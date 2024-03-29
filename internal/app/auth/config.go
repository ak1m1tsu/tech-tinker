package auth

import (
	"crypto/rsa"
	"time"

	rsalib "github.com/ak1m1tsu/tech-tinker/internal/lib/rsa"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Debug bool `yaml:"debug"`
	HTTP  struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"http"`
	JWT struct {
		PrivateKeyPath string          `yaml:"private_key"`
		PrivateKey     *rsa.PrivateKey `yaml:"-"`
		TTL            time.Duration   `yaml:"ttl"`
	} `yaml:"jwt"`
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

	cfg.JWT.PrivateKey, err = rsalib.PrivateKeyFromFile(cfg.JWT.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
