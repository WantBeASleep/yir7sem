package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	app "yir/gateway/internal/app/gateway"
	"yir/gateway/internal/config"
	"yir/gateway/internal/controller"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/middleware"
	"yir/gateway/internal/pb/authpb"
	"yir/gateway/internal/pb/medpb"
	"yir/gateway/internal/pb/uzipb"
	"yir/gateway/internal/service"
	"yir/gateway/internal/service/medservice"
	"yir/gateway/internal/usecase/secret"
	"yir/gateway/repository"
	pbs3 "yir/s3upload/api"
	"yir/s3upload/pkg/client"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	cfg_defaultPath = "config/config.yaml"
)

func main() {
	custom.Init()
	defer custom.Logger.Sync()

	cfg := config.MustLoad(cfg_defaultPath)
	custom.Logger.Info("Successfully loaded configuration")

	// подключение к gRPC сервисам
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 1. Подключенеи к Auth
	authConn, err := grpc.NewClient(cfg.Auth.Host+":"+cfg.Auth.GRPCport, opts...)
	if err != nil {
		custom.Logger.Fatal("failed connect to Auth",
			zap.Error(err),
		)
	}
	defer func() {
		custom.Logger.Info("Closed connection with service: Auth)")
		authConn.Close()
	}()
	authService := &service.AuthService{
		Client: authpb.NewAuthClient(authConn),
	}
	custom.Logger.Info("Successfully connected to Auth")

	// 2. Подключение к Med
	medConn, err := grpc.NewClient(cfg.Med.Host+":"+cfg.Med.GRPCport, opts...)
	if err != nil {
		custom.Logger.Fatal("failed connect to Med",
			zap.Error(err),
		)
	}
	defer func() {
		custom.Logger.Info("Closed connection with service: Med)")
		medConn.Close()
	}()
	medService := &medservice.MedService{
		CardClient:    medpb.NewMedCardClient(medConn),
		WorkerClient:  medpb.NewMedWorkersClient(medConn),
		PatientClient: medpb.NewMedPatientClient(medConn),
	}
	custom.Logger.Info("Successfully connected to Med")

	// 3. Подключение к Uzi
	uziConn, err := grpc.NewClient(cfg.Uzi.Host+":"+cfg.Uzi.GRPCport, opts...)
	if err != nil {
		custom.Logger.Fatal("failed connect to Uzi")
	}
	defer func() {
		custom.Logger.Info("Closed connection with service: Uzi)")
		uziConn.Close()
	}()
	uziService := &service.UziService{
		Client: uzipb.NewUziAPIClient(uziConn),
	}
	custom.Logger.Info("Successfully connected to Uzi")

	// 4. Подключение к S3
	s3Conn, err := grpc.NewClient(cfg.Gateway.Host+":"+cfg.S3.GRPCport, opts...)
	if err != nil {
		custom.Logger.Fatal("failed connect to S3",
			zap.Error(err),
		)
	}
	defer func() {
		custom.Logger.Info("Closed connection with service: S3)")
		s3Conn.Close()
	}()

	s3Client := &repository.S3Repo{
		S3: client.NewS3Client(pbs3.NewS3Client(s3Conn)),
	}
	custom.Logger.Info("Successfully connected to S3")

	// Конфиг для кафки сгружаем
	Producer := repository.New([]string{cfg.Kafka.Host + ":" + cfg.Kafka.Port}, "uzi_upload")

	key, err := secret.LoadPublicKey()
	if err != nil {
		custom.Logger.Fatal(
			"Didn't get pubkey",
			zap.Error(err),
		)
	}
	authMiddleware := &middleware.AuthMiddleware{
		PubKey: key,
	}
	// Gateway: Init Routes, Run app
	s := app.New(cfg.Gateway, controller.InitRouter(authService, medService, uziService, s3Client, authMiddleware, Producer))
	go func() {
		err := s.Run()
		if err != nil {
			custom.Logger.Error(
				"failed to run gateway",
				zap.Error(err),
			)
		}
	}()

	custom.Logger.Info("Gateway is starting",
		zap.String("host", cfg.Gateway.Host),
		zap.String("port", cfg.Gateway.Port),
	)
	// Gracefull Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	custom.Logger.Info("Received shutdown signal)",
		zap.String("signal", sig.String()),
	)
	s.Shutdown(context.Background())
	custom.Logger.Info("Shutdown down Gateway...)")
}
