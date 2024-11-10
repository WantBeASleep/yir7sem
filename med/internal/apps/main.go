package apps

import (
	"context"
	"fmt"
	"net"
	"net/http"
	cardController "yir/med/internal/controller/card"
	medWorkerController "yir/med/internal/controller/medworkers"
	patientController "yir/med/internal/controller/patient"

	"sync"
	"yir/med/internal/repository/config"

	pb "yir/med/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	cardController      *cardController.Server
	medWorkerController *medWorkerController.Server
	patientController   *patientController.Server
	logger              *zap.Logger
}

func New(
	cardController *cardController.Server,
	medWorkerController *medWorkerController.Server,
	patientController *patientController.Server,
	logger *zap.Logger,
) *App {
	return &App{
		cardController:      cardController,
		medWorkerController: medWorkerController,
		patientController:   patientController,
		logger:              logger,
	}
}

func (a *App) Run(cfg *config.App) error {
	s := grpc.NewServer()
	pb.RegisterMedCardServer(s, a.cardController)
	pb.RegisterMedWorkersServer(s, a.medWorkerController)
	pb.RegisterMedPatientServer(s, a.patientController)

	mux := runtime.NewServeMux()
	ctx := context.Background()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterMedCardHandlerFromEndpoint(ctx, mux, cfg.Host+":"+cfg.GRPCPort, opts); err != nil {
		return fmt.Errorf("register card service HTTP handler: %w", err)
	}
	if err := pb.RegisterMedWorkersHandlerFromEndpoint(ctx, mux, cfg.Host+":"+cfg.GRPCPort, opts); err != nil {
		return fmt.Errorf("register medworker service HTTP handler: %w", err)
	}
	if err := pb.RegisterMedPatientHandlerFromEndpoint(ctx, mux, cfg.Host+":"+cfg.GRPCPort, opts); err != nil {
		return fmt.Errorf("register patient service HTTP handler: %w", err)
	}

	grpcLis, err := net.Listen("tcp", cfg.Host+":"+cfg.GRPCPort)
	if err != nil {
		return fmt.Errorf("failed to listen on gRPC port: %w", err)
	}

	httpLis, err := net.Listen("tcp", cfg.Host+":"+cfg.HTTPPort)
	if err != nil {
		return fmt.Errorf("failed to listen on HTTP port: %w", err)
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		a.logger.Info("Starting gRPC server...")
		if err := s.Serve(grpcLis); err != nil {
			errChan <- err
		}
	}()

	go func() {
		defer wg.Done()
		a.logger.Info("Starting HTTP server...")
		if err := http.Serve(httpLis, mux); err != nil {
			errChan <- err
		}
	}()

	a.logger.Info("gRPC and HTTP servers are running")
	wg.Wait()
	s.Stop()
	a.logger.Info("gRPC and HTTP servers stopped")

	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}
