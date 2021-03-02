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

// CacheClient is a stategate CacheService gRPC client
type CacheClient struct {
	client stategate.CacheServiceClient
	conn   *grpc.ClientConn
}

// NewCacheClient creates a new stategate CacheService client
func NewCacheClient(ctx context.Context, target string, opts ...Opt) (*CacheClient, error) {
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
	return &CacheClient{
		client: stategate.NewCacheServiceClient(conn),
		conn:   conn,
	}, nil
}

func (m *CacheClient) Set(ctx context.Context, value *stategate.Cache) error {
	_, err := m.client.Set(ctx, value)
	return err
}

func (m *CacheClient) Del(ctx context.Context, ref *stategate.CacheRef) error {
	_, err := m.client.Del(ctx, ref)
	return err
}

func (m *CacheClient) Get(ctx context.Context, ref *stategate.CacheRef) (*stategate.Cache, error) {
	return m.client.Get(ctx, ref)
}

// Close closes the gRPC client connection
func (c *CacheClient) Close() error {
	return c.conn.Close()
}
