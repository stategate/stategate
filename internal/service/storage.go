package service

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/channel"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/storage"
	"github.com/autom8ter/machine/pubsub"
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

func (s Service) SetObject(ctx context.Context, object *eventgate.Object) (*empty.Empty, error) {
	c, ok := auth.GetContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	group := &errgroup.Group{}
	claims, _ := structpb.NewStruct(c.Claims)

	e := &eventgate.Event{
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

func (s Service) GetObject(ctx context.Context, ref *eventgate.ObjectRef) (*eventgate.Object, error) {
	return s.storage.GetObject(ctx, ref)
}

func (s Service) StreamEvents(opts *eventgate.StreamOpts, server eventgate.EventGateService_StreamEventsServer) error {
	_, ok := auth.GetContext(server.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "unauthenticated")
	}
	if err := s.ps.Subscribe(server.Context(), opts.GetType(), "", func(msg interface{}) bool {
		if event, ok := msg.(*eventgate.Event); ok {
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

func (s Service) SearchEvents(ctx context.Context, opts *eventgate.SearchOpts) (*eventgate.Events, error) {
	return s.storage.SearchEvents(ctx, opts)
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
