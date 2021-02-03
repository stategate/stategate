package config

type Config struct {
	Port           int64           `yaml:"port"`
	Logging        *Logging        `yaml:"logging"`
	Authorization  *Authorization  `yaml:"authorization"`
	Authentication *Authentication `yaml:"authentication"`
	Backend        *Backend        `yaml:"backend"`
}

type Backend struct {
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
		package eventgate.authz.requests

		default allow = false
`
	}
	if c.Authorization.ResponsePolicy == "" {
		// target = data.eventgate.responses.authz.allow
		c.Authorization.ResponsePolicy = `
		package eventgate.authz.responses

		default allow = false
`
	}
	if c.Backend == nil {
		c.Backend = &Backend{}
	}
	if c.Backend.Config == nil {
		c.Backend.Config = map[string]string{}
	}
}
