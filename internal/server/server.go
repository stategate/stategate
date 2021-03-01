package server

import (
	"context"
	"crypto/tls"
	"fmt"
	stategate "github.com/autom8ter/stategate/gen/grpc/go"
	"github.com/autom8ter/stategate/internal/auth"
	"github.com/autom8ter/stategate/internal/cache"
	"github.com/autom8ter/stategate/internal/gql"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/autom8ter/stategate/internal/service"
	"github.com/autom8ter/stategate/internal/storage"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	if c.TLSKeyFile != "" && c.TLSCertFile != "" {
		cer, err := tls.LoadX509KeyPair(c.TLSCertFile, c.TLSKeyFile)
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

	metricsRouter := http.NewServeMux()
	metricsRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	metricsRouter.Handle("/metrics", promhttp.Handler())
	metricsRouter.HandleFunc("/debug/pprof/", pprof.Index)
	metricsRouter.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	metricsRouter.HandleFunc("/debug/pprof/profile", pprof.Profile)
	metricsRouter.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	metricsRouter.HandleFunc("/debug/pprof/trace", pprof.Trace)
	metricServer = &http.Server{Handler: metricsRouter}
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
	unary := []grpc.UnaryServerInterceptor{
		grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(lgger.Zap()),
	}
	stream := []grpc.StreamServerInterceptor{
		grpc_prometheus.StreamServerInterceptor,
		grpc_zap.StreamServerInterceptor(lgger.Zap()),
	}
	a, err := auth.NewAuth(c.AuthDisabled, c.RequestPolicy, c.ResponsePolicy, c.JWKSUri, lgger)
	if err != nil {
		return err
	}
	unary = append(unary, a.UnaryInterceptor())
	stream = append(stream, a.StreamInterceptor())
	unary = append(unary, grpc_validator.UnaryServerInterceptor(), grpc_recovery.UnaryServerInterceptor())
	stream = append(stream, grpc_validator.StreamServerInterceptor(), grpc_recovery.StreamServerInterceptor())

	gopts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(unary...),
		grpc.ChainStreamInterceptor(stream...),
	}
	if tlsConfig != nil {
		gopts = append(gopts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	var (
		strgProvider  storage.Provider
		cacheProvider cache.Provider
	)
	if c.StorageProvider != nil && len(c.StorageProvider) > 0 {
		strgProvider, err = storage.GetStorageProvider(lgger, c.StorageProvider)
		if err != nil {
			return errors.Wrap(err, "failed to setup storage provider")
		}
		defer strgProvider.Close()
	}
	cacheProvider, err = cache.GetCacheProvider(lgger, c.CacheProvider)
	if err != nil {
		return errors.Wrap(err, "failed to setup cache provider")
	}
	defer cacheProvider.Close()

	svc, err := service.NewService(ctx, strgProvider, cacheProvider, lgger)
	if err != nil {
		return errors.Wrap(err, "failed to setup service")
	}
	gserver := grpc.NewServer(gopts...)
	stategate.RegisterEventServiceServer(gserver, svc.EventServiceServer())
	stategate.RegisterEntityServiceServer(gserver, svc.EntityServiceServer())
	stategate.RegisterPeerServiceServer(gserver, svc.PeerServiceServer())
	stategate.RegisterCacheServiceServer(gserver, svc.CacheServiceServer())
	stategate.RegisterMutexServiceServer(gserver, svc.MutexServiceServer())
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

	restMux := runtime.NewServeMux()

	if err := stategate.RegisterEventServiceHandler(ctx, restMux, conn); err != nil {
		return errors.Wrap(err, "failed to register REST event endpoints")
	}
	if err := stategate.RegisterEntityServiceHandler(ctx, restMux, conn); err != nil {
		return errors.Wrap(err, "failed to register REST entity endpoints")
	}
	if err := stategate.RegisterPeerServiceHandler(ctx, restMux, conn); err != nil {
		return errors.Wrap(err, "failed to register REST peer endpoints")
	}
	if err := stategate.RegisterCacheServiceHandler(ctx, restMux, conn); err != nil {
		return errors.Wrap(err, "failed to register REST cache endpoints")
	}
	if err := stategate.RegisterMutexServiceHandler(ctx, restMux, conn); err != nil {
		return errors.Wrap(err, "failed to register REST mutex endpoints")
	}
	gresolver := gql.NewResolver(conn)
	defer gresolver.Close()
	if err := restMux.HandlePath(http.MethodPost, "/graphql", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		gresolver.QueryHandler()(w, r)
	}); err != nil {
		return errors.Wrap(err, "failed to register POST /graphql endpoints")
	}
	if err := restMux.HandlePath(http.MethodGet, "/graphql", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		gresolver.QueryHandler()(w, r)
	}); err != nil {
		return errors.Wrap(err, "failed to register POST /graphql endpoints")
	}
	mux.Handle("/", restMux)
	httpServer := &http.Server{
		Handler: mux,
	}
	if tlsConfig != nil {
		httpServer.TLSConfig = tlsConfig
		httpServer.TLSNextProto = map[string]func(*http.Server, *tls.Conn, http.Handler){}
	}
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
	group.Go(func() error {
		httpMatchermatcher := apiMux.Match(cmux.Any())
		defer httpMatchermatcher.Close()
		lgger.Debug("starting http server", zap.String("address", httpMatchermatcher.Addr().String()))
		if err := httpServer.Serve(httpMatchermatcher); err != nil && err != http.ErrServerClosed && !strings.Contains(err.Error(), "use of closed network connection") {
			return errors.Wrap(err, "http server failure")
		}
		return nil
	})
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

		shutdownGroup.Go(func() error {
			httpServer.Shutdown(shutdownctx)
			lgger.Debug("shutdown http server")
			return nil
		})
		shutdownGroup.Go(func() error {
			metricServer.Shutdown(shutdownctx)
			lgger.Debug("shutdown metrics server")
			return nil
		})
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
