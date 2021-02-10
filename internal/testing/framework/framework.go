package framework

import (
	"context"
	"fmt"
	eventgate_client_go "github.com/autom8ter/eventgate/eventgate-client-go"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/server"
	"github.com/ory/dockertest/v3"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
	"os"
	"testing"
	"time"
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
	Func func(t *testing.T, client *eventgate_client_go.Client)
}

type Provider struct {
	ctx    context.Context
	cancel func()
	config *server.Config
	group  *errgroup.Group
	lgger  *logger.Logger
	client *eventgate_client_go.Client
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
	f := &Provider{
		ctx:    ctx,
		cancel: cancel,
		config: config,
		group:  group,
		lgger:  logger.New(true),
		client: nil,
	}
	time.Sleep(1 * time.Second)
	f.group.Go(func() error {
		return server.ListenAndServe(f.ctx, f.lgger, f.config)
	})
	time.Sleep(1 * time.Second)
	client, err := eventgate_client_go.NewClient(
		ctx,
		fmt.Sprintf("localhost:%v", f.config.Port),
		eventgate_client_go.WithTokenSource(oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: jwt,
		})))
	if err != nil {
		t.Fatal(err.Error())
	}
	f.client = client
	return f
}

func (f *Provider) Teardown(t *testing.T) {
	f.cancel()
	if err := f.client.Close(); err != nil {
		t.Fatal(err)
	}
	if err := f.group.Wait(); err != nil {
		t.Fatal(err)
	}
}

func (f *Provider) runTestCase(t *testing.T, testCase *TestCase) {
	t.Run(testCase.Name, func(t *testing.T) {
		testCase.Func(t, f.client)
	})
}
