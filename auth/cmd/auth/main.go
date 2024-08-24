// курить над грейсфулом
package main

import (
	"flag"
	"fmt"
	"yir/auth/internal/apps"
	"yir/auth/internal/config"
	authApi "yir/auth/internal/controller/v0/auth"
	"yir/auth/internal/core/jwt"
	"yir/auth/internal/repositories/db/repositories"
	authUsecases "yir/auth/internal/usecases/auth"
	"yir/internal/log"

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

	jwtService, err := jwt.NewService(&cfg.Token, logger)
	if err != nil {
		panic(fmt.Errorf("jwt service create: %w", err))
	}

	authRepo, err := repositories.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("auth repo create: %w", err))
	}
	logger.Info("DB load")

	usecases := authUsecases.NewAuthUseCase(authRepo, jwtService, logger)

	authGRPCController := authApi.NewAuthServer(usecases)

	app := apps.New(authGRPCController, logger)

	if err := app.Run(&cfg.App); err != nil {
		logger.Error("Application error", zap.Error(err))
	}
}
