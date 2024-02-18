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
	JWT struct {
		PrivateKey *rsa.PrivateKey
		TTL        time.Duration
	}
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithJWTPublicKey(key *rsa.PrivateKey) *Config {
	c.JWT.PrivateKey = key
	return c
}

func (c *Config) WithJWTTTL(ttl time.Duration) *Config {
	c.JWT.TTL = ttl
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
