package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Gateway Gateway `yaml:"gateway"`
	Auth    Auth    `yaml:"auth"`
	Token   TokenValidation
	Med     Med   `yaml:"med"`
	Uzi     Uzi   `yaml:"uzi"`
	S3      S3    `yaml:"s3"`
	Kafka   Kafka `yaml:"kafka"`
}

type Gateway struct {
	Env            string        `yaml:"env" env-default:"DEV"`
	Host           string        `yaml:"host" env-default:"localhost"`
	Port           string        `yaml:"http" env-default:"8080"`
	RequestTimeout time.Duration `yaml:"request_timeout" env-default:"5m"`
	TcpTimeout     time.Duration `yaml:"tcp_timeout" env-default:"5m"`
}

type Auth struct {
	Env      string `yaml:"env" env-default:"DEV"`
	Host     string `yaml:"host" env-default:"localhost"`
	GRPCport string `yaml:"grpc" env-default:"50055"`
	HTTPport string `yaml:"http" env-default:"8081"`
}

type TokenValidation struct {
	PublicKey string `env:"TOKEN_PUBLIC_KEY" env-required:"true"`
}

type Med struct {
	Env      string `yaml:"env" env-default:"DEV"`
	Host     string `yaml:"host" env-default:"localhost"`
	GRPCport string `yaml:"grpc" env-default:"50056"`
	HTTPport string `yaml:"http" env-default:"8082"`
}

type Uzi struct {
	Env      string `yaml:"env" env-default:"DEV"`
	Host     string `yaml:"host" env-default:"localhost"`
	GRPCport string `yaml:"grpc" env-default:"50057"`
	HTTPport string `yaml:"http" env-default:"8083"`
}

type S3 struct {
	Env      string `yaml:"env" env-default:"DEV"`
	Host     string `yaml:"host" env-default:"localhost"`
	GRPCport string `yaml:"grpc" env-default:"50058"`
}

type Kafka struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"9092"`
}

func MustLoad(path string) Config {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(fmt.Errorf("read config failed: %w", err))
	}
	return cfg
}
