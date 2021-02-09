package inmem

import (
	"context"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"sync"
)

type Service struct {
	mu     sync.RWMutex
	events *eventgate.Events
}

func NewService() *Service {
	return &Service{
		mu:     sync.RWMutex{},
		events: &eventgate.Events{},
	}
}

func (s *Service) SaveEvent(ctx context.Context, event *eventgate.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events.Events = append(s.events.Events, event)
	return nil
}

func (s *Service) GetEvents(ctx context.Context, opts *eventgate.HistoryOpts) (*eventgate.Events, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	events := &eventgate.Events{}
	skip := int64(0)
	for _, e := range s.events.GetEvents() {
		if e.GetTime().AsTime().After(opts.GetMin().AsTime()) || e.GetTime().AsTime().Before(opts.GetMax().AsTime()) {
			if opts.GetOffset() > 0 && skip < opts.GetOffset() {
				skip += 1
				continue
			}
			events.Events = append(events.Events, e)
		}
		if opts.GetLimit() > 0 && len(events.Events) >= int(opts.GetLimit()) {
			return events, nil
		}
	}
	return events, nil
}

func (s *Service) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = &eventgate.Events{}
	return nil
}
