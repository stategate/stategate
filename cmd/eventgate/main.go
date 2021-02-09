package main

import (
	"context"
	"crypto/tls"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/config"
	"github.com/autom8ter/eventgate/internal/gql"
	"github.com/autom8ter/eventgate/internal/helpers"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/providers"
	"github.com/autom8ter/eventgate/internal/storage"
	"github.com/autom8ter/machine"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/open-policy-agent/opa/rego"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	configPath string
)

func init() {
	pflag.CommandLine.StringVar(&configPath, "config", helpers.EnvOr("EVENTGATE_CONFIG", "config.yaml"), "path to config file (env: EVENTGATE_CONFIG)")
	pflag.Parse()
}

func main() {
	run(context.Background())
}

func run(ctx context.Context) {
	bits, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("failed to read config file: %s", err.Error())
		return
	}
	c := &config.Config{}
	if err := yaml.UnmarshalStrict(bits, c); err != nil {
		fmt.Printf("failed to read config file: %s", err.Error())
		return
	}
	c.SetDefaults()
	var lgger = logger.New(
		c.Logging.Debug,
		zap.String("channel_provider", c.Backend.ChannelProvider.Name),
		zap.String("storage_provider", c.Backend.StorageProvider.Name),
	)

	lgger.Debug("loaded config", zap.Any("config", c))
	var (
		m         = machine.New(ctx)
		interrupt = make(chan os.Signal, 1)
		apiLis    net.Listener
		tlsConfig *tls.Config
	)

	if c.TLS != nil && c.TLS.Key != "" && c.TLS.Cert != "" {
		cer, err := tls.LoadX509KeyPair(c.TLS.Cert, c.TLS.Key)
		if err != nil {
			lgger.Error("failed to load tls config", zap.Error(err))
			return
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
			lgger.Error("failed to create listener", zap.Error(err))
			return
		}
		apiLis, err = net.ListenTCP("tcp", addr)
		if err != nil {
			lgger.Error("failed to create api server listener", zap.Error(err))
			return
		}
	}
	defer apiLis.Close()
	apiMux := cmux.New(apiLis)
	apiMux.SetReadTimeout(1 * time.Second)
	grpcMatcher := apiMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	defer grpcMatcher.Close()
	m.Go(func(routine machine.Routine) {
		if err := apiMux.Serve(); err != nil && !strings.Contains(err.Error(), "closed network connection") {
			lgger.Error("listener mux error", zap.Error(err))
		}
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
		m.Go(func(routine machine.Routine) {
			addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", c.Port+1))
			if err != nil {
				lgger.Error("failed to create metrics listener", zap.Error(err))
				return
			}
			metricsLis, err := net.ListenTCP("tcp", addr)
			if err != nil {
				lgger.Error("failed to create metrics listener", zap.Error(err))
				return
			}
			defer metricsLis.Close()
			lgger.Info("starting metrics server", zap.String("address", metricsLis.Addr().String()))
			if err := metricServer.Serve(metricsLis); err != nil && err != http.ErrServerClosed {
				lgger.Error("metrics server failure", zap.Error(err))
			}
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
		lgger.Error(err.Error())
		return
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
			lgger.Error("failed to setup storage provider", zap.Error(err))
			return
		}
		defer strg.Close()
	}

	service, closer, err := providers.GetChannelProvider(providers.ChannelProvider(c.Backend.ChannelProvider.Name), strg, lgger, c.Backend.ChannelProvider.Config)
	if err != nil {
		lgger.Error("failed to setup channel provider", zap.Error(err))
		return
	}
	defer closer()
	gserver := grpc.NewServer(gopts...)
	eventgate.RegisterEventGateServiceServer(gserver, service)
	reflection.Register(gserver)
	grpc_prometheus.Register(gserver)

	m.Go(func(routine machine.Routine) {
		lgger.Info("starting grpc server",
			zap.String("address", grpcMatcher.Addr().String()),
		)
		if err := gserver.Serve(grpcMatcher); err != nil && err != http.ErrServerClosed {
			lgger.Error("grpc server failure", zap.Error(err))
		}
	})
	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("localhost:%v", c.Port),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		lgger.Error("failed to setup graphql server", zap.Error(err))
		return
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
			lgger.Error("failed to register REST endpoints", zap.Error(err))
			return
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

		m.Go(func(routine machine.Routine) {
			httpMatchermatcher := apiMux.Match(cmux.Any())
			defer httpMatchermatcher.Close()
			lgger.Info("starting http server", zap.String("address", httpMatchermatcher.Addr().String()))
			if err := httpServer.Serve(httpMatchermatcher); err != nil && err != http.ErrServerClosed {
				lgger.Error("http server failure", zap.Error(err))
			}
		})
	}

	select {
	case <-interrupt:
		m.Cancel()
		break
	case <-ctx.Done():
		m.Cancel()
		break
	}
	lgger.Debug("shutdown signal received")
	if httpServer != nil {
		go func() {
			shutdownctx, shutdowncancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer shutdowncancel()
			if err := httpServer.Shutdown(shutdownctx); err != nil {
				lgger.Error("http server shutdown failure", zap.Error(err))
			} else {
				lgger.Debug("shutdown http server")
			}
		}()
	}
	if metricServer != nil {
		go func() {
			shutdownctx, shutdowncancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer shutdowncancel()
			if err := metricServer.Shutdown(shutdownctx); err != nil {
				lgger.Error("metrics server shutdown failure", zap.Error(err))
			} else {
				lgger.Debug("shutdown metrics server")
			}

		}()
	}
	go func() {
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
	}()
	m.Wait()
	lgger.Debug("shutdown successful")
}
