package config

import (
	"github.com/nats-io/stan.go"
)

type Config struct {
	Port       int64  `yaml:"port"`
	Debug      bool   `yaml:"debug"`
	RegoPolicy string `yaml:"rego_policy"`
	RegoQuery  string `yaml:"rego_query"`
	JwksURI    string `yaml:"jwks_uri"`
	NatsURL    string `yaml:"nats_url"`
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
}
