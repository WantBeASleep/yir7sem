package main

import (
	"flag"
	"fmt"
	"service/auth/internal/apps"
	"service/auth/internal/config"
	authApi "service/auth/internal/controller/auth"
	"service/auth/internal/core/jwt"
	dbRepos "service/auth/internal/repositories/db/repositories"
	serviceRepos "service/auth/internal/repositories/services"
	authUsecases "service/auth/internal/usecases/auth"
	"service/pkg_log/log"

	"go.uber.org/zap"
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

// подумать над неймингом, здесь просто костыль
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

	jwtService, err := jwt.NewService(&cfg.Token)
	if err != nil {
		panic(fmt.Errorf("jwt service create: %w", err))
	}

	authRepo, err := dbRepos.NewRepository(&cfg.DB)
	if err != nil {
		panic(fmt.Errorf("auth repo create: %w", err))
	}
	logger.Info("DB load")

	// medRepo := serviceRepos.NewService()
	medServiceAddress := fmt.Sprintf("%s:%s", cfg.MedService.Host, cfg.MedService.GRPCPort)
	medRepo, err := serviceRepos.NewService(medServiceAddress)
	if err != nil {
		panic(fmt.Errorf("failed to create med service client: %w", err))
	}

	usecases := authUsecases.NewAuthUseCase(authRepo, medRepo, jwtService, logger)

	authGRPCController := authApi.NewServer(usecases)

	app := apps.New(authGRPCController, logger)

	if err := app.Run(&cfg.App); err != nil {
		logger.Error("Application error", zap.Error(err))
	}
}

/*
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
*/
