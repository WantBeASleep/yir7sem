package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	pb "yir/s3upload/api"
	uzicontroller "yir/s3upload/internal/api/s3upload"
	"yir/s3upload/internal/config"
	"yir/s3upload/internal/s3"
	uziusecase "yir/s3upload/internal/usecases/uzi"

	"yir/pkg/log"

	"go.uber.org/zap"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

const (
	defaultCfgPath = "config/config.yml"
	shorthand      = " (shorthand)"
)

var (
	configPath string

	bucketName string = "uzi"
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
		panic("not support")
	}
	logger.Info("CFG && logger load")

	s3Repo, err := s3.NewRepo(&cfg.S3, bucketName)
	if err != nil {
		panic(fmt.Errorf("init s3 repo"))
	}

	uziUseCase := uziusecase.NewUziUseCase(s3Repo, logger)
	uziContoller := uzicontroller.NewController(uziUseCase)

	grpcServer := grpc.NewServer()
	pb.RegisterS3Server(grpcServer, uziContoller)

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

	wg.Wait()

	logger.Info("GRPC and HTTP servers end serve")
}
