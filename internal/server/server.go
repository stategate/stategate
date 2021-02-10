package server

import (
	"context"
	"crypto/tls"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/gql"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/providers"
	"github.com/autom8ter/eventgate/internal/storage"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/open-policy-agent/opa/rego"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func ListenAndServe(ctx context.Context, lgger *logger.Logger, c *Config) error {
	group, ctx := errgroup.WithContext(ctx)

	var (
		interrupt = make(chan os.Signal, 1)
		apiLis    net.Listener
		tlsConfig *tls.Config
	)

	if c.TLS != nil && c.TLS.Key != "" && c.TLS.Cert != "" {
		cer, err := tls.LoadX509KeyPair(c.TLS.Cert, c.TLS.Key)
		if err != nil {
			lgger.Error("failed to load tls config", zap.Error(err))
			return err
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cer},
		}
	}

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)
	{
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", c.Port))
		if err != nil {
			return errors.Wrap(err, "failed to create listener")
		}
		apiLis, err = net.ListenTCP("tcp", addr)
		if err != nil {
			return errors.Wrap(err, "failed to create api server listener")
		}
	}
	defer apiLis.Close()
	apiMux := cmux.New(apiLis)
	apiMux.SetReadTimeout(1 * time.Second)
	grpcMatcher := apiMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	defer grpcMatcher.Close()
	group.Go(func() error {
		if err := apiMux.Serve(); err != nil && !strings.Contains(err.Error(), "closed network connection") {
			return errors.Wrap(err, "listener mux error")
		}
		return nil
	})

	var metricServer *http.Server

	if c.Metrics {
		router := http.NewServeMux()
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		router.Handle("/metrics", promhttp.Handler())
		router.HandleFunc("/debug/pprof/", pprof.Index)
		router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		router.HandleFunc("/debug/pprof/profile", pprof.Profile)
		router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		router.HandleFunc("/debug/pprof/trace", pprof.Trace)
		metricServer = &http.Server{Handler: router}
		if tlsConfig != nil {
			metricServer.TLSConfig = tlsConfig
			metricServer.TLSNextProto = map[string]func(*http.Server, *tls.Conn, http.Handler){}
		}
		group.Go(func() error {
			addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", c.Port+1))
			if err != nil {
				return errors.Wrap(err, "failed to create metrics listener")
			}
			metricsLis, err := net.ListenTCP("tcp", addr)
			if err != nil {
				return errors.Wrap(err, "failed to create metrics listener")
			}
			defer metricsLis.Close()
			lgger.Debug("starting metrics server", zap.String("address", metricsLis.Addr().String()))
			if err := metricServer.Serve(metricsLis); err != nil && err != http.ErrServerClosed {
				return errors.Wrap(err, "metrics server failure")
			}
			return nil
		})
	}

	const (
		regoQuery = "data.eventgate.authz.allow"
	)

	requestPolicy := rego.New(
		rego.Query(regoQuery),
		rego.Module("requests.rego", c.Authorization.RequestPolicy),
	)
	respPolicy := rego.New(
		rego.Query(regoQuery),
		rego.Module("responses.rego", c.Authorization.ResponsePolicy),
	)
	a, err := auth.NewAuth(c.Authentication.JwksURI, lgger, requestPolicy, respPolicy)
	if err != nil {
		return err
	}
	unary := []grpc.UnaryServerInterceptor{
		grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(lgger.Zap()),
		a.UnaryInterceptor(),
		grpc_validator.UnaryServerInterceptor(),
		grpc_recovery.UnaryServerInterceptor(),
	}
	stream := []grpc.StreamServerInterceptor{
		grpc_prometheus.StreamServerInterceptor,
		grpc_zap.StreamServerInterceptor(lgger.Zap()),
		a.StreamInterceptor(),
		grpc_validator.StreamServerInterceptor(),
		grpc_recovery.StreamServerInterceptor(),
	}
	if c.Logging.Payloads {
		unary = append(unary, grpc_zap.PayloadUnaryServerInterceptor(lgger.Zap(), func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
			return true
		}))
		stream = append(stream, grpc_zap.PayloadStreamServerInterceptor(lgger.Zap(), func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
			return true
		}))
	}
	gopts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(unary...),
		grpc.ChainStreamInterceptor(stream...),
	}
	if tlsConfig != nil {
		gopts = append(gopts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}
	var strg storage.Provider
	if c.Backend.StorageProvider != nil && c.Backend.StorageProvider.Name != "" {
		strg, err = providers.GetStorageProvider(providers.StorageProvider(c.Backend.StorageProvider.Name), lgger, c.Backend.StorageProvider.Config)
		if err != nil {
			return errors.Wrap(err, "failed to setup storage provider")
		}
		defer strg.Close()
	}

	service, closer, err := providers.GetChannelProvider(providers.ChannelProvider(c.Backend.ChannelProvider.Name), strg, lgger, c.Backend.ChannelProvider.Config)
	if err != nil {
		return errors.Wrap(err, "failed to setup channel provider")
	}
	defer closer()
	gserver := grpc.NewServer(gopts...)
	eventgate.RegisterEventGateServiceServer(gserver, service)
	reflection.Register(gserver)
	grpc_prometheus.Register(gserver)

	group.Go(func() error {
		lgger.Debug("starting grpc server",
			zap.String("address", grpcMatcher.Addr().String()),
		)
		if err := gserver.Serve(grpcMatcher); err != nil {
			return errors.Wrap(err, "grpc server failure")
		}
		return nil
	})

	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("localhost:%v", c.Port),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return errors.Wrap(err, "failed to setup graphql server")
	}
	defer conn.Close()
	mux := http.NewServeMux()
	if c.GraphQL {
		resolver := gql.NewResolver(eventgate.NewEventGateServiceClient(conn), lgger)
		mux.Handle("/graphql", resolver.QueryHandler())
	}

	if c.Rest {
		restMux := runtime.NewServeMux()
		if err := eventgate.RegisterEventGateServiceHandler(ctx, restMux, conn); err != nil {
			return errors.Wrap(err, "failed to register REST endpoints")
		}
		mux.Handle("/", restMux)
	}

	var httpServer *http.Server
	if c.GraphQL || c.Rest {
		httpServer = &http.Server{
			Handler: mux,
		}
		if tlsConfig != nil {
			httpServer.TLSConfig = tlsConfig
			httpServer.TLSNextProto = map[string]func(*http.Server, *tls.Conn, http.Handler){}
		}
		if c.GrpcWeb {
			wrappedGrpc := grpcweb.WrapServer(
				gserver,
				grpcweb.WithWebsockets(true),
				grpcweb.WithWebsocketPingInterval(15*time.Second),
			)
			httpServer.Handler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				if wrappedGrpc.IsGrpcWebRequest(req) {
					wrappedGrpc.ServeHTTP(resp, req)
				} else {
					mux.ServeHTTP(resp, req)
				}
			})
			if c.Cors != nil {
				c := cors.New(cors.Options{
					AllowedOrigins: c.Cors.AllowedOrigins,
					AllowedMethods: c.Cors.AllowedMethods,
					AllowedHeaders: c.Cors.AllowedHeaders,
					ExposedHeaders: c.Cors.ExposedHeaders,
				})
				httpServer.Handler = c.Handler(httpServer.Handler)
			}
		}

		group.Go(func() error {
			httpMatchermatcher := apiMux.Match(cmux.Any())
			defer httpMatchermatcher.Close()
			lgger.Debug("starting http server", zap.String("address", httpMatchermatcher.Addr().String()))
			if err := httpServer.Serve(httpMatchermatcher); err != nil && err != http.ErrServerClosed {
				return errors.Wrap(err, "http server failure")
			}
			return nil
		})
	}
	group.Go(func() error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		select {
		case <-interrupt:
			cancel()
			break
		case <-ctx.Done():
			break
		}

		lgger.Debug("shutdown signal received")
		shutdownctx, shutdowncancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer shutdowncancel()
		shutdownGroup, shutdownctx := errgroup.WithContext(shutdownctx)

		if httpServer != nil {
			shutdownGroup.Go(func() error {
				if err := httpServer.Shutdown(shutdownctx); err != nil {
					return errors.Wrap(err, "http server shutdown failure")
				}
				return nil
			})
		}
		if metricServer != nil {
			shutdownGroup.Go(func() error {
				if err := metricServer.Shutdown(shutdownctx); err != nil {
					return errors.Wrap(err, "metric server shutdown failure")
				}
				return nil
			})
		}
		shutdownGroup.Go(func() error {
			stopped := make(chan struct{})
			go func() {
				gserver.GracefulStop()
				stopped <- struct{}{}
			}()
			select {
			case <-time.After(15 * time.Second):
				gserver.Stop()
			case <-stopped:
				close(stopped)
				break
			}
			lgger.Debug("shutdown gRPC server")
			return nil
		})
		return shutdownGroup.Wait()
	})

	return group.Wait()
}
