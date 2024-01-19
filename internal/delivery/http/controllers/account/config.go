package account

import "crypto/rsa"

type Config struct {
	rsaPubKey *rsa.PublicKey
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) WithRSAPublicKey(key *rsa.PublicKey) *Config {
	c.rsaPubKey = key
	return c
}
