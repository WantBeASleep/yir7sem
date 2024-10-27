package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	"yir/pkg/kafka"
	"yir/pkg/log"
	s3api "yir/s3upload/api"
	pb "yir/uzi/api"
	uziapi "yir/uzi/internal/api/uzi"
	"yir/uzi/internal/config"
	"yir/uzi/internal/db/uzirepo"
	"yir/uzi/internal/s3service"
	uziusecase "yir/uzi/internal/usecases/uzi"

	"github.com/IBM/sarama"

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

	s3serviceConn, err := grpc.Dial(cfg.Services.S3ServiceHost, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("conn to s3 service: %w", err))
	}
	s3serviceClient := s3api.NewS3UploadClient(s3serviceConn)
	s3Repo := s3service.NewS3Service(s3serviceClient)

	uziUseCase := uziusecase.NewUziUseCase(uziRepo, s3Repo, logger)

	uziController, err := uziapi.NewServer(uziUseCase)
	if err != nil {
		panic(fmt.Errorf("init ctrl: %w", err))
	}

	uziBroker := uziapi.NewBroker(logger, uziUseCase)

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

	wg.Add(1)
	go func() {
		config := sarama.NewConfig()
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
		config.Consumer.Group.Session.Timeout = 10 * time.Second
		config.Consumer.Group.Heartbeat.Interval = 3 * time.Second
		config.Consumer.Return.Errors = true

		
	}()

	wg.Wait()

	logger.Info("GRPC and HTTP servers end serve")
}
