package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type stream struct {
	ctx  context.Context
	ss   grpc.ServerStream
	a    *Auth
	info *grpc.StreamServerInfo
}

func (s *stream) SetHeader(md metadata.MD) error {
	return s.ss.SetHeader(md)
}

func (s *stream) SendHeader(md metadata.MD) error {
	return s.ss.SendHeader(md)
}

func (s *stream) SetTrailer(md metadata.MD) {
	s.ss.SetTrailer(md)
}

func (s *stream) Context() context.Context {
	return s.ctx
}

func (s *stream) SendMsg(m interface{}) error {
	ctx := s.ctx
	c, ok := GetContext(ctx)
	if !ok {
		return status.Error(codes.PermissionDenied, "permission denied")
	}
	c.Response = toMap(m)
	allowed, err := s.a.booleanExpression(ctx, c)
	if err != nil {
		s.a.logger.Error(err.Error())
		return status.Error(codes.Internal, "failed to evaluate authz policy")
	}
	if !allowed {
		return status.Error(codes.PermissionDenied, "permission denied")
	}
	return s.ss.SendMsg(m)
}

func (s *stream) RecvMsg(m interface{}) error {
	ctx := s.ctx
	c, ok := GetContext(ctx)
	if !ok {
		return status.Error(codes.PermissionDenied, "permission denied")
	}
	c.Request = toMap(m)
	allowed, err := s.a.booleanExpression(ctx, c)
	if err != nil {
		s.a.logger.Error(err.Error())
		return status.Error(codes.Internal, "failed to evaluate authz policy")
	}
	if !allowed {
		return status.Error(codes.PermissionDenied, "permission denied")
	}
	return s.ss.RecvMsg(m)
}
