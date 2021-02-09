package sqs

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/constants"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/storage"
	"github.com/autom8ter/machine/pubsub"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Service struct {
	logger   *logger.Logger
	conn     *sqs.SQS
	ps       pubsub.PubSub
	cancel   func()
	queueUrl *string
	storage  storage.Provider
}

func NewService(logger *logger.Logger, sess *session.Session, storage storage.Provider) (*Service, error) {
	s := &Service{
		logger:  logger,
		conn:    sqs.New(sess),
		ps:      pubsub.NewPubSub(),
		storage: storage,
	}
	queueName := constants.BackendChannel
	s.conn.CreateQueue(&sqs.CreateQueueInput{
		Attributes: nil,
		QueueName:  &queueName,
		Tags:       nil,
	})
	queueUrl, err := s.conn.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		return nil, err
	}
	s.queueUrl = queueUrl.QueueUrl
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				resp, err := s.conn.ReceiveMessage(&sqs.ReceiveMessageInput{
					AttributeNames: []*string{
						aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
					},
					MessageAttributeNames: []*string{
						aws.String(sqs.QueueAttributeNameAll),
					},
					QueueUrl:            s.queueUrl,
					MaxNumberOfMessages: aws.Int64(1),
				})
				if err != nil {
					s.logger.Error("failed to receive event", zap.Error(err))
					continue
				}
				if len(resp.Messages) > 0 && resp.Messages[0].Body != nil {
					var event eventgate.Event

					var body = *resp.Messages[0].Body
					if err := proto.Unmarshal([]byte(body), &event); err != nil {
						s.logger.Error("failed to unmarshal event", zap.Error(err))
						continue
					}
					if err := s.ps.Publish(event.GetChannel(), &event); err != nil {
						s.logger.Error("failed to publish event", zap.Error(err))
					}
				}
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

	group := errgroup.Group{}
	group.Go(func() error {
		_, err = s.conn.SendMessage(&sqs.SendMessageInput{
			MessageBody: aws.String(string(bits)),
			QueueUrl:    s.queueUrl,
		})
		if err != nil {
			return err
		}
		return nil
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
	s.cancel()
	s.ps.Close()
	return nil
}

func (s *Service) History(ctx context.Context, opts *eventgate.HistoryOpts) (*eventgate.Events, error) {
	if s.storage == nil {
		return nil, status.Error(codes.Unimplemented, "backend timeseries storage provider not registered")
	}
	return s.storage.GetEvents(ctx, opts)
}
