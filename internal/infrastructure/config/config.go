// Package config contains application configuration loading logic.
package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config holds the application configuration.
type Config struct {
	Env    string     `yaml:"env" env:"ENV" env-default:"local"`
	Random bool       `yaml:"random" env:"RANDOM" env-default:"false"`
	GPRC   GRPCConfig `yaml:"grpc"`
}

// GRPCConfig holds gRPC server configuration.
type GRPCConfig struct {
	Port    int           `yaml:"port" env:"GRPC_PORT"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT"`
}

// MustLoad loads the configuration or panics on error.
func MustLoad() *Config {
	configPath := fetchConfigPath()

	var cfg Config

	if configPath != "" {
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			panic("config file does not exist: " + configPath)
		}

		if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
			panic("failed to read config: " + err.Error())
		}
	} else {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			panic("failed to read config from env: " + err.Error())
		}
	}

	return &cfg
}

// fetchConfigPath retrieves the config file path from flag or environment.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
