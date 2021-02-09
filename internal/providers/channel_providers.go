package providers

import (
	gsub "cloud.google.com/go/pubsub"
	"context"
	"crypto/tls"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/channel/inmem"
	"github.com/autom8ter/eventgate/internal/channel/kafka"
	nats2 "github.com/autom8ter/eventgate/internal/channel/nats"
	"github.com/autom8ter/eventgate/internal/channel/pubsub"
	"github.com/autom8ter/eventgate/internal/channel/redis"
	"github.com/autom8ter/eventgate/internal/channel/sqs"
	"github.com/autom8ter/eventgate/internal/channel/stan"
	"github.com/autom8ter/eventgate/internal/constants"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/storage"
	"github.com/aws/aws-sdk-go/aws/session"
	rediss "github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nuid"
	stan2 "github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	kafkaa "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"google.golang.org/api/option"
	"os"
	"strings"
	"time"
)

type ChannelProvider string

const (
	NATS         ChannelProvider = "nats"
	STAN         ChannelProvider = "stan"
	INMEM        ChannelProvider = "inmem"
	REDIS        ChannelProvider = "redis"
	KAFKA        ChannelProvider = "kafka"
	GOOGLEPUBSUB ChannelProvider = "google-pubsub"
	AWSSQS       ChannelProvider = "aws-sqs"
)

var AllChannelProviders = []ChannelProvider{INMEM, REDIS, NATS, STAN, KAFKA, GOOGLEPUBSUB, AWSSQS}

func GetChannelProvider(provider ChannelProvider, storage storage.Provider, lgger *logger.Logger, providerConfig map[string]string) (eventgate.EventGateServiceServer, func(), error) {
	if providerConfig == nil {
		return nil, nil, errors.Errorf("empty backend config")
	}
	var tlsConfig *tls.Config
	if providerConfig["tls_cert_file"] != "" && providerConfig["tls_key_file"] != "" {
		cer, err := tls.LoadX509KeyPair(providerConfig["tls_cert"], providerConfig["tls_key"])
		if err != nil {
			lgger.Error("failed to load tls config", zap.Error(err))
			return nil, nil, err
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cer},
		}
	}
	switch provider {
	case INMEM:
		svc := inmem.NewService(lgger, storage)
		return svc, func() {
			if err := svc.Close(); err != nil {
				lgger.Error("failed to close backend", zap.Error(err))
			}
		}, nil
	case KAFKA:
		kafkaAddr := providerConfig["addr"]
		if kafkaAddr == "" {
			return nil, nil, errors.New("kafka config: empty addr")
		}

		dialer := &kafkaa.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
		}
		if tlsConfig != nil {
			dialer.TLS = tlsConfig
		}
		r := kafkaa.NewReader(kafkaa.ReaderConfig{
			Brokers: []string{kafkaAddr},
			Topic:   constants.BackendChannel,
			Dialer:  dialer,
		})
		w := kafkaa.NewWriter(kafkaa.WriterConfig{
			Brokers: []string{kafkaAddr},
			Topic:   constants.BackendChannel,
			Dialer:  dialer,
		})
		svc, err := kafka.NewService(lgger, r, w, storage)
		if err != nil {
			return nil, nil, err
		}
		return svc, func() {
			if err := svc.Close(); err != nil {
				lgger.Error("failed to close backend", zap.Error(err))
			}
		}, nil
	case REDIS:
		redisHost := providerConfig["addr"]
		if redisHost == "" {
			return nil, nil, errors.New("redis config: empty addr")
		}
		opts := &rediss.Options{
			Addr:     redisHost,
			Username: providerConfig["username"],
			Password: providerConfig["password"],
		}
		if tlsConfig != nil {
			opts.TLSConfig = tlsConfig
		}
		client := rediss.NewClient(opts)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if _, err := client.Ping(ctx).Result(); err != nil {
			return nil, nil, err
		}
		svc, err := redis.NewService(lgger, client, storage)
		if err != nil {
			return nil, nil, err
		}
		return svc, func() {
			if err := svc.Close(); err != nil {
				lgger.Error("failed to close backend", zap.Error(err))
			}
		}, nil
	case STAN:
		cluster := providerConfig["cluster"]
		if cluster == "" {
			return nil, nil, errors.New("stan config: empty cluster")
		}
		natsUrl := providerConfig["addr"]
		if natsUrl == "" {
			return nil, nil, errors.New("nats config: empty addr")
		}
		hostname, err := os.Hostname()
		if err != nil {
			return nil, nil, err
		}
		hostname = strings.NewReplacer(".", "_").Replace(hostname)
		var (
			conn     *nats.Conn
			clientID string
		)
		for i := 0; i < 10; i++ {
			var (
				clientID = fmt.Sprintf("%s-%s", hostname, nuid.Next())
				opts     = []nats.Option{nats.Name(clientID)}
			)
			if tlsConfig != nil {
				opts = append(opts, nats.ClientCert(providerConfig["tls_cert_file"], providerConfig["tls_key_file"]))
			}

			conn, err = nats.Connect(
				natsUrl,
				opts...)

			if err == nil && conn != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		if err != nil {
			return nil, nil, err
		}
		sconn, err := stan2.Connect(cluster, clientID, stan2.NatsConn(conn))
		svc, err := stan.NewService(lgger, sconn, storage)
		if err != nil {
			return nil, nil, err
		}
		return svc, func() {
			if err := svc.Close(); err != nil {
				lgger.Error("failed to close backend", zap.Error(err))
			}
		}, nil
	case NATS:
		natsUrl := providerConfig["addr"]
		if natsUrl == "" {
			return nil, nil, errors.New("nats config: empty addr")
		}
		hostname, err := os.Hostname()
		if err != nil {
			return nil, nil, err
		}
		hostname = strings.NewReplacer(".", "_").Replace(hostname)
		var (
			conn *nats.Conn
		)
		for i := 0; i < 10; i++ {
			var (
				clientID = fmt.Sprintf("%s-%s", hostname, nuid.Next())
				opts     = []nats.Option{nats.Name(clientID)}
			)

			if tlsConfig != nil {
				opts = append(opts, nats.ClientCert(providerConfig["tls_cert_file"], providerConfig["tls_key_file"]))
			}
			conn, err = nats.Connect(
				natsUrl,
				opts...,
			)
			if err == nil && conn != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		if err != nil {
			return nil, nil, err
		}
		svc, err := nats2.NewService(lgger, conn, storage)
		if err != nil {
			return nil, nil, err
		}
		return svc, func() {
			if err := svc.Close(); err != nil {
				lgger.Error("failed to close backend", zap.Error(err))
			}
		}, nil
	case GOOGLEPUBSUB:
		ctx := context.Background()
		project := providerConfig["project"]
		if project == "" {
			return nil, nil, errors.New("google pubsub config: empty project")
		}
		credentialsFile := providerConfig["credentials_file"]
		var (
			client *gsub.Client
			err    error
		)
		if credentialsFile != "" {
			client, err = gsub.NewClient(ctx, project, option.WithCredentialsFile(credentialsFile))
			if err != nil {
				return nil, nil, err
			}
		} else {
			client, err = gsub.NewClient(ctx, project)
			if err != nil {
				return nil, nil, err
			}
		}
		client.CreateTopic(ctx, constants.BackendChannel)
		svc, err := pubsub.NewService(lgger, client, storage)
		if err != nil {
			return nil, nil, err
		}
		return svc, func() {
			if err := svc.Close(); err != nil {
				lgger.Error("failed to close backend", zap.Error(err))
			}
		}, nil
	case AWSSQS:
		sess, err := session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})
		if err != nil {
			return nil, nil, err
		}
		svc, err := sqs.NewService(lgger, sess, storage)
		if err != nil {
			return nil, nil, err
		}
		return svc, func() {
			if err := svc.Close(); err != nil {
				lgger.Error("failed to close backend", zap.Error(err))
			}
		}, nil
	default:
		return nil, nil, errors.Errorf("unsupported backend provider: %s. must be one of: %v", provider, AllChannelProviders)
	}
}
