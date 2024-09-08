package main

import (
	"yir/uzi/internal/config"
	"yir/uzi/internal/repositories/db/repositories"
)

func main() {
	cfg := config.MustLoad("config/config.yaml")
	_, err := repositories.NewRepository(&cfg.DB)
	if err != nil {
		panic(err)
	}
}
