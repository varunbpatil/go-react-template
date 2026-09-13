// Package config provides the application configuration parsed from environment variables.
package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

// Config is the application configuration.
type Config struct {
	Log  LogConfig  `envPrefix:"LOG_"`
	User UserConfig `envPrefix:"USER_"`
	GRPC GRPCConfig `envPrefix:"GRPC_"`
	HTTP HTTPConfig `envPrefix:"HTTP_"`
}

// LogConfig is the log configuration.
type LogConfig struct {
	Level  string `env:"LEVEL"  envDefault:"info"`
	Format string `env:"FORMAT" envDefault:"text"`
}

// UserConfig is the Users domain configuration.
type UserConfig struct {
	UserProperty1 string `env:"PROPERTY1"`
}

type GRPCConfig struct {
	Address string `env:"ADDRESS" envDefault:":50051"`
}

type HTTPConfig struct {
	Address string `env:"ADDRESS" envDefault:":8080"`
}

func Parse() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("config: %w", err)
	}
	return &cfg, nil
}
