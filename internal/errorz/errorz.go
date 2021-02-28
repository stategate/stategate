package errorz

import (
	"context"
	"encoding/json"
	"github.com/autom8ter/stategate/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Type int

const (
	ErrUnknown         Type = 0
	ErrUnauthenticated Type = 1
	ErrUnauthorized    Type = 2
	ErrNotFound        Type = 5
	ErrTimeout         Type = 6
	ErrLocked          Type = 7
)

type Error struct {
	Type     Type              `json:"type"`
	Info     string            `json:"info"`
	Err      error             `json:"err"`
	Metadata map[string]string `json:"metadata"`
}

func (e Error) Public() error {
	data := map[string]interface{}{
		"type":     e.Type,
		"info":     e.Info,
		"metadata": e.Metadata,
	}
	bits, _ := json.Marshal(&data)
	switch e.Type {
	case ErrNotFound:
		return status.Error(codes.NotFound, string(bits))
	case ErrUnauthenticated:
		return status.Error(codes.Unauthenticated, string(bits))
	case ErrUnauthorized:
		return status.Error(codes.PermissionDenied, string(bits))
	case ErrUnknown:
		return status.Error(codes.Internal, string(bits))
	case ErrTimeout:
		return status.Error(codes.DeadlineExceeded, string(bits))
	case ErrLocked:
		return status.Error(codes.AlreadyExists, string(bits))
	default:
		if e.Err == context.Canceled {
			return status.Error(codes.DeadlineExceeded, string(bits))
		}
		return status.Error(codes.Internal, string(bits))
	}
}

func (e Error) Log(logger *logger.Logger) {
	logger.Error(e.Info,
		zap.Int("type", int(e.Type)),
		zap.Error(e.Err),
		zap.Any("metadata", e.Metadata))
}
