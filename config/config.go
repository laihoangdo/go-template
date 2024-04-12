package config

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	Server ServerConfig
	Mysql  MysqlConfig
	Logger string
}

type ServerConfig struct {
	Port         string        `env:"PORT" envDefault:"8080"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"10s"`
	writeTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"10s"`
}

type MysqlConfig struct {
	Host     string `env:"MYSQL_HOST" envDefault:"127.0.0.1"`
	Port     string `env:"MYSQL_PORT" envDefault:"3306"`
	Username string `env:"MYSQL_USERNAME" envDefault:"root"`
	Password string `env:"MYSQL_PASSWORD" envDefault:"mysql"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
