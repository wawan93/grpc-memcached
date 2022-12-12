package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPC      GRPC      `yaml:"grpc"`
	Memcached Memcached `yaml:"memcached"`
	Storage   string    `yaml:"storage" env:"STORAGE"`
}

type GRPC struct {
	Port int `env-required:"true" yaml:"port" env:"GRPC_PORT"`
}

type Memcached struct {
	Addr string `yaml:"addr" env:"MEMCACHED_ADDR"`
}

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
