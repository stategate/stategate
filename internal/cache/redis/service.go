package redis

import (
	"context"
	"fmt"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/constants"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"time"
)

type Service struct {
	logger *logger.Logger
	conn   *redis.Client
}

func NewService(logger *logger.Logger, conn *redis.Client) *Service {
	return &Service{logger: logger, conn: conn}
}

func (s *Service) Get(ctx context.Context, ref *stategate.CacheRef) (*stategate.Cache, *errorz.Error) {
	res, err := s.conn.Get(ctx, cachedKeyName(ref.GetDomain(), ref.GetKey())).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to get cached value",
				Err:  err,
				Metadata: map[string]string{
					"cache_key":    ref.GetKey(),
					"cache_domain": ref.GetDomain(),
				},
			}
		}
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to get cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    ref.GetKey(),
				"cache_domain": ref.GetDomain(),
			},
		}
	}
	var cachedVal stategate.Cache
	if err := proto.Unmarshal([]byte(res), &cachedVal); err != nil {
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to get cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    ref.GetKey(),
				"cache_domain": ref.GetDomain(),
			},
		}
	}
	return &cachedVal, nil
}

func (s *Service) Set(ctx context.Context, value *stategate.Cache) *errorz.Error {
	bits, err := proto.Marshal(value)
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    value.GetKey(),
				"cache_domain": value.GetDomain(),
			},
		}
	}
	exp := time.Duration(0)
	if value.GetExp() != nil {
		exp = value.GetExp().AsTime().Sub(time.Now())
	}
	value.GetExp().AsTime().Sub(time.Now())
	if err := s.conn.Set(ctx, cachedKeyName(value.GetDomain(), value.GetKey()), bits, exp).Err(); err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    value.GetKey(),
				"cache_domain": value.GetDomain(),
			},
		}
	}
	return nil
}

func (s *Service) Del(ctx context.Context, ref *stategate.CacheRef) *errorz.Error {
	if err := s.conn.Del(ctx, cachedKeyName(ref.GetDomain(), ref.GetKey())).Err(); err != nil {
		if err == redis.Nil {
			return nil
		}
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to delete cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    ref.GetKey(),
				"cache_domain": ref.GetDomain(),
			},
		}
	}
	return nil
}

func (s *Service) Lock(ctx context.Context, ref *stategate.Mutex) *errorz.Error {
	exp := time.Duration(0)
	if ref.GetExp() != nil {
		exp = ref.GetExp().AsTime().Sub(time.Now())
	}
	gotlock, err := s.conn.SetNX(ctx, cachedLockName(ref.GetDomain(), ref.GetKey()), true, exp).Result()
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to lock mutex",
			Err:  err,
			Metadata: map[string]string{
				"mutex_key":    ref.GetKey(),
				"mutex_domain": ref.GetDomain(),
			},
		}
	}
	if !gotlock {
		return &errorz.Error{
			Type: errorz.ErrLocked,
			Info: "failed to lock mutex",
			Err:  err,
			Metadata: map[string]string{
				"mutex_key":    ref.GetKey(),
				"mutex_domain": ref.GetDomain(),
			},
		}
	}
	return nil
}

func (s *Service) Unlock(ctx context.Context, ref *stategate.MutexRef) *errorz.Error {
	err := s.conn.Del(ctx, cachedLockName(ref.GetDomain(), ref.GetKey())).Err()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to unlock mutex",
			Err:  err,
			Metadata: map[string]string{
				"mutex_key":    ref.GetKey(),
				"mutex_domain": ref.GetDomain(),
			},
		}
	}
	return nil
}

func (s *Service) PublishEvent(ctx context.Context, event *stategate.Event) *errorz.Error {
	if ctx.Err() != nil {
		return &errorz.Error{
			Type:     errorz.ErrTimeout,
			Info:     "failed to publish event",
			Err:      ctx.Err(),
			Metadata: map[string]string{},
		}
	}
	bits, err := proto.Marshal(event)
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to encode event",
			Err:  err,
			Metadata: map[string]string{
				"entity_key":  event.GetEntity().GetKey(),
				"entity_type": event.GetEntity().GetType(),
				"event_id":    event.GetId(),
			},
		}
	}
	if _, err := s.conn.Publish(ctx, constants.EventChannel, bits).Result(); err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to publish event",
			Err:  err,
			Metadata: map[string]string{
				"entity_key":  event.GetEntity().GetKey(),
				"entity_type": event.GetEntity().GetType(),
				"event_id":    event.GetId(),
			},
		}
	}
	return nil
}

func (s *Service) GetEventChannel(ctx context.Context) (chan *stategate.Event, *errorz.Error) {
	if ctx.Err() != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrTimeout,
			Info:     "failed to setup event channel",
			Err:      ctx.Err(),
			Metadata: map[string]string{},
		}
	}
	sub := s.conn.Subscribe(ctx, constants.EventChannel)
	events := make(chan *stategate.Event)
	go func() {
		ch := sub.Channel()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				sub.Close()
				return
			case msg := <-ch:
				var event stategate.Event
				if err := proto.Unmarshal([]byte(msg.Payload), &event); err != nil {
					s.logger.Error("failed to unmarshal event", zap.Error(err))
					continue
				}
				events <- &event
			}
		}
	}()
	return events, nil
}

func (s *Service) PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error {
	if ctx.Err() != nil {
		return &errorz.Error{
			Type:     errorz.ErrTimeout,
			Info:     "failed to publish message",
			Err:      ctx.Err(),
			Metadata: map[string]string{},
		}
	}
	bits, err := proto.Marshal(message)
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to encode message",
			Err:  err,
			Metadata: map[string]string{
				"message_domain":  message.GetDomain(),
				"message_type":    message.GetType(),
				"message_channel": message.GetChannel(),
				"message_id":      message.GetId(),
			},
		}
	}
	if _, err := s.conn.Publish(ctx, constants.MessageChannel, bits).Result(); err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to publish message",
			Err:  err,
			Metadata: map[string]string{
				"message_domain":  message.GetDomain(),
				"message_type":    message.GetType(),
				"message_channel": message.GetChannel(),
				"message_id":      message.GetId(),
			},
		}
	}
	return nil
}

func (s *Service) GetMessageChannel(ctx context.Context) (chan *stategate.PeerMessage, *errorz.Error) {
	if ctx.Err() != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrTimeout,
			Info:     "failed to setup message channel",
			Err:      ctx.Err(),
			Metadata: map[string]string{},
		}
	}
	sub := s.conn.Subscribe(ctx, constants.MessageChannel)
	messages := make(chan *stategate.PeerMessage)
	go func() {
		ch := sub.Channel()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				sub.Close()
				return
			case msg := <-ch:
				var m stategate.PeerMessage
				if err := proto.Unmarshal([]byte(msg.Payload), &m); err != nil {
					s.logger.Error("failed to unmarshal message", zap.Error(err))
					continue
				}
				messages <- &m
			}
		}
	}()
	return messages, nil
}

func (s *Service) Close() error {
	return s.conn.Close()
}

func cachedKeyName(domain string, key string) string {
	return fmt.Sprintf("%s.%s", domain, key)
}

func cachedLockName(domain string, key string) string {
	return fmt.Sprintf("mutex.%s.%s", domain, key)
}
