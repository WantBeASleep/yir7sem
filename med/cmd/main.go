package main

import (
	"flag"
	"fmt"
	apps "yir/med/internal/apps"
	cardController "yir/med/internal/controller/card"
	workerController "yir/med/internal/controller/medworkers"

	patientController "yir/med/internal/controller/patient"

	patientRep "yir/med/internal/repository"
	cardRep "yir/med/internal/repository/card"
	"yir/med/internal/repository/config"
	medworkerRep "yir/med/internal/repository/medworkers"
	"yir/med/internal/usecase"
	"yir/pkg_log/log"

	"go.uber.org/zap"
)

const (
	defaultCfgPath = "all/config/config.yml"
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
	CardRepo, err := cardRep.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("card repo create: %w", err))
	}
	WorkerRepo, err := medworkerRep.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("worker repo create: %w", err))
	}
	PatientRepo, err := patientRep.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("patient repo create: %w", err))
	}
	logger.Info("DB load")

	cardusecases := usecase.NewCardUseCase(CardRepo, logger)
	workerusecases := usecase.NewMedWorkerUseCase(WorkerRepo, logger)
	patientusecases := usecase.NewPatientUseCase(*PatientRepo, logger)

	CardGRPCController := cardController.NewServer(cardusecases, logger)
	MedworkersGRPCController := workerController.NewServer(workerusecases, logger)
	PatientGRPCController := patientController.NewServer(patientusecases)

	app := apps.New(CardGRPCController, MedworkersGRPCController, PatientGRPCController, logger)
	if err := app.Run(&cfg.App); err != nil {
		logger.Error("App error", zap.Error(err))
	}
}
