package config

import (
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	Server Server
	Logger Logger
	Mysql  Mysql
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

type Mysql struct {
	Host     string `envconfig:"BAC_MYSQL_HOST" default:"localhost:3306"`
	Driver   string `envconfig:"BAC_MYSQL_DRIVER" default:"mysql"`
	User     string `envconfig:"BAC_MYSQL_USER" default:"root"`
	Password string `envconfig:"BAC_MYSQL_PASSWORD" default:"1234"`
	Database string `envconfig:"BAC_MYSQL_DATABASE" default:"board"`
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
