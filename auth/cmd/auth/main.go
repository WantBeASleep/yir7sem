package main

import (
	"flag"
	"fmt"
	"yir/auth/internal/apps"
	"yir/auth/internal/config"
	authApi "yir/auth/internal/controller/auth"
	"yir/auth/internal/core/jwt"
	dbRepos "yir/auth/internal/repositories/db/repositories"
	serviceRepos "yir/auth/internal/repositories/services"
	authUsecases "yir/auth/internal/usecases/auth"
	"yir/pkg/log"

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

// подумать над неймингом, здесь просто костыль
func main() {
	flag.Parse()
	cfg := config.MustLoad(configPath)

	var logger *zap.Logger
	if cfg.App.Env == "DEV" {
		logger = log.New(log.DevEnv, "")
	} else {
		panic("not support")
	}
	logger.Info("CFG && logger load")

	jwtService, err := jwt.NewService(&cfg.Token)
	if err != nil {
		panic(fmt.Errorf("jwt service create: %w", err))
	}

	authRepo, err := dbRepos.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("auth repo create: %w", err))
	}
	logger.Info("DB load")

	medRepo := serviceRepos.NewService()

	usecases := authUsecases.NewAuthUseCase(authRepo, medRepo, jwtService, logger)

	authGRPCController := authApi.NewServer(usecases)

	app := apps.New(authGRPCController, logger)

	if err := app.Run(&cfg.App); err != nil {
		logger.Error("Application error", zap.Error(err))
	}
}
