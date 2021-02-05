package kafka

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/machine/pubsub"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	logger *logger.Logger
	reader *kafka.Reader
	writer *kafka.Writer
	ps     pubsub.PubSub
}

func NewService(logger *logger.Logger, reader *kafka.Reader, writer *kafka.Writer) (*Service, error) {
	s := &Service{
		logger: logger,
		reader: reader,
		writer: writer,
		ps:     pubsub.NewPubSub(),
	}
	go func() {
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				s.logger.Error("failed to read event", zap.Error(err))
				return
			}
			var event eventgate.Event
			if err := proto.Unmarshal(msg.Value, &event); err != nil {
				s.logger.Error("failed to unmarshal event", zap.Error(err))
				continue
			}
			if err := s.ps.Publish(event.GetChannel(), &event); err != nil {
				s.logger.Error("failed to unmarshal event", zap.Error(err))
			}
		}
	}()
	return s, nil
}

func (s *Service) Send(ctx context.Context, r *eventgate.Event) (*empty.Empty, error) {
	_, ok := auth.GetContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	toSend := &eventgate.Event{
		Id:       r.GetId(),
		Channel:  r.GetChannel(),
		Data:     r.GetData(),
		Metadata: r.GetMetadata(),
		Time:     r.GetTime(),
	}
	if toSend.Id == "" {
		toSend.Id = uuid.New().String()
	}
	if toSend.Time == nil {
		toSend.Time = timestamppb.New(time.Now())
	}
	bits, err := proto.Marshal(toSend)
	if err != nil {
		return nil, err
	}
	if err := s.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(toSend.Id),
		Value: bits,
	}); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Service) Receive(r *eventgate.ReceiveOpts, server eventgate.EventGateService_ReceiveServer) error {
	_, ok := auth.GetContext(server.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "unauthenticated")
	}
	if err := s.ps.Subscribe(server.Context(), r.GetChannel(), r.GetConsumerGroup(), func(msg interface{}) bool {
		if event, ok := msg.(*eventgate.Event); ok {
			if err := server.Send(event); err != nil {
				s.logger.Error("failed to send subscription event", zap.Error(err))
			}
		} else {
			s.logger.Error("invalid event type", zap.Any("event_type", fmt.Sprintf("%T", msg)))
		}
		return true
	}); err != nil {
		return status.Error(codes.Internal, "reception failure")
	}
	return nil
}

func (s *Service) Close() error {
	group := &errgroup.Group{}
	group.Go(s.writer.Close)
	group.Go(s.reader.Close)
	if err := group.Wait(); err != nil {
		return err
	}
	s.ps.Close()
	return nil
}
