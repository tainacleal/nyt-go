package nyttop

import "time"

type Option func(*cfg)

type cfg struct {
	timeout time.Duration
}

func newConfig(opts ...Option) *cfg {
	cfg := defaultConfig()

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func defaultConfig() *cfg {
	return &cfg{
		timeout: 30 * time.Second,
	}
}

// WithTimeout sets the timeout for the HTTP Client. The default is 30 seconds.
func WithTimeout(t time.Duration) Option {
	return func(c *cfg) {
		c.timeout = t
	}
}
