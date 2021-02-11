package server

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port           int64           `yaml:"port"`
	TLS            *TLS            `yaml:"tls"`
	Cors           *Cors           `yaml:"cors"`
	Logging        *Logging        `yaml:"logging"`
	Authorization  *Authorization  `yaml:"authorization"`
	Authentication *Authentication `yaml:"authentication"`
	Backend        *Backend        `yaml:"backend"`
}

type Cors struct {
	// Normalized list of plain allowed origins
	AllowedOrigins []string `yaml:"allowed_origins"`
	// Normalized list of allowed headers
	AllowedHeaders []string `yaml:"allowed_headers"`
	// Normalized list of allowed methods
	AllowedMethods []string `yaml:"allowed_methods"`
	// Normalized list of exposed headers
	ExposedHeaders []string `yaml:"exposed_headers"`
}

type TLS struct {
	Cert string `yaml:"cert_file"`
	Key  string `yaml:"key_file"`
}

type Backend struct {
	ChannelProvider *Provider `yaml:"channel_provider"`
	StorageProvider *Provider `yaml:"storage_provider"`
}

type Provider struct {
	Name   string            `yaml:"name"`
	Config map[string]string `yaml:"config"`
}

type Authentication struct {
	JwksURI string `yaml:"jwks_uri"`
}

type Authorization struct {
	RequestPolicy  string `yaml:"requests"`
	ResponsePolicy string `yaml:"responses"`
}

type Logging struct {
	Debug    bool `yaml:"debug"`
	Payloads bool `yaml:"payloads"`
}

func (c *Config) SetDefaults() {
	if c.Port == 0 {
		c.Port = 8080
	}
	if c.Logging == nil {
		c.Logging = &Logging{}
	}
	if c.Authentication == nil {
		c.Authentication = &Authentication{}
	}
	if c.Authorization == nil {
		c.Authorization = &Authorization{}
	}
	if c.Authorization.RequestPolicy == "" {
		// target = data.eventgate.requests.authz.allow
		c.Authorization.RequestPolicy = `
		package eventgate.authz

		default allow = false
`
	}
	if c.Authorization.ResponsePolicy == "" {
		// target = data.eventgate.responses.authz.allow
		c.Authorization.ResponsePolicy = `
		package eventgate.authz

		default allow = false
`
	}
	if c.Backend == nil {
		c.Backend = &Backend{}
	}
	if c.Backend.ChannelProvider == nil {
		c.Backend.ChannelProvider = &Provider{}
	}
	if c.Backend.StorageProvider == nil {
		c.Backend.StorageProvider = &Provider{}
	}
	if c.Backend.StorageProvider.Config == nil {
		c.Backend.StorageProvider.Config = map[string]string{}
	}
	if c.Backend.ChannelProvider.Config == nil {
		c.Backend.ChannelProvider.Config = map[string]string{}
	}
}

func ConfigFromFile(path string) (*Config, error) {
	bits, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	if err := yaml.UnmarshalStrict(bits, c); err != nil {
		return nil, err
	}
	c.SetDefaults()
	return c, nil
}
