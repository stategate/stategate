package redis

import (
	"context"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/constants"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

type Service struct {
	logger *logger.Logger
	conn   *redis.Client
}

func NewService(logger *logger.Logger, conn *redis.Client) *Service {
	return &Service{logger: logger, conn: conn}
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
	return nil
}
