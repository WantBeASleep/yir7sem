package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"sync"
	"yir/pkg/log"
	pb "yir/uzi/api"
	uziapi "yir/uzi/internal/api/uzi"
	"yir/uzi/internal/config"
	"yir/uzi/internal/db/uzirepo"
	uziusecase "yir/uzi/internal/usecases/uzi"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

// MVP CODE Tech Debt
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

	uziRepo, err := uzirepo.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("init repo: %w", err))
	}

	uziUseCase := uziusecase.NewUziUseCase(uziRepo, logger)

	uziController, err := uziapi.NewServer(uziUseCase)
	if err != nil {
		panic(fmt.Errorf("init ctrl: %w", err))
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUziAPIServer(grpcServer, uziController)

	mux := runtime.NewServeMux()
	if err := pb.RegisterUziAPIHandlerFromEndpoint(context.TODO(), mux, cfg.App.Host+":"+cfg.App.GRPCPort, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		panic(fmt.Errorf("register http handlers: %w", err))
	}

	GRPCLis, err := net.Listen("tcp", cfg.App.Host+":"+cfg.App.GRPCPort)
	if err != nil {
		panic(fmt.Errorf("lister grpc host:port: %w", err))
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if err := grpcServer.Serve(GRPCLis); err != nil {
			logger.Error("GRPC server serve error", zap.Error(err))
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if err := http.ListenAndServe(":"+cfg.App.HTTPPort, mux); err != nil {
			logger.Error("HTTP server serve error", zap.Error(err))
		}
		wg.Done()
	}()

	wg.Wait()

	logger.Info("GRPC and HTTP servers end serve")
}
