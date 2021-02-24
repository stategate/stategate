package inmem

import (
	"context"
	"github.com/autom8ter/machine/v2"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/constants"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/autom8ter/stategate/internal/logger"
)

type Service struct {
	logger *logger.Logger
	ps     machine.Machine
}

func NewService(logger *logger.Logger) *Service {
	return &Service{
		logger: logger,
		ps:     machine.New(),
	}
}

func (s *Service) Publish(ctx context.Context, event *stategate.Event) *errorz.Error {
	s.ps.Publish(ctx, machine.Msg{
		Channel: constants.BackendChannel,
		Body:    event,
	})
	return nil
}

func (s *Service) GetChannel(ctx context.Context) (chan *stategate.Event, error) {
	events := make(chan *stategate.Event)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		s.ps.Subscribe(ctx, constants.BackendChannel, func(ctx context.Context, msg machine.Message) (bool, error) {
			events <- msg.GetBody().(*stategate.Event)
			return true, nil
		})
	}()
	return events, nil
}

func (s *Service) Close() error {
	s.ps.Close()
	return nil
}
