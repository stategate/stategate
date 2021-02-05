package redis

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/constants"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/machine/pubsub"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	eventsChan string
	logger     *logger.Logger
	conn       *redis.Client
	ps         pubsub.PubSub
	sub        *redis.PubSub
}

func NewService(logger *logger.Logger, client *redis.Client) (*Service, error) {
	s := &Service{
		logger:     logger,
		conn:       client,
		ps:         pubsub.NewPubSub(),
		eventsChan: constants.BackendChannel,
	}
	sub := s.conn.Subscribe(context.Background(), s.eventsChan)
	go func() {
		ch := sub.Channel()
		for {
			select {
			case msg := <-ch:
				var event eventgate.Event
				if err := proto.Unmarshal([]byte(msg.Payload), &event); err != nil {
					s.logger.Error("failed to unmarshal event", zap.Error(err))
					return
				}
				if err := s.ps.Publish(event.GetChannel(), &event); err != nil {
					s.logger.Error("failed to unmarshal event", zap.Error(err))
					return
				}
			}
		}
	}()
	s.sub = sub
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
	if _, err := s.conn.Publish(ctx, s.eventsChan, bits).Result(); err != nil {
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
	if err := s.sub.Close(); err != nil {
		return err
	}
	s.ps.Close()
	return nil
}
