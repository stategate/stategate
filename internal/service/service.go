package service

import (
	"context"
	"fmt"
	"github.com/autom8ter/machine/v2"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/auth"
	"github.com/autom8ter/stategate/internal/channel"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/autom8ter/stategate/internal/storage"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

type Service struct {
	storage storage.Provider
	channel channel.Provider
	lgger   *logger.Logger
	ps      machine.Machine
	cancel  func()
}

func NewService(storage storage.Provider, channel channel.Provider, lgger *logger.Logger, machne machine.Machine) (*Service, error) {
	ctx, cancel := context.WithCancel(context.Background())
	svc := &Service{
		storage: storage,
		lgger:   lgger,
		channel: channel,
		ps:      machne,
		cancel:  cancel,
	}
	ch, err := channel.GetChannel(ctx)
	if err != nil {
		return nil, err
	}
	go func() {
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-ch:
				svc.ps.Publish(ctx, machine.Msg{
					Channel: channelName(event.GetEntity().GetDomain(), event.GetEntity().GetType()),
					Body:    event,
				})
			}
		}
	}()
	return svc, nil
}

func (s Service) setEntity(ctx context.Context, entity *stategate.Entity) (*empty.Empty, error) {
	c, _ := auth.GetContext(ctx)
	claims, _ := structpb.NewStruct(c.Claims)

	e := &stategate.Event{
		Id:     uuid.New().String(),
		Entity: entity,
		Method: c.Method,
		Claims: claims,
		Time:   time.Now().UnixNano(),
	}
	if err := s.storage.SetEntity(ctx, entity); err != nil {
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

func (s Service) getEntity(ctx context.Context, ref *stategate.EntityRef) (*stategate.Entity, error) {
	obj, err := s.storage.GetEntity(ctx, ref)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return obj, nil
}

func (s Service) revertEntity(ctx context.Context, opts *stategate.EventRef) (*stategate.Entity, error) {
	event, err := s.storage.GetEvent(ctx, &stategate.EventRef{
		Domain: opts.GetDomain(),
		Type:   opts.GetType(),
		Key:    opts.GetKey(),
		Id:     opts.GetId(),
	})
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	entity := event.GetEntity()
	if _, err := s.setEntity(ctx, entity); err != nil {
		return nil, err
	}
	return entity, nil
}

func (s Service) editEntity(ctx context.Context, entity *stategate.Entity) (*stategate.Entity, error) {
	c, _ := auth.GetContext(ctx)
	claims, _ := structpb.NewStruct(c.Claims)
	result, err := s.storage.EditEntity(ctx, entity)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	if err := result.Validate(); err != nil {
		return nil, err
	}
	e := &stategate.Event{
		Id:     uuid.New().String(),
		Entity: result,
		Method: c.Method,
		Claims: claims,
		Time:   time.Now().UnixNano(),
	}
	if err := s.storage.SaveEvent(ctx, e); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	if err := s.channel.Publish(ctx, e); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return result, nil
}

func (s Service) delEntity(ctx context.Context, ref *stategate.EntityRef) (*empty.Empty, error) {
	c, _ := auth.GetContext(ctx)
	claims, _ := structpb.NewStruct(c.Claims)
	if err := s.storage.DelEntity(ctx, ref); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	entity, err := s.getEntity(ctx, ref)
	if err != nil {
		return nil, err
	}
	e := &stategate.Event{
		Id:     uuid.New().String(),
		Entity: entity,
		Method: c.Method,
		Claims: claims,
		Time:   time.Now().UnixNano(),
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

func (s Service) streamEvents(opts *stategate.StreamOpts, server stategate.EventService_StreamServer) error {
	s.ps.Subscribe(server.Context(), channelName(opts.GetDomain(), opts.GetType()), func(ctx context.Context, msg machine.Message) (bool, error) {
		if err := server.Send(msg.GetBody().(*stategate.Event)); err != nil {
			return false, errors.Wrap(err, "failed to stream event")
		}
		return true, nil
	})
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

func (s Service) getEvent(ctx context.Context, opts *stategate.EventRef) (*stategate.Event, error) {
	event, err := s.storage.GetEvent(ctx, opts)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return event, nil
}

func (s Service) searchEntities(ctx context.Context, opts *stategate.SearchEntityOpts) (*stategate.Entities, error) {
	entitys, err := s.storage.SearchEntities(ctx, opts)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return entitys, nil
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

func (s *Service) EntityServiceServer() stategate.EntityServiceServer {
	return &entityService{svc: s}
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

func (e eventService) Get(ctx context.Context, ref *stategate.EventRef) (*stategate.Event, error) {
	return e.svc.getEvent(ctx, ref)
}

type entityService struct {
	svc *Service
}

func (o entityService) Set(ctx context.Context, entity *stategate.Entity) (*empty.Empty, error) {
	return o.svc.setEntity(ctx, entity)
}

func (o entityService) Edit(ctx context.Context, entity *stategate.Entity) (*stategate.Entity, error) {
	return o.svc.editEntity(ctx, entity)
}

func (o entityService) Revert(ctx context.Context, opts *stategate.EventRef) (*stategate.Entity, error) {
	return o.svc.revertEntity(ctx, opts)
}

func (o entityService) Get(ctx context.Context, ref *stategate.EntityRef) (*stategate.Entity, error) {
	return o.svc.getEntity(ctx, ref)
}

func (o entityService) Del(ctx context.Context, ref *stategate.EntityRef) (*empty.Empty, error) {
	return o.svc.delEntity(ctx, ref)
}

func (o entityService) Search(ctx context.Context, opts *stategate.SearchEntityOpts) (*stategate.Entities, error) {
	return o.svc.searchEntities(ctx, opts)
}

func channelName(tenant, typ string) string {
	return fmt.Sprintf("%s.%s", tenant, typ)
}
