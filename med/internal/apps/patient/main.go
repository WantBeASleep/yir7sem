package apps

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	patientctrl "yir/med/internal/controller/patient"
	pb "yir/med/internal/pb/patient"
	"yir/med/internal/repository/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Patient struct {
	controller *patientctrl.Server
	logger     *zap.Logger
}

func New(controller *patientctrl.Server, logger *zap.Logger) *Patient {
	return &Patient{
		controller: controller,
		logger:     logger,
	}
}

func (p *Patient) Run(cfg *config.App) error {
	s := grpc.NewServer()
	pb.RegisterMedPatientServer(s, p.controller)
	mux := runtime.NewServeMux()
	if err := pb.RegisterMedPatientHandlerFromEndpoint(context.TODO(), mux, cfg.Host+":"+cfg.GRPCPort, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		return fmt.Errorf("register http handlers: %w", err)
	}
	GRPCLis, err := net.Listen("tcp", cfg.Host+":"+cfg.GRPCPort)
	if err != nil {
		return fmt.Errorf("lister grpc host:port: %w", err)
	}
	var wg sync.WaitGroup
	errFeedBack := make(chan error, 2)

	wg.Add(1)

	go func() {
		if err := s.Serve(GRPCLis); err != nil {
			p.logger.Error("GRPC server serve error", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()

	go func() {
		if err := http.ListenAndServe(":"+cfg.HTTPPort, mux); err != nil {
			p.logger.Error("HTTP server serve error", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()
	p.logger.Info("GRPC and HTTP servers start serve")
	wg.Wait()
	s.Stop()
	p.logger.Info("GRPC and HTTP servers end serve")

	select {
	case err := <-errFeedBack:
		return err

	default:
		return nil
	}
}
