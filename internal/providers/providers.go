package providers

import (
	"context"
	"crypto/tls"
	"fmt"
	rediss "github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nuid"
	"github.com/pkg/errors"
	"github.com/stategate/stategate/internal/api"
	"github.com/stategate/stategate/internal/logger"
	"github.com/stategate/stategate/internal/providers/amqpb"
	"github.com/stategate/stategate/internal/providers/mongo"
	natss "github.com/stategate/stategate/internal/providers/nats"
	"github.com/stategate/stategate/internal/providers/redis"
	"github.com/streadway/amqp"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

type Name string

const (
	REDIS Name = "redis"
	MONGO Name = "mongo"
	NATS  Name = "nats"
	AMQP  Name = "amqp"
)

var (
	AllCacheProviders   = []Name{REDIS}
	AllStorageProviders = []Name{MONGO}
	AllChannelProviders = []Name{REDIS, NATS, AMQP}
)

func GetStorageProvider(lgger *logger.Logger, providerConfig map[string]string) (api.StorageProvider, error) {
	if providerConfig == nil {
		return nil, errors.Errorf("empty backend storage provider config")
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
	switch Name(name) {
	case MONGO:
		db := providerConfig["database"]
		if db == "" {
			return nil, errors.New("mongo config: empty database")
		}
		uri := providerConfig["addr"]
		if uri == "" {
			return nil, errors.New("mongo config: empty addr")
		}
		opts := options.Client()
		opts.ApplyURI(uri)
		if tlsConfig != nil {
			opts.TLSConfig = tlsConfig
		}
		client, err := mongo2.NewClient(opts)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			return nil, err
		}
		return mongo.NewProvider(client), nil
	default:
		return nil, errors.Errorf("unsupported backend provider: %s. must be one of: %v", name, AllStorageProviders)
	}
}

func GetCacheProvider(lgger *logger.Logger, providerConfig map[string]string) (api.CacheProvider, error) {
	if providerConfig == nil {
		return nil, errors.Errorf("empty backend channel provider config")
	}
	name := providerConfig["name"]
	if name == "" {
		return nil, errors.New("cache provider: empty name")
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
	switch Name(name) {
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
	default:
		return nil, errors.Errorf("unsupported cache provider: %s. must be one of: %v", name, AllCacheProviders)
	}
}

func GetChannelProvider(lgger *logger.Logger, providerConfig map[string]string) (api.ChannelProvider, error) {
	if providerConfig == nil {
		return nil, errors.Errorf("empty channel provider config")
	}
	name := providerConfig["name"]
	if name == "" {
		return nil, errors.New("cache provider: empty name")
	}
	addr := providerConfig["addr"]
	user := providerConfig["user"]
	pw := providerConfig["password"]
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

	switch Name(name) {
	case AMQP:
		if addr == "" {
			return nil, errors.New("amqp config: empty addr")
		}
		conn, err := amqp.Dial(addr)
		if err != nil {
			return nil, errors.Wrap(err, "failed to dial amqp")
		}
		return amqpb.NewService(lgger, conn)
	case NATS:
		if addr == "" {
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
			if user != "" {
				opts = append(opts, nats.UserInfo(user, pw))
			}
			conn, err = nats.Connect(
				addr,
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
		return natss.NewService(lgger, conn)
	case REDIS:
		if addr == "" {
			return nil, errors.New("redis config: empty addr")
		}
		opts := &rediss.Options{
			Addr:     addr,
			Username: user,
			Password: pw,
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
	default:
		return nil, errors.Errorf("unsupported channel provider: %s. must be one of: %v", name, AllChannelProviders)
	}
}
