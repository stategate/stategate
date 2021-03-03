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
	"io"
)

// EventClient is a stategate EventService gRPC client
type EventClient struct {
	client stategate.EventServiceClient
	conn   *grpc.ClientConn
}

// NewEventClient creates a new stategate EventService client
// EntityService serves API methods to clients that modify/query the current state of an entity
// An Entity is a single object with a type, domain, key, and k/v values
func NewEventClient(ctx context.Context, target string, opts ...Opt) (*EventClient, error) {
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
	return &EventClient{
		client: stategate.NewEventServiceClient(conn),
		conn:   conn,
	}, nil
}

// Stream creates an event stream/subscription to changes to entities  until fn returns false OR the context cancels.. Glob matching is supported.
func (c *EventClient) Stream(ctx context.Context, in *stategate.StreamEventOpts, fn func(even *stategate.Event) bool) error {
	stream, err := c.client.Stream(ctx, in)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			if !fn(msg) {
				return nil
			}
		}
	}
}

// Search queries historical events(changes to entities).
func (c *EventClient) Search(ctx context.Context, in *stategate.SearchEventOpts) (*stategate.Events, error) {
	return c.client.Search(ctx, in)
}

// Get gets an existing Event
func (c *EventClient) Get(ctx context.Context, in *stategate.EventRef) (*stategate.Event, error) {
	return c.client.Get(ctx, in)
}

// Close closes the gRPC client connection
func (c *EventClient) Close() error {
	return c.conn.Close()
}
