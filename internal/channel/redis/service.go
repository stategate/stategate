package redis

import (
	"context"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/constants"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

type Service struct {
	eventsChan string
	logger     *logger.Logger
	conn       *redis.Client
}

func NewService(logger *logger.Logger, conn *redis.Client) *Service {
	return &Service{eventsChan: constants.BackendChannel, logger: logger, conn: conn}
}

func (s *Service) Publish(ctx context.Context, event *stategate.Event) error {
	bits, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	if _, err := s.conn.Publish(ctx, s.eventsChan, bits).Result(); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetChannel(ctx context.Context) (chan *stategate.Event, error) {
	sub := s.conn.Subscribe(context.Background(), s.eventsChan)
	events := make(chan *stategate.Event)
	go func() {
		ch := sub.Channel()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				sub.Close()
				return
			case msg := <-ch:
				var event stategate.Event
				if err := proto.Unmarshal([]byte(msg.Payload), &event); err != nil {
					s.logger.Error("failed to unmarshal event", zap.Error(err))
					return
				}
				events <- &event
			}
		}
	}()
	return events, nil
}

func (s *Service) Close() error {
	return nil
}
