package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	App struct {
		Env    string `yaml:"env" env-default:"LOCAL"`
		Server struct {
			Host     string        `yaml:"host" env-required:"true"`
			GRPCPort int           `yaml:"grpc-port" env-required:"true"`
			HTTPPort int           `yaml:"http-port" env-required:"true"`
			Timeout  time.Duration `yaml:"timeout" env-default:"5s"`
		} `yaml:"server"`
	} `yaml:"app"`

	Database struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     int    `yaml:"port" env-required:"true"`
		Name     string `yaml:"name" env-required:"true"`
		User     string `yaml:"user" env-default:"postgres"`
		Password string `yaml:"password" env-default:"postgres"`

		path string `yaml:"storage-path" env-required:"true"`
	} `yaml:"database"`
}

func NewConfig(filePath string) (*Config, error) {
	if filePath == "" {
		filePath = os.Getenv("CONFIG_PATH")
	}
	if filePath == "" {
		return nil, fmt.Errorf("no config file path provided")
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s does not exist", filePath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	return &cfg, nil
}
