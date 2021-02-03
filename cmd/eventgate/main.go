package main

import (
	"context"
	"fmt"
	eventgate "github.com/autom8ter/eventgate/gen/grpc/go"
	"github.com/autom8ter/eventgate/internal/auth"
	"github.com/autom8ter/eventgate/internal/backend"
	"github.com/autom8ter/eventgate/internal/config"
	"github.com/autom8ter/eventgate/internal/gql"
	"github.com/autom8ter/eventgate/internal/helpers"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/machine"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/open-policy-agent/opa/rego"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
	var lgger = logger.New(c.Logging.Debug)

	lgger.Debug("loaded config", zap.Any("config", c))
	var (
		metricsLis net.Listener
		m          = machine.New(ctx)
		interrupt  = make(chan os.Signal, 1)
		apiLis     net.Listener
	)
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
	httpMatchermatcher := apiMux.Match(cmux.Any())
	defer httpMatchermatcher.Close()
	m.Go(func(routine machine.Routine) {
		if err := apiMux.Serve(); err != nil && !strings.Contains(err.Error(), "closed network connection") {
			lgger.Error("listener mux error", zap.Error(err))
		}
	})
	{
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", c.Port+1))
		if err != nil {
			lgger.Error("failed to create listener", zap.Error(err))
			return
		}
		metricsLis, err = net.ListenTCP("tcp", addr)
		if err != nil {
			lgger.Error("failed to create listener", zap.Error(err))
			return
		}
	}
	defer metricsLis.Close()
	var metricServer *http.Server
	{
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
	}
	lgger.Info("starting metrics server", zap.String("address", metricsLis.Addr().String()))
	m.Go(func(routine machine.Routine) {
		if err := metricServer.Serve(metricsLis); err != nil && err != http.ErrServerClosed {
			lgger.Error("metrics server failure", zap.Error(err))
		}
	})

	const (
		requestQuery  = "data.eventgate.authz.requests.allow"
		responseQuery = "data.eventgate.authz.responses.allow"
	)

	requestPolicy := rego.New(
		rego.Query(requestQuery),
		rego.Module("requests.rego", c.Authorization.RequestPolicy),
	)
	respPolicy := rego.New(
		rego.Query(responseQuery),
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
	service, closer, err := backend.GetProvider(backend.Provider(c.Backend.Name), c.Backend.Config)
	if err != nil {
		lgger.Error(err.Error())
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
	resolver := gql.NewResolver(eventgate.NewEventGateServiceClient(conn), lgger)
	mux := http.NewServeMux()

	mux.Handle("/graphql", resolver.QueryHandler())

	{
		restMux := runtime.NewServeMux()
		if err := eventgate.RegisterEventGateServiceHandler(ctx, restMux, conn); err != nil {
			lgger.Error("failed to register REST endpoints", zap.Error(err))
			return
		}
		mux.Handle("/", restMux)
	}
	httpServer := &http.Server{
		Handler: mux,
	}
	m.Go(func(routine machine.Routine) {
		lgger.Info("starting http server", zap.String("address", httpMatchermatcher.Addr().String()))
		if err := httpServer.Serve(httpMatchermatcher); err != nil && err != http.ErrServerClosed {
			lgger.Error("http server failure", zap.Error(err))
		}
	})
	select {
	case <-interrupt:
		m.Cancel()
		break
	case <-ctx.Done():
		m.Cancel()
		break
	}
	lgger.Debug("shutdown signal received")
	go func() {
		shutdownctx, shutdowncancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer shutdowncancel()
		if err := httpServer.Shutdown(shutdownctx); err != nil {
			lgger.Error("http server shutdown failure", zap.Error(err))
		} else {
			lgger.Debug("shutdown http server")
		}
	}()
	go func() {
		shutdownctx, shutdowncancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer shutdowncancel()
		if err := metricServer.Shutdown(shutdownctx); err != nil {
			lgger.Error("metrics server shutdown failure", zap.Error(err))
		} else {
			lgger.Debug("shutdown metrics server")
		}

	}()
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
