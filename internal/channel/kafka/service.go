package kafka

import (
	"context"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/autom8ter/stategate/internal/logger"
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

func (s *Service) Publish(ctx context.Context, event *stategate.Event) *errorz.Error {
	bits, err := proto.Marshal(event)
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to encode event",
			Err:  err,
			Metadata: map[string]string{
				"object_key":  event.GetState().GetKey(),
				"object_type": event.GetState().GetType(),
				"event_id":    event.GetId(),
			},
		}
	}
	if err := s.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(event.Id),
		Value: bits,
	}); err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to publish event",
			Err:  err,
			Metadata: map[string]string{
				"object_key":  event.GetState().GetKey(),
				"object_type": event.GetState().GetType(),
				"event_id":    event.GetId(),
			},
		}
	}
	return nil
}

func (s *Service) GetChannel(ctx context.Context) (chan *stategate.Event, error) {
	events := make(chan *stategate.Event)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := s.reader.ReadMessage(context.Background())
				if err != nil {
					s.logger.Error("failed to read event", zap.Error(err))
					continue
				}
				var event stategate.Event
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
