package stategate_client_go

import (
	"context"
	"golang.org/x/sync/errgroup"
)

// ClientSet holds an EntityClient, EventClient, and PeerClient
type ClientSet struct {
	entity *EntityClient
	event  *EventClient
	peer   *PeerClient
}

// NewClientSet returns an initialized ClientSet
func NewClientSet(ctx context.Context, target string, opts ...Opt) (*ClientSet, error) {
	entity, err := NewEntityClient(ctx, target, opts...)
	if err != nil {
		return nil, err
	}
	event, err := NewEventClient(ctx, target, opts...)
	if err != nil {
		return nil, err
	}
	peer, err := NewPeerClient(ctx, target, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientSet{
		entity: entity,
		event:  event,
		peer:   peer,
	}, nil
}

// Entity returns the clientset's EntityService client
func (c *ClientSet) Entity() *EntityClient {
	return c.entity
}

// Event returns the clientset's EventService client
func (c *ClientSet) Event() *EventClient {
	return c.event
}

// Peer returns the clientset's PeerService client
func (c *ClientSet) Peer() *PeerClient {
	return c.peer
}

// Peer returns the clientset's PeerService client
func (c *ClientSet) Close() error {
	group := &errgroup.Group{}
	group.Go(c.Event().Close)
	group.Go(c.Peer().Close)
	group.Go(c.Entity().Close)
	return group.Wait()
}
