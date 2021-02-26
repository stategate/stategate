package channel

import (
	"context"
	"crypto/tls"
	"fmt"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/channel/inmem"
	"github.com/autom8ter/stategate/internal/channel/kafka"
	nats2 "github.com/autom8ter/stategate/internal/channel/nats"
	"github.com/autom8ter/stategate/internal/channel/redis"
	"github.com/autom8ter/stategate/internal/channel/stan"
	"github.com/autom8ter/stategate/internal/constants"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/autom8ter/stategate/internal/logger"
	rediss "github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nuid"
	stan2 "github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	kafkaa "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

type Provider interface {
	PublishEvent(ctx context.Context, event *stategate.Event) *errorz.Error
	GetEventChannel(ctx context.Context) (chan *stategate.Event, *errorz.Error)
	PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error
	GetMessageChannel(ctx context.Context) (chan *stategate.PeerMessage, *errorz.Error)
	Close() error
}

type ProviderName string

const (
	NATS  ProviderName = "nats"
	STAN  ProviderName = "stan"
	INMEM ProviderName = "inmem"
	REDIS ProviderName = "redis"
	KAFKA ProviderName = "kafka"
)

var AllProviderNames = []ProviderName{INMEM, REDIS, NATS, STAN, KAFKA}

func GetChannelProvider(lgger *logger.Logger, providerConfig map[string]string) (Provider, error) {
	if providerConfig == nil {
		return nil, errors.Errorf("empty backend channel provider config")
	}
	name := providerConfig["name"]
	if name == "" {
		return nil, errors.New("storage provider: empty name")
	}
	var tlsConfig *tls.Config
	if providerConfig["client_cert_file"] != "" && providerConfig["client_key_file"] != "" {
		cer, err := tls.LoadX509KeyPair(providerConfig["tls_cert"], providerConfig["tls_key"])
		if err != nil {
			lgger.Error("failed to load tls config", zap.Error(err))
			return nil, err
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cer},
		}
	}
	switch ProviderName(name) {
	case INMEM:
		svc := inmem.NewService(lgger)
		return svc, nil
	case KAFKA:
		kafkaAddr := providerConfig["addr"]
		if kafkaAddr == "" {
			return nil, errors.New("kafka config: empty addr")
		}

		dialer := &kafkaa.Dialer{
			Timeout:   10 * time.Second,
			DualStack: true,
		}
		if tlsConfig != nil {
			dialer.TLS = tlsConfig
		}
		re := kafkaa.NewReader(kafkaa.ReaderConfig{
			Brokers: []string{kafkaAddr},
			Topic:   constants.EventChannel,
			Dialer:  dialer,
		})
		rm := kafkaa.NewReader(kafkaa.ReaderConfig{
			Brokers: []string{kafkaAddr},
			Topic:   constants.MessageChannel,
			Dialer:  dialer,
		})
		we := kafkaa.NewWriter(kafkaa.WriterConfig{
			Brokers: []string{kafkaAddr},
			Topic:   constants.EventChannel,
			Dialer:  dialer,
		})
		wm := kafkaa.NewWriter(kafkaa.WriterConfig{
			Brokers: []string{kafkaAddr},
			Topic:   constants.MessageChannel,
			Dialer:  dialer,
		})
		svc, err := kafka.NewService(lgger, re, rm, we, wm)
		if err != nil {
			return nil, err
		}
		return svc, nil
	case REDIS:
		redisHost := providerConfig["addr"]
		if redisHost == "" {
			return nil, errors.New("redis config: empty addr")
		}
		opts := &rediss.Options{
			Addr:     redisHost,
			Username: providerConfig["user"],
			Password: providerConfig["password"],
		}
		if tlsConfig != nil {
			opts.TLSConfig = tlsConfig
		}
		client := rediss.NewClient(opts)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if _, err := client.Ping(ctx).Result(); err != nil {
			return nil, err
		}
		return redis.NewService(lgger, client), nil
	case STAN:
		cluster := providerConfig["cluster"]
		if cluster == "" {
			return nil, errors.New("stan config: empty cluster")
		}
		natsUrl := providerConfig["addr"]
		if natsUrl == "" {
			return nil, errors.New("nats config: empty addr")
		}
		hostname, err := os.Hostname()
		if err != nil {
			return nil, err
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
				opts = append(opts, nats.ClientCert(providerConfig["client_cert_file"], providerConfig["client_key_file"]))
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
			return nil, err
		}
		sconn, err := stan2.Connect(cluster, clientID, stan2.NatsConn(conn))
		svc, err := stan.NewService(lgger, sconn)
		if err != nil {
			return nil, err
		}
		return svc, nil
	case NATS:
		natsUrl := providerConfig["addr"]
		if natsUrl == "" {
			return nil, errors.New("nats config: empty addr")
		}
		hostname, err := os.Hostname()
		if err != nil {
			return nil, err
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
				opts = append(opts, nats.ClientCert(providerConfig["client_cert_file"], providerConfig["client_key_file"]))
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
			return nil, err
		}
		svc, err := nats2.NewService(lgger, conn)
		if err != nil {
			return nil, err
		}
		return svc, nil
	default:
		return nil, errors.Errorf("unsupported backend provider: %s. must be one of: %v", name, AllProviderNames)
	}
}
