package apps

import (
	"context"
	"fmt"
	"net"
	"net/http"
	pb "service/api/medworkers"
	"service/internal/config"
	MedWorkerApi "service/internal/controller/medworkers"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MedWorker struct {
	controller *MedWorkerApi.Server
	logger     *zap.Logger
}

func New(
	controller *MedWorkerApi.Server,
	logger *zap.Logger,
) *MedWorker {
	return &MedWorker{controller: controller, logger: logger}
}

func (a *MedWorker) Run(cfg *config.App) error {
	s := grpc.NewServer()
	pb.RegisterMedWorkersServer(s, a.controller)

	mux := runtime.NewServeMux()
	if err := pb.RegisterMedWorkersHandlerFromEndpoint(context.TODO(), mux, cfg.Host+":"+cfg.GRPCPort, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		return fmt.Errorf("register http handlers: %w", err)
	}
	GRPCLis, err := net.Listen("tcp", cfg.Host+":"+cfg.GRPCPort)
	if err != nil {
		return fmt.Errorf("listen grpc host:port: %w", err)
	}

	var wg sync.WaitGroup
	errFeedBack := make(chan error, 2)

	wg.Add(1)

	go func() {
		if err := s.Serve(GRPCLis); err != nil {
			a.logger.Error("GRPC Server serve error", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()
	go func() {
		if err := http.ListenAndServe(":"+cfg.HTTPPort, mux); err != nil {
			a.logger.Error("HTTP SERVER SERVE ERROR", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()

	a.logger.Info("HTTP and gRPC servers are starting to serve")
	wg.Wait()
	s.Stop()
	a.logger.Info("HTTP and gRPC servers are ending to serve")

	select {
	case err := <-errFeedBack:
		return err

	default:
		return nil
	}
}
