package framework

import (
	"context"
	"fmt"
	eventgate_client_go "github.com/autom8ter/eventgate/eventgate-client-go"
	"github.com/autom8ter/eventgate/internal/logger"
	"github.com/autom8ter/eventgate/internal/server"
	"golang.org/x/oauth2"
	"golang.org/x/sync/errgroup"
	"testing"
)

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
	ctx context.Context
	cancel func()
	config *server.Config
	group *errgroup.Group
	lgger *logger.Logger
	client *eventgate_client_go.Client
}

// creates an starts a new factory instance
func NewProvider(t *testing.T, ctx context.Context, config *server.Config, jwt string) *Provider {

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
	f.group.Go(func() error {
		return server.ListenAndServe(f.ctx, f.lgger, f.config)
	})
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

