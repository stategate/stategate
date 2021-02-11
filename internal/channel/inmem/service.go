package inmem

import (
	"context"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/constants"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/machine/pubsub"
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

func (s *Service) Publish(ctx context.Context, event *eventgate.Event) error {
	return s.ps.Publish(event.GetObject().GetType(), event)
}

func (s *Service) GetChannel(ctx context.Context) (chan *eventgate.Event, error) {
	events := make(chan *eventgate.Event)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		s.ps.Subscribe(ctx, constants.BackendChannel, "", func(msg interface{}) bool {
			events <- msg.(*eventgate.Event)
			return true
		})
	}()
	return events, nil
}

func (s *Service) Close() error {
	s.ps.Close()
	return nil
}
