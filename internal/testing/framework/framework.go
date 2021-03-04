package framework

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/spf13/cast"
	"github.com/stategate/stategate/internal/logger"
	"github.com/stategate/stategate/internal/server"
	stategate_client_go "github.com/stategate/stategate/stategate-client-go"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
	"os"
	"testing"
)

func init() {
	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

var pool *dockertest.Pool

func Run(t *testing.T, providers []*Provider, testCases []*TestCase) {
	for _, p := range providers {
		defer p.Teardown(t)
		for _, testCase := range testCases {
			p.runTestCase(t, testCase)
		}
	}
}

type TestCase struct {
	Name string
	Func func(t *testing.T, clientset *stategate_client_go.ClientSet)
}

type Provider struct {
	ctx       context.Context
	cancel    func()
	config    *server.Config
	group     *errgroup.Group
	lgger     *logger.Logger
	clientset *stategate_client_go.ClientSet
}

type Container struct {
	resource *dockertest.Resource
}

func NewContainer(t *testing.T, repo, tag string, env []string) *Container {
	resource, err := pool.Run(repo, tag, env)
	if err != nil {
		t.Fatal(err.Error())
	}
	return &Container{resource: resource}
}

func (c *Container) GetPort(id string) string {
	return c.resource.GetPort(id)
}

func (c *Container) Close(t *testing.T) {
	if err := pool.Purge(c.resource); err != nil {
		t.Fatal(err)
	}
}

// creates an starts a new factory instance
func NewProvider(t *testing.T, ctx context.Context, jwt string, config *server.Config) *Provider {
	config.SetDefaults()
	ctx, cancel := context.WithCancel(ctx)
	group, ctx := errgroup.WithContext(ctx)
	var lgger = logger.New(
		config.Debug,
		zap.Any("storage_provider", cast.ToString(config.StorageProvider["name"])),
		zap.Any("cache_provider", cast.ToString(config.CacheProvider["name"])),
		zap.Any("channel_provider", cast.ToString(config.ChannelProvider["name"])),
	)
	f := &Provider{
		ctx:       ctx,
		cancel:    cancel,
		config:    config,
		group:     group,
		lgger:     lgger,
		clientset: nil,
	}
	f.group.Go(func() error {
		return server.ListenAndServe(f.ctx, f.lgger, f.config)
	})
	clientset, err := stategate_client_go.NewClientSet(
		ctx,
		fmt.Sprintf("localhost:%v", f.config.Port),
		stategate_client_go.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: jwt,
		})))
	if err != nil {
		t.Fatal(err)
	}
	f.clientset = clientset
	return f
}

func (f *Provider) Teardown(t *testing.T) {
	f.cancel()
	if err := f.clientset.Close(); err != nil {
		t.Fatal(err)
	}
	if err := f.group.Wait(); err != nil {
		t.Fatal(err)
	}
}

func (f *Provider) runTestCase(t *testing.T, testCase *TestCase) {
	t.Run(testCase.Name, func(t *testing.T) {
		testCase.Func(t, f.clientset)
	})
}
