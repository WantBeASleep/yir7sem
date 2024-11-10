package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App App `yaml:"app"`
	DB  DB  `yaml:"db"`
}

type App struct {
	Env      string `yaml:"env" env:"ENV" env-default:"PROD"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	GRPCPort string `yaml:"grpc" env:"GRPC_PORT" env-default:"50053"`
	HTTPPort string `yaml:"http" env:"HTTP_PORT" env-default:"8082"`
}

type DB struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Name     string `yaml:"name" env:"DB_NAME" env-required:"true"`
	User     string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASS" env-required:"true"`
}

func MustLoad(cfgPath string) *Config {
	var cfg Config
	fmt.Println(cfgPath)
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		panic(fmt.Errorf("read config: %v", err))
	}
	return &cfg
}
