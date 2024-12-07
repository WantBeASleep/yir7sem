package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

func Load[T any]() (*T, error) {
	cfg := new(T)
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("read config to cfg: %w", err)
	}

	return cfg, nil
}
