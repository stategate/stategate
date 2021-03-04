package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	stategate "github.com/stategate/stategate/gen/grpc/go"
	"github.com/stategate/stategate/internal/constants"
	"github.com/stategate/stategate/internal/errorz"
	"github.com/stategate/stategate/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	logger *logger.Logger
	conn   *nats.Conn
	sub    *nats.Subscription
}

func NewService(logger *logger.Logger, conn *nats.Conn) (*Service, error) {
	return &Service{
		logger: logger,
		conn:   conn,
	}, nil
}

func (s *Service) PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error {
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
	if err := s.conn.Publish(constants.MessageChannel, bits); err != nil {
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
	peers := make(chan *stategate.PeerMessage)
	sub, err := s.conn.Subscribe(constants.MessageChannel, func(msg *nats.Msg) {
		var peer stategate.PeerMessage
		if err := proto.Unmarshal(msg.Data, &peer); err != nil {
			s.logger.Error("failed to unmarshal peer", zap.Error(err))
			return
		}
		peers <- &peer
	})
	if err != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to get peer channel",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	s.sub = sub
	return peers, nil
}

func (s *Service) PublishEvent(ctx context.Context, event *stategate.Event) *errorz.Error {
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
	if err := s.conn.Publish(constants.EventChannel, bits); err != nil {
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
	events := make(chan *stategate.Event)
	sub, err := s.conn.Subscribe(constants.EventChannel, func(msg *nats.Msg) {
		var event stategate.Event
		if err := proto.Unmarshal(msg.Data, &event); err != nil {
			s.logger.Error("failed to unmarshal event", zap.Error(err))
			return
		}
		events <- &event
	})
	if err != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to get event channel",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	s.sub = sub
	return events, nil
}

func (s *Service) Close() error {
	if s.sub != nil {
		if err := s.sub.Drain(); err != nil {
			return err
		}
	}
	s.conn.Close()
	return nil
}
