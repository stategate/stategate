package auth

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type stream struct {
	ss   grpc.ServerStream
	a    *Auth
	ctx context.Context
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
	md := metautils.ExtractOutgoing(s.ctx)
	respMeta := make(map[string]string)
	for k, arr := range md {
		if len(arr) > 0 {
			respMeta[k] = arr[0]
		}
	}
	c, ok := GetContext(s.Context())
	if !ok {
		return status.Error(codes.PermissionDenied, "failed to get context")
	}
	c.Metadata = respMeta
	c.Body = toMap(m)
	allowed, err := s.a.evaluateResponse(s.ctx, c)
	if err != nil {
		s.a.logger.Error(err.Error())
		return status.Error(codes.Internal, "failed to evaluate authz policy during server stream")
	}
	if !allowed {
		return status.Error(codes.PermissionDenied, respDenied)
	}
	return s.ss.SendMsg(m)
}

func (s *stream) RecvMsg(m interface{}) error {
	md := metautils.ExtractIncoming(s.ctx)
	reqMeta := make(map[string]string)
	for k, arr := range md {
		if len(arr) > 0 {
			reqMeta[k] = arr[0]
		}

	}
	c, ok := GetContext(s.Context())
	if !ok {
		return status.Error(codes.PermissionDenied, "failed to get context")
	}
	c.Metadata = reqMeta
	c.Body = toMap(m)
	allowed, err := s.a.evaluateRequest(s.ctx, c)
	if err != nil {
		s.a.logger.Error(err.Error())
		return status.Error(codes.Internal, "failed to evaluate authz policy during client stream")
	}
	if !allowed {
		return status.Error(codes.PermissionDenied, reqDenied)
	}
	return s.ss.RecvMsg(m)
}
