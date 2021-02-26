package stategate_client_go

import (
	"context"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/logger"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
)

// PeerClient is a stategate PeerService gRPC client
type PeerClient struct {
	client stategate.PeerServiceClient
	conn   *grpc.ClientConn
}

// NewPeerClient creates a new stategate PeerService client
func NewPeerClient(ctx context.Context, target string, opts ...Opt) (*PeerClient, error) {
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
	return &PeerClient{
		client: stategate.NewPeerServiceClient(conn),
		conn:   conn,
	}, nil
}

// Broadcast broadcasts a message to N subscribers(PeerService clients calling Stream)
func (p PeerClient) Broadcast(ctx context.Context, in *stategate.Message) error {
	_, err := p.client.Broadcast(ctx, in)
	return err
}

// Stream consumes/streams messages from message producers(PeerService clients calling broadcast)
func (p PeerClient) Stream(ctx context.Context, in *stategate.StreamMessageOpts, fn func(msg *stategate.PeerMessage) bool) error {
	stream, err := p.client.Stream(ctx, in)
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

// Close closes the gRPC client connection
func (c *PeerClient) Close() error {
	return c.conn.Close()
}
