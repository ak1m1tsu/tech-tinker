package auth

import (
	"crypto/rsa"
	"time"
)

type Config struct {
	Repo struct {
	}
	Cache struct {
		TTL  time.Duration
		Size int
	}
	PrivateKey *rsa.PrivateKey
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithPublicKey(key *rsa.PrivateKey) *Config {
	c.PrivateKey = key
	return c
}

func (c *Config) WithCacheTTL(ttl time.Duration) *Config {
	c.Cache.TTL = ttl
	return c
}

func (c *Config) WithCacheSize(size int) *Config {
	c.Cache.Size = size
	return c
}
