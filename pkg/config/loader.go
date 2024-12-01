package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func Load[T any](cfgPath string) (*T, error) {
	cfg := new(T)
	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		return nil, fmt.Errorf("read config to cfg: %w", err)
	}

	return cfg, nil
}
