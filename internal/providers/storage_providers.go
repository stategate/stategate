package providers

import (
	"context"
	"crypto/tls"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/storage"
	"github.com/autom8ter/eventgate/internal/storage/elastic"
	"github.com/autom8ter/eventgate/internal/storage/inmem"
	"github.com/autom8ter/eventgate/internal/storage/mongo"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

type StorageProvider string

const (
	MONGO_STORAGE         StorageProvider = "mongo"
	ELASTICSEARCH_STORAGE StorageProvider = "elasticsearch"
	INMEM_STORAGE         StorageProvider = "inmem"
)

var AllStorageProviders = []StorageProvider{INMEM_STORAGE, MONGO_STORAGE, ELASTICSEARCH_STORAGE}

func GetStorageProvider(provider StorageProvider, lgger *logger.Logger, providerConfig map[string]string) (storage.Provider, error) {
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
	case INMEM_STORAGE:
		return inmem.NewService(), nil
	case ELASTICSEARCH_STORAGE:
		addr := providerConfig["addr"]
		if addr == "" {
			return nil, errors.New("elasticsearch config: empty addr")
		}
		cfg := elasticsearch.Config{
			Addresses: []string{
				addr,
			},
			Username: providerConfig["user"],
			Password: providerConfig["password"],
		}
		client, err := elasticsearch.NewClient(cfg)
		if err != nil {
			return nil, err
		}
		return elastic.NewService(client)
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
		svc, err := mongo.NewService(db, client, lgger)
		if err != nil {
			return nil, err
		}
		return svc, nil
	default:
		return nil, errors.Errorf("unsupported backend provider: %s. must be one of: %v", provider, AllStorageProviders)
	}
}
