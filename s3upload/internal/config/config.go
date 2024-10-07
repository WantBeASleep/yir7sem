package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App App `yaml:"app"`
	S3  S3  `yaml:"db"`
}

type App struct {
	Env      string `yaml:"env" env:"ENV" env-default:"PROD"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	GRPCPort string `yaml:"grpc" env:"GRPC_PORT" env-default:"50055"`
}

type S3 struct {
	Endpoint     string `env:"S3_ENDPOINT" env-required:"true"`
	AccessToken string `env:"S3_ACCESS_TOKEN" env-required:"true"`
	SecretToken string `env:"S3_SECRET_TOKEN" env-required:"true"`
}

// вынести в pkg
func MustLoad(cfgPath string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		panic(fmt.Errorf("read config: %v", err))
	}

	return &cfg
}
