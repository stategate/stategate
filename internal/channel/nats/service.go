package nats

import (
	"context"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/constants"
	"github.com/autom8ter/eventgate/internal/logger"
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

func (s *Service) Publish(ctx context.Context, event *eventgate.Event) error {
	bits, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	if err := s.conn.Publish(constants.BackendChannel, bits); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetChannel(ctx context.Context) (chan *eventgate.Event, error) {
	events := make(chan *eventgate.Event)
	sub, err := s.conn.Subscribe(constants.BackendChannel, func(msg *nats.Msg) {
		var event eventgate.Event
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
