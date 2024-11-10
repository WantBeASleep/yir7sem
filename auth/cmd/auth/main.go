package main

import (
	"flag"
	"fmt"
	"net"
	authApi "yir/auth/internal/api/auth"
	"yir/auth/internal/config"
	pb "yir/auth/api/auth"
	dbRepos "yir/auth/internal/db/repositories"
	"yir/auth/internal/entity"
	serviceRepos "yir/auth/internal/medservice"
	authUsecases "yir/auth/internal/usecases/auth"
	"yir/pkg_log/log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	jwtService, err := entity.NewService(&cfg.Token)
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

	authGRPCController, err := authApi.NewServer(usecases)
	if err != nil {
		panic(fmt.Errorf("work grepc controller: %w", err))
	}

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, authGRPCController)

	GRPCLis, err := net.Listen("tcp", cfg.App.Host+":"+cfg.App.GRPCPort)
	if err != nil {
		panic(fmt.Errorf("lister grpc host:port: %w", err))
	}

	if err := s.Serve(GRPCLis); err != nil {
		logger.Error("GRPC server serve error", zap.Error(err))
	}

}

/*

	

	mux := runtime.NewServeMux()
	if err := pb.RegisterAuthHandlerFromEndpoint(context.TODO(), mux, cfg.Host+":"+cfg.GRPCPort, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		return fmt.Errorf("register http handlers: %w", err)
	}

	go func() {
		if err := http.ListenAndServe(":"+cfg.HTTPPort, mux); err != nil {
			a.logger.Error("HTTP server serve error", zap.Error(err))
			errFeedBack <- err
		}
		wg.Done()
	}()
*/
