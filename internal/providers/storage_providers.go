package providers

import (
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/storage"
	"github.com/autom8ter/eventgate/internal/storage/elastic"
	"github.com/autom8ter/eventgate/internal/storage/inmem"
	"github.com/autom8ter/eventgate/internal/storage/mongo"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
)

type StorageProvider string

const (
	MONGO_STORAGE         StorageProvider = "mongo"
	ELASTICSEARCH_STORAGE StorageProvider = "elasticsearch"
	INMEM_STORAGE         StorageProvider = "inmem"
)

var AllStorageProviders = []StorageProvider{INMEM_STORAGE, MONGO_STORAGE, ELASTICSEARCH_STORAGE}

func GetStorageProvider(provider StorageProvider, lgger *logger.Logger, providerConfig map[string]string) (storage.Provider, error) {
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
			Username:              providerConfig["user"],
			Password:              providerConfig["password"],
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
		uri := providerConfig["uri"]
		if uri == "" {
			return nil, errors.New("mongo config: empty uri")
		}
		svc, err := mongo.NewService(db, uri, lgger)
		if err != nil {
			return nil, err
		}
		return svc, nil
	default:
		return nil, errors.Errorf("unsupported backend provider: %s. must be one of: %v", provider, AllStorageProviders)
	}
}
