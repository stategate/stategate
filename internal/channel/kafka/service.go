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
	logger  *logger.Logger
	readerM *kafka.Reader
	writerM *kafka.Writer
	readerE *kafka.Reader
	writerE *kafka.Writer
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
	if err := s.writerE.WriteMessages(ctx, kafka.Message{
		Key:   []byte(event.Id),
		Value: bits,
	}); err != nil {
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
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := s.readerE.ReadMessage(ctx)
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

func (s *Service) PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error {
	if ctx.Err() != nil {
		return &errorz.Error{
			Type:     errorz.ErrTimeout,
			Info:     "failed to publish message",
			Err:      ctx.Err(),
			Metadata: map[string]string{},
		}
	}
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
	if err := s.writerM.WriteMessages(ctx, kafka.Message{
		Key:   []byte(message.Id),
		Value: bits,
	}); err != nil {
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
	if ctx.Err() != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrTimeout,
			Info:     "failed to setup message channel",
			Err:      ctx.Err(),
			Metadata: map[string]string{},
		}
	}
	messages := make(chan *stategate.PeerMessage)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := s.readerM.ReadMessage(ctx)
				if err != nil {
					s.logger.Error("failed to read message", zap.Error(err))
					continue
				}
				var m stategate.PeerMessage
				if err := proto.Unmarshal(msg.Value, &m); err != nil {
					s.logger.Error("failed to unmarshal message", zap.Error(err))
					continue
				}
				messages <- &m
			}
		}
	}()
	return messages, nil
}

func NewService(logger *logger.Logger, readerE, readerM *kafka.Reader, writerE, writerM *kafka.Writer) (*Service, error) {
	return &Service{
		logger:  logger,
		readerE: readerE,
		readerM: readerM,
		writerE: writerE,
		writerM: writerM,
	}, nil
}

func (s *Service) Close() error {
	group := &errgroup.Group{}
	group.Go(s.writerM.Close)
	group.Go(s.writerE.Close)
	group.Go(s.readerM.Close)
	group.Go(s.readerE.Close)
	if err := group.Wait(); err != nil {
		return err
	}
	return nil
}
