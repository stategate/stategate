package api

import (
	"context"
	stategate "github.com/stategate/stategate/gen/grpc/go"
	"github.com/stategate/stategate/internal/errorz"
)

// ChannelProvider is an interface for broadcasting messages to stategate instances as they fan out horizontally
type ChannelProvider interface {
	PublishEvent(ctx context.Context, event *stategate.Event) *errorz.Error
	GetEventChannel(ctx context.Context) (chan *stategate.Event, *errorz.Error)
	PublishMessage(ctx context.Context, message *stategate.PeerMessage) *errorz.Error
	GetMessageChannel(ctx context.Context) (chan *stategate.PeerMessage, *errorz.Error)
	Close() error
}

// CacheProvider is an interface for caching ephemeral data
type CacheProvider interface {
	Get(ctx context.Context, ref *stategate.CacheRef) (*stategate.Cache, *errorz.Error)
	Set(ctx context.Context, value *stategate.Cache) *errorz.Error
	Del(ctx context.Context, value *stategate.CacheRef) *errorz.Error
	Lock(ctx context.Context, ref *stategate.Mutex) *errorz.Error
	Unlock(ctx context.Context, value *stategate.MutexRef) *errorz.Error
	Close() error
}

// StorageProvider is an interface for querying/persisting entities & events
type StorageProvider interface {
	SetEntity(ctx context.Context, entity *stategate.Entity) *errorz.Error
	EditEntity(ctx context.Context, entity *stategate.Entity) (*stategate.Entity, *errorz.Error)
	SearchEntities(ctx context.Context, ref *stategate.SearchEntityOpts) (*stategate.Entities, *errorz.Error)
	DelEntity(ctx context.Context, ref *stategate.EntityRef) *errorz.Error
	GetEntity(ctx context.Context, ref *stategate.EntityRef) (*stategate.Entity, *errorz.Error)
	SaveEvent(ctx context.Context, event *stategate.Event) *errorz.Error
	SearchEvents(ctx context.Context, ref *stategate.SearchEventOpts) (*stategate.Events, *errorz.Error)
	GetEvent(ctx context.Context, ref *stategate.EventRef) (*stategate.Event, *errorz.Error)
	Close() error
}
