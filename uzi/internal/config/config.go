package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App App `yaml:"app"`
	DB  DB  `yaml:"db"`

	Services Services `yaml:"services"`
}

type App struct {
	Env      string `yaml:"env" env:"ENV" env-default:"PROD"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	GRPCPort string `yaml:"grpc" env:"GRPC_PORT" env-default:"50055"`
	HTTPPort string `yaml:"http" env:"HTTP_PORT" env-default:"8080"`
}

type DB struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Name     string `yaml:"name" env:"DB_NAME" env-required:"true"`
	User     string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASS" env-required:"true"`
}

type Services struct {
	S3ServiceHost string `yaml:"s3servicehost" env:"SERVICE_HOST_S3" env-default:"localhost:50055"`
}

func (d *DB) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		d.Host,
		d.User,
		d.Password,
		d.Name,
		d.Port,
	)
}

func MustLoad(cfgPath string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		panic(fmt.Errorf("read config: %v", err))
	}

	return &cfg
}
