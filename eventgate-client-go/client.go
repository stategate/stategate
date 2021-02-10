package eventgate_client_go

import (
	"context"
	"fmt"
	"github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/logger"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
)

// Options holds configuration options
type Options struct {
	tokenSource oauth2.TokenSource
	metrics     bool
	logging     bool
	logPayload  bool
	creds       credentials.TransportCredentials
	idtoken     bool
}

// Opt is a single configuration option
type Opt func(o *Options)

// WithTransportCreds adds transport credentials to the client
func WithTransportCreds(creds credentials.TransportCredentials) Opt {
	return func(o *Options) {
		o.creds = creds
	}
}

// WithTokenSource uses oauth token add an authorization header to every outbound request
func WithTokenSource(tokenSource oauth2.TokenSource) Opt {
	return func(o *Options) {
		o.tokenSource = tokenSource
	}
}

// WithMetrics registers prometheus metrics
func WithMetrics(metrics bool) Opt {
	return func(o *Options) {
		o.metrics = metrics
	}
}

// WithLogging registers a logging middleware
func WithLogging(logging, logPayload bool) Opt {
	return func(o *Options) {
		o.logging = logging
		o.logPayload = logPayload
	}
}

// WithIDToken makes the client use the oauth id token(if it exists) instead of the oauth access token
func WithIDToken(idToken bool) Opt {
	return func(o *Options) {
		o.idtoken = idToken
	}
}

func unaryAuth(tokenSource oauth2.TokenSource, useIDToken bool) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, err := toContext(ctx, tokenSource, useIDToken)
		if err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func streamAuth(tokenSource oauth2.TokenSource, useIDToken bool) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx, err := toContext(ctx, tokenSource, useIDToken)
		if err != nil {
			return nil, err
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}

// NewClient creates a new gRPC eventgate client
func NewClient(ctx context.Context, target string, opts ...Opt) (*Client, error) {
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
		return nil, err
	}
	return &Client{
		client: eventgate.NewEventGateServiceClient(conn),
		conn:   conn,
	}, nil
}

// Client is a eventgate gRPC client
type Client struct {
	client eventgate.EventGateServiceClient
	conn   *grpc.ClientConn
}

func toContext(ctx context.Context, tokenSource oauth2.TokenSource, useIdToken bool) (context.Context, error) {
	token, err := tokenSource.Token()
	if err != nil {
		return ctx, errors.Wrap(err, "failed to get token")
	}

	if useIdToken {
		idToken := token.Extra("id_token")
		if idToken != nil {
			return metadata.AppendToOutgoingContext(
				ctx,
				"Authorization", fmt.Sprintf("Bearer %v", idToken.(string)),
			), nil
		}
	}
	return metadata.AppendToOutgoingContext(
		ctx,
		"Authorization", fmt.Sprintf("Bearer %v", token.AccessToken),
	), nil
}

// Send broadcasts an event to all consumers on a given channel
func (c *Client) Send(ctx context.Context, in *eventgate.Event) error {
	_, err := c.client.Send(ctx, in)
	return err
}

// Receive creates an event stream/subscription to a given channel until fn returns false OR the context cancels.
func (c *Client) Receive(ctx context.Context, in *eventgate.ReceiveOpts, fn func(even *eventgate.EventDetail) bool) error {
	if ctx.Err() != nil {
		return nil
	}
	stream, err := c.client.Receive(ctx, in)
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

// History returns an array of immutable historical events from a given channel.
func (c *Client) History(ctx context.Context, in *eventgate.HistoryOpts) (*eventgate.EventDetails, error) {
	return c.client.History(ctx, in)
}

// Close closes the gRPC client connection
func (c *Client) Close() error {
	return c.conn.Close()
}
