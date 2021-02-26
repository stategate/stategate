package stategate_client_go

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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
