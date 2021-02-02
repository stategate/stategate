package nats

import (
	"context"
	"fmt"
	cloudEventsProxy "github.com/autom8ter/cloudEventsProxy/gen/grpc/go"
	"github.com/autom8ter/cloudEventsProxy/internal/auth"
	"github.com/autom8ter/cloudEventsProxy/internal/logger"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *Service) Send(ctx context.Context, r *cloudEventsProxy.CloudEventInput) (*empty.Empty, error) {
	c, ok := auth.GetContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	toSend := &cloudEventsProxy.CloudEvent{
		Specversion: r.GetSpecversion(),
		Id:          uuid.New().String(),
		Source:      r.GetSource(),
		Type:        r.GetType(),
		Subject:     r.GetSubject(),
		Attributes:  r.GetAttributes(),
		Data:        r.GetData(),
		Time:        timestamppb.New(time.Now()),
	}
	if toSend.Attributes == nil {
		data, _ := structpb.NewStruct(nil)
		toSend.Attributes = data
	}
	claims, _ := structpb.NewStruct(c.Claims)
	toSend.Attributes.Fields["auth.claims"] = structpb.NewStructValue(claims)
	bits, err := proto.Marshal(toSend)
	if err != nil {
		return nil, err
	}
	if err := s.conn.Publish(getSubject(r.GetType(), r.GetSubject()), bits); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Service) Request(ctx context.Context, r *cloudEventsProxy.CloudEventInput) (*cloudEventsProxy.CloudEvent, error) {
	c, ok := auth.GetContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}
	toSend := &cloudEventsProxy.CloudEvent{
		Specversion: r.GetSpecversion(),
		Id:          uuid.New().String(),
		Source:      r.GetSource(),
		Type:        r.GetType(),
		Subject:     r.GetSubject(),
		Attributes:  r.GetAttributes(),
		Data:        r.GetData(),
		Time:        timestamppb.New(time.Now()),
	}
	if toSend.Attributes == nil {
		data, _ := structpb.NewStruct(nil)
		toSend.Attributes = data
	}
	claims, _ := structpb.NewStruct(c.Claims)
	toSend.Attributes.Fields["auth.claims"] = structpb.NewStructValue(claims)

	bits, err := proto.Marshal(toSend)
	if err != nil {
		return nil, err
	}
	resp, err := s.conn.Request(getSubject(r.GetType(), r.GetSubject()), bits, 30*time.Second)
	if err != nil {
		return nil, err
	}
	var event cloudEventsProxy.CloudEvent
	if err := proto.Unmarshal(resp.Data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (s *Service) Receive(r *cloudEventsProxy.ReceiveRequest, server cloudEventsProxy.CloudEventsService_ReceiveServer) error {
	_, ok := auth.GetContext(server.Context())
	if !ok {
		return status.Error(codes.Unauthenticated, "unauthenticated")
	}
	var (
		err error
		sub *nats.Subscription
		wg  = sync.WaitGroup{}
	)
	sub, err = s.conn.Subscribe(getSubject(r.GetType(), r.GetSubject()), func(msg *nats.Msg) {
		wg.Add(1)
		defer wg.Done()
		var event cloudEventsProxy.CloudEvent
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
		if err := sub.Unsubscribe(); err != nil {
			return err
		}
		wg.Wait()
		return nil
	}
}

func getSubject(typ string, subject string) string {
	if subject != "" {
		return fmt.Sprintf("%s/%s", typ, subject)
	}
	return typ
}
