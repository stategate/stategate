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
	machine     machine.Machine
}

func NewService(logger *logger.Logger) *Service {
	return &Service{
		logger: logger,
		machine:     machine.New(),
	}
}

func (s *Service) PublishEvent(ctx context.Context, event *stategate.Event) *errorz.Error {
	s.machine.Publish(ctx, machine.Msg{
		Channel: constants.EventChannel,
		Body:    event,
	})
	return nil
}

func (s *Service) GetEventChannel(ctx context.Context) (chan *stategate.Event, *errorz.Error) {
	events := make(chan *stategate.Event)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		s.machine.Subscribe(ctx, constants.EventChannel, func(ctx context.Context, msg machine.Message) (bool, error) {
			events <- msg.GetBody().(*stategate.Event)
			return true, nil
		})
	}()
	return events, nil
}

func (s *Service) GetMessageChannel(ctx context.Context) (chan *stategate.PeerMessage, *errorz.Error) {
	messages := make(chan *stategate.PeerMessage)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		s.machine.Subscribe(ctx, constants.MessageChannel, func(ctx context.Context, msg machine.Message) (bool, error) {
			messages <- msg.GetBody().(*stategate.PeerMessage)
			return true, nil
		})
	}()
	return messages, nil
}

func (s *Service) PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error {
	s.machine.Publish(ctx, machine.Msg{
		Channel: constants.MessageChannel,
		Body:    message,
	})
	return nil
}

func (s *Service) Close() error {
	s.machine.Close()
	return nil
}
