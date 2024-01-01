package postgres

import "time"

const (
	defaultAttempts = 5
	defaultPoolSize = 10
	defaultTimeout  = 5 * time.Second
)

type Option func(*Conn)

func WithConnectionAttempts(maxAttempts int) Option {
	return func(c *Conn) {
		c.connAttempts = maxAttempts
	}
}

func WithConnectionTimeout(timeout time.Duration) Option {
	return func(c *Conn) {
		c.connTimeout = timeout
	}
}

func WithMaxPoolSize(size int) Option {
	return func(c *Conn) {
		c.maxPoolSize = size
	}
}
