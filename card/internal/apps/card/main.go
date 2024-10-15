package apps

import (
	"context"
	"fmt"
	"net"
	"net/http"
	cardcontroller "service/internal/controller/card"
	"service/internal/repository/config"
	"sync"

	pb "service/api/cards"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Card struct {
	controller *cardcontroller.Server
	logger     *zap.Logger
}

func New(controller *cardcontroller.Server, logger *zap.Logger) *Card {
	return &Card{controller: controller, logger: logger}
}

func (c *Card) Run(cfg *config.App) error {
	s := grpc.NewServer()
	pb.RegisterMedCardServer(s, c.controller)

	mux := runtime.NewServeMux()
	if err := pb.RegisterMedCardHandlerFromEndpoint(context.TODO(), mux, cfg.Host+":"+cfg.GRPCPort, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
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
			c.logger.Error("GRPC Server serve error", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()
	go func() {
		if err := http.ListenAndServe(":"+cfg.HTTPPort, mux); err != nil {
			c.logger.Error("HTTP SERVER SERVE ERROR", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()
	c.logger.Info("HTTP and gRPC servers are starting to serve")
	wg.Wait()
	s.Stop()
	c.logger.Info("HTTP and gRPC servers are ending to serve")

	select {
	case err := <-errFeedBack:
		return err

	default:
		return nil
	}
}
