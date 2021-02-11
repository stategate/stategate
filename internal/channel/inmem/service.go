package inmem

import (
	"context"
	"github.com/autom8ter/machine/pubsub"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/constants"
	"github.com/autom8ter/stategate/internal/logger"
)

type Service struct {
	logger *logger.Logger
	ps     pubsub.PubSub
}

func NewService(logger *logger.Logger) *Service {
	return &Service{
		logger: logger,
		ps:     pubsub.NewPubSub(),
	}
}

func (s *Service) Publish(ctx context.Context, event *stategate.Event) error {
	return s.ps.Publish(constants.BackendChannel, event)
}

func (s *Service) GetChannel(ctx context.Context) (chan *stategate.Event, error) {
	events := make(chan *stategate.Event)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		s.ps.Subscribe(ctx, constants.BackendChannel, "", func(msg interface{}) bool {
			events <- msg.(*stategate.Event)
			return true
		})
	}()
	return events, nil
}

func (s *Service) Close() error {
	s.ps.Close()
	return nil
}
