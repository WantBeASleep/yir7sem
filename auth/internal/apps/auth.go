package apps

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	pb "yir/auth/api/v0/auth"
	"yir/auth/internal/config"
	authApi "yir/auth/internal/controller/v0/auth"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Auth struct {
	controller *authApi.AuthServer

	logger *zap.Logger
}

func New(
	controller *authApi.AuthServer,
	logger *zap.Logger,
) *Auth {
	return &Auth{
		controller: controller,
		logger:     logger,
	}
}

func (a *Auth) Run(cfg *config.App) error {
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, a.controller)

	mux := runtime.NewServeMux()
	if err := pb.RegisterAuthHandlerFromEndpoint(context.TODO(), mux, cfg.Host+":"+cfg.GRPCPort, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
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
			a.logger.Error("GRPC server serve error", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()

	go func() {
		if err := http.ListenAndServe(":"+cfg.HTTPPort, mux); err != nil {
			a.logger.Error("HTTP server serve error", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()

	a.logger.Info("GRPC and HTTP servers start serve")
	wg.Wait()
	s.Stop()
	a.logger.Info("GRPC and HTTP servers end serve")

	select {
	case err := <-errFeedBack:
		return err

	default:
		return nil
	}
}
