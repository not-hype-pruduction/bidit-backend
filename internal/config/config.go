package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string     `yaml:"env" env:"ENV" env-default:"local"`
	Random bool       `yaml:"random" env:"RANDOM" env-default:"false"`
	GPRC   GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env:"GRPC_PORT"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT"`
}

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

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
