package stan

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/constants"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/storage"
	"github.com/autom8ter/machine/pubsub"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	eventsChan string
	logger     *logger.Logger
	conn       stan.Conn
	ps         pubsub.PubSub
	sub        stan.Subscription
	storage    storage.Provider
}

func NewService(logger *logger.Logger, conn stan.Conn, storage storage.Provider) (*Service, error) {
	s := &Service{
		logger:     logger,
		conn:       conn,
		ps:         pubsub.NewPubSub(),
		eventsChan: constants.BackendChannel,
		storage:    storage,
	}
	sub, err := s.conn.Subscribe(s.eventsChan, func(msg *stan.Msg) {
		var event eventgate.Event
		if err := proto.Unmarshal(msg.Data, &event); err != nil {
			s.logger.Error("failed to unmarshal event", zap.Error(err))
			return
		}
		if err := s.ps.Publish(event.GetChannel(), &event); err != nil {
			s.logger.Error("failed to unmarshal event", zap.Error(err))
			return
		}
	})
	if err != nil {
		return nil, err
	}
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
	group := errgroup.Group{}
	group.Go(func() error {
		return s.conn.Publish(s.eventsChan, bits)
	})
	if s.storage != nil {
		group.Go(func() error {
			return s.storage.SaveEvent(ctx, toSend)
		})
	}
	if err := group.Wait(); err != nil {
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
		return status.Error(codes.Internal, "reception failure")
	}
	return nil
}

func (s *Service) Close() error {
	if err := s.sub.Unsubscribe(); err != nil {
		return err
	}
	if err := s.sub.Close(); err != nil {
		return err
	}
	s.ps.Close()
	return nil
}

func (s *Service) History(ctx context.Context, opts *eventgate.HistoryOpts) (*eventgate.Events, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "backend timeseries storage provider not registered")
	}
	return s.storage.GetEvents(ctx, opts)
}
