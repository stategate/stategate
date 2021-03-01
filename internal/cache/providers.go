package cache

import (
	"context"
	"crypto/tls"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/cache/redis"
	"github.com/autom8ter/stategate/internal/errorz"
	"github.com/autom8ter/stategate/internal/logger"
	rediss "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

// ChannelProvider acts as dependency injection for broadcasting messages to stategate instances as they fan out
type ChannelProvider interface {
	PublishEvent(ctx context.Context, event *stategate.Event) *errorz.Error
	GetEventChannel(ctx context.Context) (chan *stategate.Event, *errorz.Error)
	PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error
	GetMessageChannel(ctx context.Context) (chan *stategate.PeerMessage, *errorz.Error)
}

// CacheProvider acts as dependency injection for caching ephemeral data
type CacheProvider interface {
	Get(ctx context.Context, ref *stategate.CacheRef) (*stategate.Cache, *errorz.Error)
	Set(ctx context.Context, value *stategate.Cache) *errorz.Error
	Del(ctx context.Context, value *stategate.CacheRef) *errorz.Error
}

// MutexProvider acts as dependency injection for distributed mutex operations
type MutexProvider interface {
	Lock(ctx context.Context, ref *stategate.Mutex) *errorz.Error
	Unlock(ctx context.Context, value *stategate.MutexRef) *errorz.Error
}

// Provider is a channel, cache, & mutex provider
type Provider interface {
	CacheProvider
	ChannelProvider
	MutexProvider
	Close() error
}

type ProviderName string

const (
	REDIS ProviderName = "redis"
)

var AllProviderNames = []ProviderName{REDIS}

func GetCacheProvider(lgger *logger.Logger, providerConfig map[string]string) (Provider, error) {
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
	switch ProviderName(name) {
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
		return nil, errors.Errorf("unsupported cache provider: %s. must be one of: %v", name, AllProviderNames)
	}
}
