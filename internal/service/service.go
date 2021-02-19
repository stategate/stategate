package service

import (
	"context"
	"fmt"
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
				if err := svc.ps.Publish(channelName(event.GetState().GetDomain(), event.GetState().GetType()), event); err != nil {
					svc.lgger.Error("failed to unmarshal event", zap.Error(err))
				}
			}
		}
	}()
	return svc, nil
}

func (s Service) setState(ctx context.Context, object *stategate.State) (*empty.Empty, error) {
	c, _ := auth.GetContext(ctx)
	claims, _ := structpb.NewStruct(c.Claims)

	e := &stategate.Event{
		Id:     uuid.New().String(),
		State:  object,
		Claims: claims,
		Time:   time.Now().UnixNano(),
	}
	if err := s.storage.SetState(ctx, object); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	if err := s.storage.SaveEvent(ctx, e); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}

	if err := s.channel.Publish(ctx, e); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) getState(ctx context.Context, ref *stategate.StateRef) (*stategate.State, error) {
	obj, err := s.storage.GetState(ctx, ref)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return obj, nil
}

func (s Service) delState(ctx context.Context, ref *stategate.StateRef) (*empty.Empty, error) {
	if err := s.storage.DelState(ctx, ref); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) streamEvents(opts *stategate.StreamOpts, server stategate.EventService_StreamServer) error {
	if err := s.ps.Subscribe(server.Context(), channelName(opts.GetDomain(), opts.GetType()), func(msg interface{}) bool {
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

func (s Service) searchEvents(ctx context.Context, opts *stategate.SearchEventOpts) (*stategate.Events, error) {
	events, err := s.storage.SearchEvents(ctx, opts)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return events, nil
}

func (s Service) searchStates(ctx context.Context, opts *stategate.SearchStateOpts) (*stategate.StateValues, error) {
	objects, err := s.storage.SearchState(ctx, opts)
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

func (s *Service) EventServiceServer() stategate.EventServiceServer {
	return &eventService{svc: s}
}

func (s *Service) StateServiceServer() stategate.StateServiceServer {
	return &objectService{svc: s}
}

type eventService struct {
	svc *Service
}

func (e eventService) Stream(opts *stategate.StreamOpts, server stategate.EventService_StreamServer) error {
	return e.svc.streamEvents(opts, server)
}

func (e eventService) Search(ctx context.Context, opts *stategate.SearchEventOpts) (*stategate.Events, error) {
	return e.svc.searchEvents(ctx, opts)
}

type objectService struct {
	svc *Service
}

func (o objectService) Set(ctx context.Context, object *stategate.State) (*empty.Empty, error) {
	return o.svc.setState(ctx, object)
}

func (o objectService) Get(ctx context.Context, ref *stategate.StateRef) (*stategate.State, error) {
	return o.svc.getState(ctx, ref)
}

func (o objectService) Del(ctx context.Context, ref *stategate.StateRef) (*empty.Empty, error) {
	return o.svc.delState(ctx, ref)
}

func (o objectService) Search(ctx context.Context, opts *stategate.SearchStateOpts) (*stategate.StateValues, error) {
	return o.svc.searchStates(ctx, opts)
}

func channelName(tenant, typ string) string {
	return fmt.Sprintf("%s.%s", tenant, typ)
}
