package kafka

import (
	"context"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	logger *logger.Logger
	reader *kafka.Reader
	writer *kafka.Writer
}

func (s *Service) Publish(ctx context.Context, event *eventgate.Event) error {
	bits, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	if err := s.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(event.Id),
		Value: bits,
	}); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetChannel(ctx context.Context) (chan *eventgate.Event, error) {
	events := make(chan *eventgate.Event)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := s.reader.ReadMessage(context.Background())
				if err != nil {
					s.logger.Error("failed to read event", zap.Error(err))
					return
				}
				var event eventgate.Event
				if err := proto.Unmarshal(msg.Value, &event); err != nil {
					s.logger.Error("failed to unmarshal event", zap.Error(err))
					continue
				}
				events <- &event
			}
		}
	}()
	return events, nil
}

func NewService(logger *logger.Logger, reader *kafka.Reader, writer *kafka.Writer) (*Service, error) {
	return &Service{
		logger: logger,
		reader: reader,
		writer: writer,
	}, nil
}

func (s *Service) Close() error {
	group := &errgroup.Group{}
	group.Go(s.writer.Close)
	group.Go(s.reader.Close)
	if err := group.Wait(); err != nil {
		return err
	}
	return nil
}
