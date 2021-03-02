package service

import (
	"context"
	"fmt"
	"github.com/autom8ter/machine/v2"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	stategate "github.com/stategate/stategate/gen/grpc/go"
	"github.com/stategate/stategate/internal/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
)

func eventChannelName(domain, typ string) string {
	return fmt.Sprintf("%s.%s", domain, typ)
}

func msgChannelName(domain, channel, typ string) string {
	return fmt.Sprintf("%s.%s.%s", domain, channel, typ)
}

func (s Service) setEntity(ctx context.Context, entity *stategate.Entity) (*empty.Empty, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "empty storage provider")
	}
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

	if err := s.cache.PublishEvent(ctx, e); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) getEntity(ctx context.Context, ref *stategate.EntityRef) (*stategate.Entity, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "empty storage provider")
	}
	obj, err := s.storage.GetEntity(ctx, ref)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return obj, nil
}

func (s Service) revertEntity(ctx context.Context, opts *stategate.EventRef) (*stategate.Entity, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "empty storage provider")
	}
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
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "empty storage provider")
	}
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
	if err := s.cache.PublishEvent(ctx, e); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return result, nil
}

func (s Service) delEntity(ctx context.Context, ref *stategate.EntityRef) (*empty.Empty, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "empty storage provider")
	}
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
	if err := s.cache.PublishEvent(ctx, e); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) streamEvents(opts *stategate.StreamEventOpts, server stategate.EventService_StreamServer) error {
	if s.storage == nil {
		return status.Error(codes.Unimplemented, "empty storage provider")
	}
	if s.cache == nil {
		return status.Error(codes.Unimplemented, "empty cache provider")
	}
	s.events.Subscribe(server.Context(), eventChannelName(opts.GetDomain(), opts.GetType()), func(ctx context.Context, msg machine.Message) (bool, error) {
		if err := server.Send(msg.GetBody().(*stategate.Event)); err != nil {
			return false, errors.Wrap(err, "failed to stream event")
		}
		return true, nil
	})
	return nil
}

func (s Service) searchEvents(ctx context.Context, opts *stategate.SearchEventOpts) (*stategate.Events, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "empty storage provider")
	}
	events, err := s.storage.SearchEvents(ctx, opts)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return events, nil
}

func (s Service) getEvent(ctx context.Context, opts *stategate.EventRef) (*stategate.Event, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "empty storage provider")
	}
	event, err := s.storage.GetEvent(ctx, opts)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return event, nil
}

func (s Service) searchEntities(ctx context.Context, opts *stategate.SearchEntityOpts) (*stategate.Entities, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "empty storage provider")
	}
	entitys, err := s.storage.SearchEntities(ctx, opts)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return entitys, nil
}

func (s Service) broadcastMessage(ctx context.Context, message *stategate.Message) (*empty.Empty, error) {
	if s.cache == nil {
		return nil, status.Error(codes.Unimplemented, "empty cache provider")
	}
	c, _ := auth.GetContext(ctx)
	claims, _ := structpb.NewStruct(c.Claims)
	bm := &stategate.PeerMessage{
		Id:      uuid.New().String(),
		Domain:  message.GetDomain(),
		Channel: message.GetChannel(),
		Type:    message.GetType(),
		Body:    message.GetBody(),
		Claims:  claims,
		Time:    time.Now().UnixNano(),
	}
	if err := s.cache.PublishMessage(ctx, bm); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) streamMessages(opts *stategate.StreamMessageOpts, server stategate.PeerService_StreamServer) error {
	if s.cache == nil {
		return status.Error(codes.Unimplemented, "empty cache provider")
	}
	s.messages.Subscribe(server.Context(), msgChannelName(opts.GetDomain(), opts.GetChannel(), opts.GetType()), func(ctx context.Context, msg machine.Message) (bool, error) {
		if err := server.Send(msg.GetBody().(*stategate.PeerMessage)); err != nil {
			return false, errors.Wrap(err, "failed to stream message")
		}
		return true, nil
	})
	return nil
}

func (s Service) setCache(ctx context.Context, value *stategate.Cache) (*empty.Empty, error) {
	if err := s.cache.Set(ctx, value); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) getCache(ctx context.Context, ref *stategate.CacheRef) (*stategate.Cache, error) {
	if s.cache == nil {
		return nil, status.Error(codes.Unimplemented, "empty cache provider")
	}
	resp, err := s.cache.Get(ctx, ref)
	if err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return resp, nil
}

func (s Service) delCache(ctx context.Context, ref *stategate.CacheRef) (*empty.Empty, error) {
	if s.cache == nil {
		return nil, status.Error(codes.Unimplemented, "empty cache provider")
	}
	if err := s.cache.Del(ctx, ref); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) lockMutex(ctx context.Context, mutex *stategate.Mutex) (*empty.Empty, error) {
	if s.cache == nil {
		return nil, status.Error(codes.Unimplemented, "empty cache provider")
	}
	if err := s.cache.Lock(ctx, mutex); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}

func (s Service) unlockMutex(ctx context.Context, ref *stategate.MutexRef) (*empty.Empty, error) {
	if s.cache == nil {
		return nil, status.Error(codes.Unimplemented, "empty cache provider")
	}
	if err := s.cache.Unlock(ctx, ref); err != nil {
		err.Log(s.lgger)
		return nil, err.Public()
	}
	return &empty.Empty{}, nil
}
