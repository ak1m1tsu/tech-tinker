package publicapi

import (
	"crypto/rsa"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	rsalib "github.com/insan1a/tech-tinker/internal/lib/rsa"
)

type config struct {
	Debug bool `yaml:"debug"`
	HTTP  struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"http"`
	JWT struct {
		PublicKeyPath string         `yaml:"public_key"`
		PublicKey     *rsa.PublicKey `yaml:"-"`
	} `yaml:"jwt"`
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
		path = "/etc/tech-tinker/config.yaml"
		cfg  config
		err  error
	)

	if err = cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	cfg.JWT.PublicKey, err = rsalib.PublicKeyFromFile(cfg.JWT.PublicKeyPath)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
