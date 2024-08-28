package main

import (
	"flag"
	"fmt"
	"yir/internal/log"
	"yir/medworkers/internal/apps"
	"yir/medworkers/internal/config"
	MedWorkerApi "yir/medworkers/internal/controller/medworkers"
	dbRepos "yir/medworkers/internal/repositories/repositories"
	MedWorkerUsecases "yir/medworkers/internal/usecases/medworkers"

	"go.uber.org/zap"
)

const (
	defaultCfgPath = "config/config.yml"
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

	MedWorkerRepo, err := dbRepos.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("medworker repo create: %w", err))
	}
	logger.Info("DB load")

	usecases := MedWorkerUsecases.NewMedWorkerUseCase(MedWorkerRepo, logger)

	MedWorkerGRPCController := MedWorkerApi.NewServer(usecases, logger)

	app := apps.New(MedWorkerGRPCController, logger)

	if err := app.Run(&cfg.App); err != nil {
		logger.Error("App error", zap.Error(err))
	}
}
