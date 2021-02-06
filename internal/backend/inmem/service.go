package inmem

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/machine/pubsub"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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
	if err := s.ps.Publish(r.Channel, toSend); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Service) Receive(r *eventgate.ReceiveOpts, server eventgate.EventGateService_ReceiveServer) error {
	_, ok := auth.GetContext(server.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "unauthenticated")
	}
	if err := s.ps.Subscribe(server.Context(), r.GetChannel(), "", func(msg interface{}) bool {
		if event, ok := msg.(*eventgate.Event); ok {
			if err := server.Send(event); err != nil {
				s.logger.Error("failed to send subscription event", zap.Error(err))
			}
		} else {
			s.logger.Error("invalid event type", zap.Any("event_type", fmt.Sprintf("%T", msg)))
		}
		return true
	}); err != nil {
		s.logger.Error("reception failure", zap.Error(err))
		return status.Error(codes.Internal, "reception failure")
	}
	return nil
}

func (s *Service) Close() error {
	s.ps.Close()
	return nil
}
