package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Gateway Gateway `yaml:"gateway"`
	Auth    Auth    `yaml:"auth"`
	Token   TokenValidation
	Med     Med `yaml:"med"`
	Uzi     Uzi `yaml:"uzi"`
	S3      S3  `yaml:"s3"`
}

type Gateway struct {
	Env  string `yaml:"env" env:"ENV" env-default:"DEV"`
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port string `yaml:"http" env:"HTTP_PORT" env-default:"8080"`
}

type Auth struct {
	Env      string `yaml:"env" env:"ENV" env-default:"DEV"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	GRPCport string `yaml:"grpc" env:"GRPC_PORT" env-default:"50055"`
	HTTPport string `yaml:"http" env:"HTTP_PORT" env-default:"8081"`
}

type TokenValidation struct {
	PublicKey string `env:"TOKEN_PUBLIC_KEY" env-required:"true"`
}

type Med struct {
	Env      string `yaml:"env" env:"ENV" env-default:"DEV"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	GRPCport string `yaml:"grpc" env:"GRPC_PORT" env-default:"50056"`
	HTTPport string `yaml:"http" env:"HTTP_PORT" env-default:"8082"`
}

type Uzi struct {
	Env      string `yaml:"env" env:"ENV" env-default:"DEV"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	GRPCport string `yaml:"grpc" env:"GRPC_PORT" env-default:"50057"`
	HTTPport string `yaml:"http" env:"HTTP_PORT" env-default:"8083"`
}

type S3 struct {
	Env      string `yaml:"env" env:"ENV" env-default:"DEV"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	GRPCport string `yaml:"grpc" env:"GRPC_PORT" env-default:"50058"`
}

func MustLoad(path string) Config {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(fmt.Errorf("read config failed: %w", err))
	}
	return cfg
}
