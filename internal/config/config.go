package config

type Config struct {
	Port           int64   `yaml:"port"`
	Debug          bool    `yaml:"debug"`
	RequestPolicy  *Policy `yaml:"request_policy"`
	ResponsePolicy *Policy `yaml:"response_policy"`
	JwksURI        string  `yaml:"jwks_uri"`
	NatsURL        string  `yaml:"nats_url"`
}

type Policy struct {
	RegoPolicy string `yaml:"rego_policy"`
	RegoQuery  string `yaml:"rego_query"`
}

func (c *Config) SetDefaults() {
	if c.Port == 0 {
		c.Port = 8080
	}
	if c.RequestPolicy == nil {
		c.RequestPolicy = &Policy{}
	}
	if c.ResponsePolicy == nil {
		c.ResponsePolicy = &Policy{}
	}

	if c.RequestPolicy.RegoPolicy == "" {
		c.RequestPolicy.RegoPolicy = `
		package eventgate.authz

		default allow = true
`
	}
	if c.RequestPolicy.RegoQuery == "" {
		c.RequestPolicy.RegoQuery = "data.eventgate.authz.allow"
	}
	if c.ResponsePolicy.RegoPolicy == "" {
		c.ResponsePolicy.RegoPolicy = `
		package eventgate.authz

		default allow = true
`
	}
	if c.ResponsePolicy.RegoQuery == "" {
		c.ResponsePolicy.RegoQuery = "data.eventgate.authz.allow"
	}
	if c.NatsURL == "" {
		c.NatsURL = "0.0.0.0:4444"
	}
}
