package stategate_client_go

import (
	"context"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"google.golang.org/grpc"
)

// ClientSet holds an EntityClient, EventClient, and PeerClient
type ClientSet struct {
	conn   *grpc.ClientConn
	entity *EntityClient
	event  *EventClient
	peer   *PeerClient
	cache  *CacheClient
	mutex  *MutexClient
}

// NewClientSet returns an initialized ClientSet
func NewClientSet(ctx context.Context, target string, opts ...Opt) (*ClientSet, error) {
	conn, err := getConnection(ctx, target, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientSet{
		entity: &EntityClient{
			client: stategate.NewEntityServiceClient(conn),
			conn:   conn,
		},
		event: &EventClient{
			client: stategate.NewEventServiceClient(conn),
			conn:   conn,
		},
		peer: &PeerClient{
			client: stategate.NewPeerServiceClient(conn),
			conn:   conn,
		},
		cache: &CacheClient{
			client: stategate.NewCacheServiceClient(conn),
			conn:   conn,
		},
		mutex: &MutexClient{
			client: stategate.NewMutexServiceClient(conn),
			conn:   conn,
		},
		conn: conn,
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

// Cache returns the clientset's CacheService client
func (c *ClientSet) Cache() *CacheClient {
	return c.cache
}

// Mutex returns the clientset's MutexService client
func (c *ClientSet) Mutex() *MutexClient {
	return c.mutex
}

// Close closes the underlying connection
func (c *ClientSet) Close() error {
	return c.conn.Close()
}
