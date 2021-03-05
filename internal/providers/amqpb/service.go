package amqpb

import (
	"context"
	"github.com/pkg/errors"
	stategate "github.com/stategate/stategate/gen/grpc/go"
	"github.com/stategate/stategate/internal/constants"
	"github.com/stategate/stategate/internal/errorz"
	"github.com/stategate/stategate/internal/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	logger *logger.Logger
	conn   *amqp.Connection
	ch     *amqp.Channel
	eq     amqp.Queue
	mq     amqp.Queue
}

func NewService(logger *logger.Logger, conn *amqp.Connection) (*Service, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup amqp channel")
	}
	mq, err := ch.QueueDeclare(
		constants.MessageChannel,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup amqp message queue")
	}
	eq, err := ch.QueueDeclare(
		constants.EventChannel,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup amqp event queue")
	}
	return &Service{
		logger: logger,
		conn:   conn,
		ch:     ch,
		mq:     mq,
		eq:     eq,
	}, nil
}

func (s Service) PublishEvent(ctx context.Context, event *stategate.Event) *errorz.Error {
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
	if err := s.ch.Publish(
		"",        // exchange
		s.eq.Name, // routing key
		false,     // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/protobuf",
			Body:        bits,
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

func (s Service) GetEventChannel(ctx context.Context) (chan *stategate.Event, *errorz.Error) {
	events := make(chan *stategate.Event)
	msgs, err := s.ch.Consume(
		s.eq.Name, // queue
		"",
		true,
		false,
		false,
		true,
		nil,
	)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				close(events)
				return
			case msg := <-msgs:
				var event stategate.Event
				if err := proto.Unmarshal([]byte(msg.Body), &event); err != nil {
					s.logger.Error("failed to unmarshal event", zap.Error(err))
					continue
				}
				events <- &event
			}
		}
	}()
	if err != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to get event channel",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	return events, nil
}

func (s Service) PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error {
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
	if err := s.ch.Publish(
		"",        // exchange
		s.mq.Name, // routing key
		false,     // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/protobuf",
			Body:        bits,
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

func (s Service) GetMessageChannel(ctx context.Context) (chan *stategate.PeerMessage, *errorz.Error) {
	messages := make(chan *stategate.PeerMessage)
	msgs, err := s.ch.Consume(
		s.mq.Name, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		true,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, &errorz.Error{
			Type:     errorz.ErrUnknown,
			Info:     "failed to get peer message channel",
			Err:      err,
			Metadata: map[string]string{},
		}
	}
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				close(messages)
				return
			case msg := <-msgs:
				var m stategate.PeerMessage
				if err := proto.Unmarshal([]byte(msg.Body), &m); err != nil {
					s.logger.Error("failed to unmarshal message", zap.Error(err))
					continue
				}
				messages <- &m
			}
		}
	}()

	return messages, nil
}

func (s Service) Close() error {
	if s.ch != nil {
		if err := s.ch.Close(); err != nil {
			return errors.Wrap(err, "failed to close amqp channel")
		}
	}
	return s.conn.Close()
}
