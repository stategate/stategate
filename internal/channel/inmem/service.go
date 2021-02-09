package inmem

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/storage"
	"github.com/autom8ter/machine/pubsub"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	logger  *logger.Logger
	ps      pubsub.PubSub
	storage storage.Provider
}

func NewService(logger *logger.Logger, storage storage.Provider) *Service {
	return &Service{
		logger:  logger,
		ps:      pubsub.NewPubSub(),
		storage: storage,
	}
}

func (s *Service) Send(ctx context.Context, r *eventgate.Event) (*empty.Empty, error) {
	c, ok := auth.GetContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	claims, _ := structpb.NewStruct(c.Claims)
	toSend := &eventgate.Event{
		Id:       r.GetId(),
		Channel:  r.GetChannel(),
		Data:     r.GetData(),
		Metadata: r.GetMetadata(),
		Time:     r.GetTime(),
		Claims:   claims,
	}
	if toSend.Id == "" {
		toSend.Id = uuid.New().String()
	}
	if toSend.Time == nil {
		toSend.Time = timestamppb.New(time.Now())
	}
	if s.storage != nil {
		if err := s.storage.SaveEvent(ctx, toSend); err != nil {
			return nil, err
		}
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

func (s *Service) History(ctx context.Context, opts *eventgate.HistoryOpts) (*eventgate.Events, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "backend timeseries storage provider not registered")
	}
	return s.storage.GetEvents(ctx, opts)
}
