package main

import (
	"flag"
	"fmt"
	"yir/internal/log"
	apps "yir/med/internal/apps/patient"
	controller "yir/med/internal/controller/patient"
	"yir/med/internal/repository"
	"yir/med/internal/repository/config"
	"yir/med/internal/usecase"

	"go.uber.org/zap"
)

const (
	defaultCfgPath = "config/config.yaml"
	shorthand      = " (shorthand)"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", defaultCfgPath, "set config path")
	flag.StringVar(&configPath, "c", defaultCfgPath, "set config path"+shorthand)
}

func main() {
	flag.Parse()
	cfg := config.MustLoad(configPath)

	var logger *zap.Logger
	if cfg.App.Env == "DEV" {
		logger = log.New(log.DevEnv, "")
	} else {
		panic("not supported")
	}
	logger.Info("Cfg && logger load")

	PatientRepo, err := repository.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("patient repo create: %w", err))
	}
	logger.Info("DB load")

	usecases := usecase.NewPatientUseCase(*PatientRepo, logger)

	server := controller.NewServer(usecases)

	app := apps.New(server, logger)

	if err := app.Run(&cfg.App); err != nil {
		logger.Error("App error", zap.Error(err))
	}
}
