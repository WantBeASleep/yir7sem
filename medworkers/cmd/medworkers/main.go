package main

import (
	"context"
	"flag"
	"fmt"
	"service/internal/apps"
	"service/internal/config"
	MedWorkerApi "service/internal/controller/medworkers"
	dbRepos "service/internal/repositories/repositories"
	"service/internal/services"
	MedWorkerUsecases "service/internal/usecases/medworkers"
	"service/pkg_log/log"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Создаем соединение с gRPC сервисом
	conn, err := grpc.DialContext(ctx, cfg.PatientService.Host+":"+cfg.PatientService.GRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("failed to connect to patient service", zap.Error(err))
	}
	defer conn.Close()

	// Инициализируем gRPC клиент для сервиса пациентов
	patientService := services.NewGRPCPatientService(conn)

	MedWorkerRepo, err := dbRepos.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("medworker repo create: %w", err))
	}
	logger.Info("DB load")

	usecases := MedWorkerUsecases.NewMedWorkerUseCase(MedWorkerRepo, patientService, logger)

	MedWorkerGRPCController := MedWorkerApi.NewServer(usecases, logger)

	app := apps.New(MedWorkerGRPCController, logger)

	if err := app.Run(&cfg.App); err != nil {
		logger.Error("App error", zap.Error(err))
	}
}
