package auth

import (
	"crypto/rsa"
	"github.com/google/uuid"
	"time"
)

type Config struct {
	key      *rsa.PrivateKey
	tokenTTL time.Duration
	secret   string
}

func NewConfig() *Config {
	return &Config{
		tokenTTL: time.Hour,
		secret:   uuid.NewString(),
	}
}

func (c *Config) WithRSAPrivateKey(key *rsa.PrivateKey) *Config {
	c.key = key
	return c
}

func (c *Config) WithTokenTTL(ttl time.Duration) *Config {
	c.tokenTTL = ttl
	return c
}

func (c *Config) WithSecret(secret string) *Config {
	c.secret = secret
	return c
}
