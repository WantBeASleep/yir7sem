package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	app "yir/gateway/internal/app/gateway"
	"yir/gateway/internal/config"
	"yir/gateway/internal/controller"
	"yir/gateway/internal/logger"
	"yir/gateway/internal/service/auth"
	"yir/gateway/internal/service/s3"
	"yir/gateway/internal/service/uzi"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	cfg_defaultPath = "config/config.yaml"
)

func main() {
	logger.Init()
	defer logger.Logger.Sync()

	cfg := config.MustLoad(cfg_defaultPath)
	logger.Logger.Info("Configuration loaded successfully")

	// подключение к gRPC сервисам
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 1. Подключенеи к Auth
	authService := new(auth.Auth) // имя package и имя структуры должны быть разными, но пока так
	err := authService.Connect(cfg.Auth.Host, cfg.Auth.GRPCport, opts)
	if err != nil {
		logger.Logger.Error("failed connection to auth",
			zap.Error(err),
		)
	}
	defer authService.Close()

	// 2. Подключение к Med
	// МУСОРА СОСАТЬ МУСОРА СОСАТЬ, 3 ПРОТИКА СОСУТ ХУЙ В РОТИКА

	// 3. Подключение к Uzi
	uziService := new(uzi.Uzi)
	err = uziService.Connect(cfg.Uzi.Host, cfg.Uzi.GRPCport, opts)
	if err != nil {
		logger.Logger.Error("failed connection to uzi",
			zap.Error(err),
		)
	}
	defer uziService.Close()

	// 4. Подключение к S3
	s3Service := new(s3.S3)
	err = s3Service.Connect(cfg.S3.Host, cfg.S3.GRPCport, opts)
	if err != nil {
		logger.Logger.Error("failed connection to s3",
			zap.Error(err),
		)
	}

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
