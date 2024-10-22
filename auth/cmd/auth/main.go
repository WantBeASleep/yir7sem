package main

import (
	"flag"
	"fmt"
	"service/auth/internal/apps"
	"service/auth/internal/config"
	authApi "service/auth/internal/controller/auth"
	"service/auth/internal/core/jwt"
	dbRepos "service/auth/internal/repositories/db/repositories"
	serviceRepos "service/auth/internal/repositories/services"
	authUsecases "service/auth/internal/usecases/auth"
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

	// medRepo := serviceRepos.NewService()
	medServiceAddress := fmt.Sprintf("%s:%s", cfg.MedService.Host, cfg.MedService.GRPCPort)
	medRepo, err := serviceRepos.NewService(medServiceAddress)
	if err != nil {
		panic(fmt.Errorf("failed to create med service client: %w", err))
	}

	usecases := authUsecases.NewAuthUseCase(authRepo, medRepo, jwtService, logger)

	authGRPCController := authApi.NewServer(usecases)

	app := apps.New(authGRPCController, logger)

	if err := app.Run(&cfg.App); err != nil {
		logger.Error("Application error", zap.Error(err))
	}
}
