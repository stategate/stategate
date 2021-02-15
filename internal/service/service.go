package service

import (
	"context"
	"fmt"
	"github.com/autom8ter/machine/pubsub"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/auth"
	"github.com/autom8ter/stategate/internal/channel"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/autom8ter/stategate/internal/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

type Service struct {
	storage storage.Provider
	channel channel.Provider
	lgger   *logger.Logger
	ps      pubsub.PubSub
	cancel  func()
}

func NewService(storage storage.Provider, channel channel.Provider, lgger *logger.Logger) (*Service, error) {
	ctx, cancel := context.WithCancel(context.Background())
	svc := &Service{
		storage: storage,
		lgger:   lgger,
		channel: channel,
		ps:      pubsub.NewPubSub(),
		cancel:  cancel,
	}
	ch, err := channel.GetChannel(ctx)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-ch:
				if err := svc.ps.Publish(event.GetObject().GetType(), event); err != nil {
					svc.lgger.Error("failed to unmarshal event", zap.Error(err))
					return
				}
			}
		}
	}()
	return svc, nil
}

func (s Service) SetObject(ctx context.Context, object *stategate.Object) (*empty.Empty, error) {
	c, _ := auth.GetContext(ctx)
	group := &errgroup.Group{}
	claims, _ := structpb.NewStruct(c.Claims)

	e := &stategate.Event{
		Id:     uuid.New().String(),
		Object: object,
		Claims: claims,
		Time:   time.Now().UnixNano(),
	}

	group.Go(func() error {
		return s.storage.SetObject(ctx, object)
	})
	group.Go(func() error {
		return s.storage.SaveEvent(ctx, e)
	})

	if err := group.Wait(); err != nil {
		return nil, errors.Wrap(err, "failed to save state change")
	}
	if err := s.channel.Publish(ctx, e); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s Service) GetObject(ctx context.Context, ref *stategate.ObjectRef) (*stategate.Object, error) {
	return s.storage.GetObject(ctx, ref)
}

func (s Service) StreamEvents(opts *stategate.StreamOpts, server stategate.StateGateService_StreamEventsServer) error {
	if err := s.ps.Subscribe(server.Context(), opts.GetType(), "", func(msg interface{}) bool {
		if event, ok := msg.(*stategate.Event); ok {
			if err := server.Send(event); err != nil {
				s.lgger.Error("failed to send subscription event", zap.Error(err))
			}
		} else {
			s.lgger.Error("invalid event type", zap.Any("event_type", fmt.Sprintf("%T", msg)))
		}
		return true
	}); err != nil {
		return status.Error(codes.Internal, "reception failure")
	}
	return nil
}

func (s Service) SearchEvents(ctx context.Context, opts *stategate.SearchEventOpts) (*stategate.Events, error) {
	return s.storage.SearchEvents(ctx, opts)
}

func (s Service) SearchObjects(ctx context.Context, opts *stategate.SearchObjectOpts) (*stategate.Objects, error) {
	return s.storage.SearchObjects(ctx, opts)
}

func (s Service) Close() error {
	s.cancel()
	if err := s.channel.Close(); err != nil {
		return err
	}
	if err := s.storage.Close(); err != nil {
		return err
	}
	s.ps.Close()
	return nil
}
