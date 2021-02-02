package config

import (
	"github.com/nats-io/stan.go"
	"helm.sh/helm/v3/pkg/repo"
)

type Config struct {
	Port        int64         `yaml:"port"`
	Debug       bool          `yaml:"debug"`
	RegoPolicy  string        `yaml:"rego_policy"`
	RegoQuery   string        `yaml:"rego_query"`
	JwksURI     string        `yaml:"jwks_uri"`
	Repos       []*repo.Entry `yaml:"repos"`
	NatsURL     string        `yaml:"nats_url"`
	NatsCluster string        `yaml:"nats_cluster"`
}

func (c *Config) SetDefaults() {
	if c.Port == 0 {
		c.Port = 8820
	}
	if c.RegoPolicy == "" {
		c.RegoPolicy = `
		package cloudEventsProxy.authz

		default allow = true
`
	}
	if c.RegoQuery == "" {
		c.RegoQuery = "data.cloudEventsProxy.authz.allow"
	}
	if c.NatsURL == "" {
		c.NatsURL = stan.DefaultNatsURL
	}
	if c.NatsCluster == "" {
		c.NatsCluster = "test-cluster"
	}
}
