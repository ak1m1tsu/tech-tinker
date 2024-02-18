package auth

import (
	"crypto/rsa"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	rsalib "github.com/insan1a/tech-tinker/internal/lib/rsa"
)

type config struct {
	HTTP struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"http"`
	RSA struct {
		PrivateKeyPath string          `yaml:"private_key"`
		PrivateKey     *rsa.PrivateKey `yaml:"-"`
	}
	Cache struct {
		TTL  time.Duration `yaml:"ttl"`
		Size int           `yaml:"size"`
	} `yaml:"cache"`
	DB struct {
		URL                string        `env:"DB_URL" env-required:"true"`
		PoolSize           int32         `yaml:"pool_size"`
		ConnectionAttempts uint          `yaml:"connection_attempts"`
		ConnectionTimeout  time.Duration `yaml:"connection_timeout"`
		RetryDelay         time.Duration `yaml:"retry_delay"`
	} `yaml:"db"`
}

func newConfig() (*config, error) {
	var (
		path = "/etc/auth-api/config.yaml"
		cfg  config
		err  error
	)

	if err = cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	cfg.RSA.PrivateKey, err = rsalib.PrivateKeyFromFile(cfg.RSA.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}