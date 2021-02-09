package storage

import (
	"context"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
)

type Provider interface {
	SaveEvent(ctx context.Context, event *eventgate.Event) error
	GetEvents(ctx context.Context, opts *eventgate.HistoryOpts) (*eventgate.Events, error)
	Close() error
}
