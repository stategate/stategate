package gql

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/autom8ter/eventgate/gen/gql/go/generated"
	"github.com/autom8ter/eventgate/gen/gql/go/model"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger *logger.Logger
	client eventgate.EventGateServiceClient
}

func NewResolver(client eventgate.EventGateServiceClient, logger *logger.Logger) *Resolver {
	return &Resolver{logger: logger, client: client}
}

func (r *Resolver) QueryHandler() http.Handler {
	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers:  r,
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			auth := initPayload.Authorization()
			ctx = metadata.AppendToOutgoingContext(ctx, "Authorization", auth)
			return ctx, nil
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(&apollotracing.Tracer{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	return r.authMiddleware(srv)
}

func (r *Resolver) authMiddleware(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		for k, arr := range req.Header {
			if len(arr) > 0 {
				ctx = metadata.AppendToOutgoingContext(ctx, k, arr[0])
			}
		}
		handler.ServeHTTP(w, req.WithContext(ctx))
	}
}

func (r *Resolver) toEvent(msg *eventgate.EventDetail) *model.EventDetail {
	return &model.EventDetail{
		ID:       msg.GetId(),
		Channel:  msg.GetChannel(),
		Data:     msg.GetData().AsMap(),
		Metadata: msg.GetMetadata().AsMap(),
		Claims:   msg.GetClaims().AsMap(),
		Time:     msg.Time.AsTime(),
	}
}
