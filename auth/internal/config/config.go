package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App        App        `yaml:"app"`
	DB         DB         `yaml:"db"`
	Token      Token      `yaml:"token"`
	MedService MedService `yaml:"medservice"`
}

type App struct {
	Env      string `yaml:"env" env:"ENV" env-default:"PROD"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	GRPCPort string `yaml:"grpc" env:"GRPC_PORT" env-default:"50053"`
	HTTPPort string `yaml:"http" env:"HTTP_PORT" env-default:"8081"`
}

type DB struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Name     string `yaml:"name" env:"DB_NAME" env-required:"true"`
	User     string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASS" env-required:"true"`
}

type Token struct {
	AccessLifeTime  time.Duration `yaml:"access_time" env:"TOKEN_AC_TIME" env-default:"5m"`
	RefreshLifeTime time.Duration `yaml:"refresh_time" env:"TOKEN_RT_TIME" env-default:"1h"`
	PrivateKey      string        `env:"TOKEN_PRIVATE_KEY" env-required:"true"`
	PublicKey       string        `env:"TOKEN_PUBLIC_KEY" env-required:"true"`
}

type MedService struct {
	Host     string `yaml:"host" env:"MEDSERVICE_HOST" env-required:"true"`
	GRPCPort string `yaml:"grpc" env:"MEDSERVICE_GRPC_PORT" env-default:"50054"`
}

func MustLoad(cfgPath string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		panic(fmt.Errorf("read config: %v", err))
	}

	return &cfg
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
