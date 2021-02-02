package service

import (
	"context"
	cloudEventsProxy "github.com/autom8ter/cloudEventsProxy/gen/grpc/go"
	"github.com/autom8ter/cloudEventsProxy/internal/logger"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
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

func (s *Service) Send(ctx context.Context, c *cloudEventsProxy.CloudEventInput) (*empty.Empty, error) {
	bits, err := proto.Marshal(c)
	if err != nil {
		return nil, err
	}
	if err := s.conn.Publish(c.GetType(), bits); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Service) Request(ctx context.Context, c *cloudEventsProxy.CloudEventInput) (*cloudEventsProxy.CloudEvent, error) {

	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) Receive(r *cloudEventsProxy.ReceiveRequest, server cloudEventsProxy.CloudEventsService_ReceiveServer) error {
	var (
		err error
		sub *nats.Subscription
		wg  = sync.WaitGroup{}
	)
	sub, err = s.conn.Subscribe(r.GetType(), func(msg *nats.Msg) {
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
	for {
		select {
		case <-ctx.Done():
			if err := sub.Unsubscribe(); err != nil {
				return err
			}
			wg.Wait()
			return nil
		}
	}
}
