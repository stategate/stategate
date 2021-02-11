package storage

import (
	"context"
	"crypto/tls"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/storage/mongo"
	"github.com/pkg/errors"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

type Provider interface {
	SetObject(ctx context.Context, object *eventgate.Object) error
	SaveEvent(ctx context.Context, event *eventgate.Event) error
	GetObject(ctx context.Context, ref *eventgate.ObjectRef) (*eventgate.Object, error)
	SearchEvents(ctx context.Context, ref *eventgate.SearchOpts) (*eventgate.Events, error)
	Close() error
}

type ProviderName string

const (
	MONGO_STORAGE ProviderName = "mongo"
)

var AllProviderNames = []ProviderName{MONGO_STORAGE}

func GetStorageProvider(provider ProviderName, lgger *logger.Logger, providerConfig map[string]string) (Provider, error) {
	if providerConfig == nil {
		return nil, errors.Errorf("empty backend storage provider config")
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
	switch provider {
	case MONGO_STORAGE:
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
		return mongo.NewProvider(client.Database(db)), nil
	default:
		return nil, errors.Errorf("unsupported backend provider: %s. must be one of: %v", provider, AllProviderNames)
	}
}
