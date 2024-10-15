package main

import (
	"flag"
	"fmt"
	apps "service/internal/apps/card"
	controller "service/internal/controller/card"
	"service/internal/repository"
	"service/internal/repository/config"
	"service/internal/usecase"
	"service/pkg_log/log"

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
	CardRepo, err := repository.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("card repo create: %w", err))
	}
	logger.Info("DB load")
	usecases := usecase.NewCardUseCase(CardRepo, logger)
	CardGRPCController := controller.NewServer(usecases, logger)
	app := apps.New(CardGRPCController, logger)

	if err := app.Run(&cfg.App); err != nil {
		logger.Error("App error", zap.Error(err))
	}
}
