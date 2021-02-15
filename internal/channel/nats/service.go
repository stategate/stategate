package nats

import (
	"context"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/constants"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
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

func (s *Service) Publish(ctx context.Context, event *stategate.Event) *errorz.Error {
	bits, err := proto.Marshal(event)
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to encode event",
			Err:  err,
			Metadata: map[string]string{
				"object_key":  event.GetObject().GetKey(),
				"object_type": event.GetObject().GetType(),
				"event_id":    event.GetId(),
			},
		}
	}
	if err := s.conn.Publish(constants.BackendChannel, bits); err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to publish event",
			Err:  err,
			Metadata: map[string]string{
				"object_key":  event.GetObject().GetKey(),
				"object_type": event.GetObject().GetType(),
				"event_id":    event.GetId(),
			},
		}
	}
	return nil
}

func (s *Service) GetChannel(ctx context.Context) (chan *stategate.Event, error) {
	events := make(chan *stategate.Event)
	sub, err := s.conn.Subscribe(constants.BackendChannel, func(msg *nats.Msg) {
		var event stategate.Event
		if err := proto.Unmarshal(msg.Data, &event); err != nil {
			s.logger.Error("failed to unmarshal event", zap.Error(err))
			return
		}
		events <- &event
	})
	if err != nil {
		return nil, err
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
