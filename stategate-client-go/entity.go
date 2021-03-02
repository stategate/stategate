package stategate_client_go

import (
	"context"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	stategate "github.com/stategate/stategate/gen/grpc/go"
	"github.com/stategate/stategate/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// EntityClient is a stategate EntityService gRPC client
type EntityClient struct {
	client stategate.EntityServiceClient
	conn   *grpc.ClientConn
}

// NewEntityClient creates a new stategate EntityService client
func NewEntityClient(ctx context.Context, target string, opts ...Opt) (*EntityClient, error) {
	if target == "" {
		return nil, errors.New("empty target")
	}
	dialopts := []grpc.DialOption{}
	var uinterceptors []grpc.UnaryClientInterceptor
	var sinterceptors []grpc.StreamClientInterceptor
	options := &Options{}
	for _, o := range opts {
		o(options)
	}
	if options.creds == nil {
		dialopts = append(dialopts, grpc.WithInsecure())
	} else {
		dialopts = append(dialopts, grpc.WithTransportCredentials(options.creds))
	}
	uinterceptors = append(uinterceptors, grpc_validator.UnaryClientInterceptor())

	if options.metrics {
		uinterceptors = append(uinterceptors, grpc_prometheus.UnaryClientInterceptor)
		sinterceptors = append(sinterceptors, grpc_prometheus.StreamClientInterceptor)
	}

	if options.tokenSource != nil {
		uinterceptors = append(uinterceptors, unaryAuth(options.tokenSource, options.idtoken))
		sinterceptors = append(sinterceptors, streamAuth(options.tokenSource, options.idtoken))
	}
	if options.logging {
		lgger := logger.New(true, zap.Bool("client", true))
		uinterceptors = append(uinterceptors, grpc_zap.UnaryClientInterceptor(lgger.Zap()))
		sinterceptors = append(sinterceptors, grpc_zap.StreamClientInterceptor(lgger.Zap()))

		if options.logPayload {
			uinterceptors = append(uinterceptors, grpc_zap.PayloadUnaryClientInterceptor(lgger.Zap(), func(ctx context.Context, fullMethodName string) bool {
				return true
			}))
			sinterceptors = append(sinterceptors, grpc_zap.PayloadStreamClientInterceptor(lgger.Zap(), func(ctx context.Context, fullMethodName string) bool {
				return true
			}))
		}
	}
	dialopts = append(dialopts,
		grpc.WithChainUnaryInterceptor(uinterceptors...),
		grpc.WithChainStreamInterceptor(sinterceptors...),
		grpc.WithBlock(),
	)
	conn, err := grpc.DialContext(ctx, target, dialopts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create stategate client")
	}
	return &EntityClient{
		client: stategate.NewEntityServiceClient(conn),
		conn:   conn,
	}, nil
}

// Get gets an entity's current state
func (c *EntityClient) Get(ctx context.Context, in *stategate.EntityRef) (*stategate.Entity, error) {
	return c.client.Get(ctx, in)
}

// Edit overwrites the k/v pairs present in the entity request without replacing the entire entity.
// It then adds the state change to the event log, then broadcast the event to all interested consumers(EventService.Stream)
func (c *EntityClient) Edit(ctx context.Context, in *stategate.Entity) (*stategate.Entity, error) {
	return c.client.Edit(ctx, in)
}

// Search queries the current state of entities
func (c *EntityClient) Search(ctx context.Context, in *stategate.SearchEntityOpts) (*stategate.Entities, error) {
	return c.client.Search(ctx, in)
}

// Set sets the current state value of an entity, adds it to the event log, then broadcast the event to all interested consumers
func (c *EntityClient) Set(ctx context.Context, in *stategate.Entity) error {
	_, err := c.client.Set(ctx, in)
	return err
}

// Revert reverts an Entity to a previous version of itself
func (c *EntityClient) Revert(ctx context.Context, in *stategate.EventRef) (*stategate.Entity, error) {
	return c.client.Revert(ctx, in)
}

// Del deletes an application state value(k/v pairs)
func (c *EntityClient) Del(ctx context.Context, in *stategate.EntityRef) error {
	_, err := c.client.Del(ctx, in)
	return err
}

// Close closes the gRPC client connection
func (c *EntityClient) Close() error {
	return c.conn.Close()
}
