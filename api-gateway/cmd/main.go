package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	app "yir/api-gateway/internal/app/gateway"
	"yir/api-gateway/internal/config"
	"yir/api-gateway/internal/controller"
	"yir/api-gateway/internal/logger"

	"go.uber.org/zap"
)

var (
	cfg_defaultPath = "config/config.yaml"
)

func main() {
	logger.Init()
	defer logger.Logger.Sync()

	cfg := config.MustLoad(cfg_defaultPath)
	logger.Logger.Info("Configuration loaded successfully")

	// подключение к сервису {1..N} по grpc clientу

	// Инициализация гейтвея и запуск
	s := app.New(cfg.Gateway.Host, cfg.Gateway.Port, controller.InitRouter())
	go s.Run()
	logger.Logger.Info("Gateway is starting",
		zap.String("host", cfg.Gateway.Host),
		zap.String("port", cfg.Gateway.Port),
	)

	// Gracefull Shutdown
	// TODO: надо еще завершать соединения с серверами
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	logger.Logger.Info("Received shutdown signal)",
		zap.String("signal", sig.String()),
	)
	// close {1..N} service connections
	logger.Logger.Info("Closed connection with service: Auth)")
	logger.Logger.Info("Closed connection with service: Med)")
	logger.Logger.Info("Closed connection with service: Uzi)")
	logger.Logger.Info("Closed connection with service: S3)")
	s.Shutdown(context.Background())
	logger.Logger.Info("Shutdown down Gateway...)")
}
