package framework

import (
	"context"
	"fmt"
	"github.com/autom8ter/stategate/internal/logger"
	"github.com/autom8ter/stategate/internal/server"
	stategate_client_go "github.com/autom8ter/stategate/stategate-client-go"
	"github.com/ory/dockertest/v3"
	"github.com/spf13/cast"
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
	Func func(t *testing.T, eclient *stategate_client_go.EventClient, oclient *stategate_client_go.EntityClient)
}

type Provider struct {
	ctx     context.Context
	cancel  func()
	config  *server.Config
	group   *errgroup.Group
	lgger   *logger.Logger
	eclient *stategate_client_go.EventClient
	oclient *stategate_client_go.EntityClient
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
		zap.Any("channel_provider", cast.ToString(config.ChannelProvider["name"])),
		zap.Any("storage_provider", cast.ToString(config.StorageProvider["name"])),
	)
	f := &Provider{
		ctx:     ctx,
		cancel:  cancel,
		config:  config,
		group:   group,
		lgger:   lgger,
		eclient: nil,
		oclient: nil,
	}
	f.group.Go(func() error {
		return server.ListenAndServe(f.ctx, f.lgger, f.config)
	})
	eclient, err := stategate_client_go.NewEventClient(
		ctx,
		fmt.Sprintf("localhost:%v", f.config.Port),
		stategate_client_go.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: jwt,
		})))
	if err != nil {
		t.Fatal(err.Error())
	}
	oclient, err := stategate_client_go.NewEntityClient(
		ctx,
		fmt.Sprintf("localhost:%v", f.config.Port),
		stategate_client_go.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: jwt,
		})))
	if err != nil {
		t.Fatal(err.Error())
	}
	f.eclient = eclient
	f.oclient = oclient
	return f
}

func (f *Provider) Teardown(t *testing.T) {
	f.cancel()
	if err := f.eclient.Close(); err != nil {
		t.Fatal(err)
	}
	if err := f.oclient.Close(); err != nil {
		t.Fatal(err)
	}
	if err := f.group.Wait(); err != nil {
		t.Fatal(err)
	}
}

func (f *Provider) runTestCase(t *testing.T, testCase *TestCase) {
	t.Run(testCase.Name, func(t *testing.T) {
		testCase.Func(t, f.eclient, f.oclient)
	})
}
