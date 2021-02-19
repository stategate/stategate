package stategate_client_go

import (
	"context"
	"fmt"
	"github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/logger"
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

// NewEventClient creates a new stategate client
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

// NewStateClient creates a new stategate state client
func NewStateClient(ctx context.Context, target string, opts ...Opt) (*StateClient, error) {
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
	return &StateClient{
		client: stategate.NewStateServiceClient(conn),
		conn:   conn,
	}, nil
}

// EventClient is a stategate EventService gRPC client
type EventClient struct {
	client stategate.EventServiceClient
	conn   *grpc.ClientConn
}

// StateClient is a stategate StateService gRPC client
type StateClient struct {
	client stategate.StateServiceClient
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

// Get gets an state's current state values
func (c *StateClient) Get(ctx context.Context, in *stategate.StateRef) (*stategate.State, error) {
	return c.client.Get(ctx, in)
}

// Search queries states of a specific type
func (c *StateClient) Search(ctx context.Context, in *stategate.SearchStateOpts) (*stategate.StateValues, error) {
	return c.client.Search(ctx, in)
}

// Set sets the current state value of an state, adds it to the event log, then broadcast the event to all interested consumers
func (c *StateClient) Set(ctx context.Context, in *stategate.State) error {
	_, err := c.client.Set(ctx, in)
	return err
}

// Del deletes an application state value(k/v pairs)
func (c *StateClient) Del(ctx context.Context, in *stategate.StateRef) error {
	_, err := c.client.Del(ctx, in)
	return err
}

// Close closes the gRPC client connection
func (c *StateClient) Close() error {
	return c.conn.Close()
}

// Stream creates an event stream/subscription to a given state type/domain until fn returns false OR the context cancels.
func (c *EventClient) Stream(ctx context.Context, in *stategate.StreamOpts, fn func(even *stategate.Event) bool) error {
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

// SearchEvents returns an array of immutable historical events for a given state.
func (c *EventClient) Search(ctx context.Context, in *stategate.SearchEventOpts) (*stategate.Events, error) {
	return c.client.Search(ctx, in)
}

// Close closes the gRPC client connection
func (c *EventClient) Close() error {
	return c.conn.Close()
}
