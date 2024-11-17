package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Gateway Gateway `yaml:"gateway"`
	Auth    Auth    `yaml:"auth"`
	Uzi     Uzi     `yaml:"uzi"`
}

type Gateway struct {
	Env            string        `yaml:"env" env-default:"DEV"`
	Host           string        `yaml:"host" env-default:"localhost"`
	GRPCport       string        `yaml:"grpc" env-default:"50055"`
	HTTPport       string        `yaml:"http" env-default:"8080"`
	RequestTimeout time.Duration `yaml:"request_timeout" env-default:"5m"`
	TcpTimeout     time.Duration `yaml:"tcp_timeout" env-default:"5m"`
}

type Auth struct {
	Url string `yaml:"url" env-default:"localhost:50055"`
}

type Uzi struct {
	Url string `yaml:"url" env-default:"localhost:50055"`
}

func MustLoad(path string) Config {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(fmt.Errorf("read config failed: %w", err))
	}
	return cfg
}
