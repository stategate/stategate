package stan

import (
	"context"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/constants"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/stan.go"
	"go.uber.org/zap"
)

type Service struct {
	eventsChan string
	logger     *logger.Logger
	conn       stan.Conn
	sub        stan.Subscription
}

func NewService(logger *logger.Logger, conn stan.Conn) (*Service, error) {
	return &Service{
		logger: logger,
		conn:   conn,
	}, nil
}

func (s *Service) Publish(ctx context.Context, event *stategate.Event) error {
	bits, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	if err := s.conn.Publish(constants.BackendChannel, bits); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetChannel(ctx context.Context) (chan *stategate.Event, error) {
	events := make(chan *stategate.Event)
	sub, err := s.conn.Subscribe(constants.BackendChannel, func(msg *stan.Msg) {
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
	return s.conn.Close()
}
