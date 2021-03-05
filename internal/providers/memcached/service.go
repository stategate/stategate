package memcached

import (
	"context"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	stategate "github.com/stategate/stategate/gen/grpc/go"
	"github.com/stategate/stategate/internal/errorz"
	"github.com/stategate/stategate/internal/logger"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	logger *logger.Logger
	conn   *memcache.Client
}

func NewService(logger *logger.Logger, conn *memcache.Client) *Service {
	return &Service{logger: logger, conn: conn}
}

func (s Service) Get(ctx context.Context, ref *stategate.CacheRef) (*stategate.Cache, *errorz.Error) {
	res, err := s.conn.Get(cachedKeyName(ref.GetDomain(), ref.GetKey()))
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return nil, &errorz.Error{
				Type: errorz.ErrNotFound,
				Info: "failed to get cached value",
				Err:  err,
				Metadata: map[string]string{
					"cache_key":    ref.GetKey(),
					"cache_domain": ref.GetDomain(),
				},
			}
		}
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to get cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    ref.GetKey(),
				"cache_domain": ref.GetDomain(),
			},
		}
	}
	var cachedVal stategate.Cache
	if err := proto.Unmarshal(res.Value, &cachedVal); err != nil {
		return nil, &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to get cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    ref.GetKey(),
				"cache_domain": ref.GetDomain(),
			},
		}
	}
	return &cachedVal, nil
}

func (s Service) Set(ctx context.Context, value *stategate.Cache) *errorz.Error {
	bits, err := proto.Marshal(value)
	if err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    value.GetKey(),
				"cache_domain": value.GetDomain(),
			},
		}
	}
	exp := int32(0)
	if value.GetExp() != nil {
		exp = int32(value.GetExp().AsTime().Unix())
	}
	if err := s.conn.Set(&memcache.Item{
		Key:        cachedKeyName(value.GetDomain(), value.GetKey()),
		Value:      bits,
		Expiration: exp,
	}); err != nil {
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to set cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    value.GetKey(),
				"cache_domain": value.GetDomain(),
			},
		}
	}
	return nil
}

func (s Service) Del(ctx context.Context, ref *stategate.CacheRef) *errorz.Error {
	if err := s.conn.Delete(cachedKeyName(ref.GetDomain(), ref.GetKey())); err != nil {
		if err == memcache.ErrCacheMiss {
			return nil
		}
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to delete cached value",
			Err:  err,
			Metadata: map[string]string{
				"cache_key":    ref.GetKey(),
				"cache_domain": ref.GetDomain(),
			},
		}
	}
	return nil
}

func (s Service) Lock(ctx context.Context, ref *stategate.Mutex) *errorz.Error {
	exp := int32(0)
	if ref.GetExp() != nil {
		exp = int32(ref.GetExp().AsTime().Unix())
	}
	if err := s.conn.Add(&memcache.Item{
		Key:        cachedLockName(ref.GetDomain(), ref.GetKey()),
		Value:      []byte("true"),
		Expiration: exp,
	}); err != nil {
		if err == memcache.ErrNotStored {
			return &errorz.Error{
				Type: errorz.ErrLocked,
				Info: "failed to lock mutex",
				Err:  err,
				Metadata: map[string]string{
					"mutex_key":    ref.GetKey(),
					"mutex_domain": ref.GetDomain(),
				},
			}
		}
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to lock mutex",
			Err:  err,
			Metadata: map[string]string{
				"mutex_key":    ref.GetKey(),
				"mutex_domain": ref.GetDomain(),
			},
		}
	}
	return nil
}

func (s Service) Unlock(ctx context.Context, ref *stategate.MutexRef) *errorz.Error {
	if err := s.conn.Delete(cachedLockName(ref.GetDomain(), ref.GetKey())); err != nil {
		if err == memcache.ErrCacheMiss {
			return nil
		}
		return &errorz.Error{
			Type: errorz.ErrUnknown,
			Info: "failed to unlock mutex",
			Err:  err,
			Metadata: map[string]string{
				"mutex_key":    ref.GetKey(),
				"mutex_domain": ref.GetDomain(),
			},
		}
	}
	return nil
}

func (s Service) Close() error {
	return nil
}

func cachedKeyName(domain string, key string) string {
	return fmt.Sprintf("%s.%s", domain, key)
}

func cachedLockName(domain string, key string) string {
	return fmt.Sprintf("mutex.%s.%s", domain, key)
}
