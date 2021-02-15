package service

import (
	"context"
	"github.com/autom8ter/machine/pubsub"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/auth"
	"github.com/autom8ter/stategate/internal/channel"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/autom8ter/stategate/internal/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
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
		if err := s.storage.SetObject(ctx, object); err != nil {
			err.Log(s.lgger)
			return err.Public()
		}
		return nil
	})
	group.Go(func() error {
		if err := s.storage.SaveEvent(ctx, e); err != nil {
			err.Log(s.lgger)
			return err.Public()
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		return nil, err
	}
	if err := s.channel.Publish(ctx, e); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) GetObject(ctx context.Context, ref *stategate.ObjectRef) (*stategate.Object, error) {
	obj, err := s.storage.GetObject(ctx, ref)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return obj, nil
}

func (s Service) StreamEvents(opts *stategate.StreamOpts, server stategate.StateGateService_StreamEventsServer) error {
	if err := s.ps.Subscribe(server.Context(), opts.GetType(), "", func(msg interface{}) bool {
		if err := server.Send(msg.(*stategate.Event)); err != nil {
			e := &errorz.Error{
				Type:     errorz.ErrUnknown,
				Info:     "failed to stream event",
				Err:      err,
				Metadata: map[string]string{},
			}
			e.Log(s.lgger)
		}
		return true
	}); err != nil {
		e := &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to stream events",
			Err:      err,
			Metadata: map[string]string{},
		}
		e.Log(s.lgger)
		return e.Public()
	}
	return nil
}

func (s Service) SearchEvents(ctx context.Context, opts *stategate.SearchEventOpts) (*stategate.Events, error) {
	events, err := s.storage.SearchEvents(ctx, opts)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return events, nil
}

func (s Service) SearchObjects(ctx context.Context, opts *stategate.SearchObjectOpts) (*stategate.Objects, error) {
	objects, err := s.storage.SearchObjects(ctx, opts)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return objects, nil
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
