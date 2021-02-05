package nats

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"sync"
	"time"
)

type Service struct {
	logger *logger.Logger
	conn   *nats.Conn
}

func NewService(logger *logger.Logger, conn *nats.Conn) (*Service, error) {
	return &Service{
		logger: logger,
		conn:   conn,
	}, nil
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
	if err := s.conn.Publish(r.GetChannel(), bits); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Service) Request(ctx context.Context, r *eventgate.Event) (*eventgate.Event, error) {
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
	resp, err := s.conn.Request(r.GetChannel(), bits, 30*time.Second)
	if err != nil {
		return nil, err
	}
	var event eventgate.Event
	if err := proto.Unmarshal(resp.Data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (s *Service) Receive(r *eventgate.Filter, server eventgate.EventGateService_ReceiveServer) error {
	_, ok := auth.GetContext(server.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "unauthenticated")
	}
	var (
		err error
		sub *nats.Subscription
		wg  = sync.WaitGroup{}
	)
	sub, err = s.conn.Subscribe(r.GetChannel(), func(msg *nats.Msg) {
		wg.Add(1)
		defer wg.Done()
		var event eventgate.Event
		if err := proto.Unmarshal(msg.Data, &event); err != nil {
			s.logger.Error("failed to unmarshal cloud event", zap.Error(err))
			return
		}
		if err := server.Send(&event); err != nil {
			s.logger.Error("failed to unmarshal cloud event", zap.Error(err))
			return
		}
	})
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(server.Context())
	defer cancel()
	select {
	case <-ctx.Done():
		if err := sub.Drain(); err != nil {
			return err
		}
		wg.Wait()
		return nil
	}
}

func getNatsSubject(schema, source, typ string, subject string) string {
	if schema == "" {
		schema = "*"
	}
	if source == "" {
		source = "*"
	}
	if typ == "" {
		typ = "*"
	}
	if subject == "" {
		subject = "*"
	}
	return strings.TrimSpace(fmt.Sprintf("%s.%s.%s.%s", schema, source, typ, subject))
}
