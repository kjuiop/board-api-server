package config

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Server Server
	Logger Logger
}

type Server struct {
	Mode           string `envconfig:"BAC_ENV" default:"dev"`
	Port           string `envconfig:"BAC_SERVER_PORT" default:"8090"`
	TrustedProxies string `envconfig:"BAC_TRUSTED_PROXIES" default:"127.0.0.1/32"`
}

type Logger struct {
	Level       string `envconfig:"BAC_LOG_LEVEL" default:"debug"`
	Path        string `envconfig:"BAC_LOG_PATH" default:"./logs/access.log"`
	PrintStdOut bool   `envconfig:"BAC_STDOUT" default:"true"`
}

func LoadEnvConfig() (*EnvConfig, error) {
	var config EnvConfig
	if err := envconfig.Process("bac", &config); err != nil {
		return nil, err
	}

	if err := config.CheckValid(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *EnvConfig) CheckValid() error {
	return nil
}
