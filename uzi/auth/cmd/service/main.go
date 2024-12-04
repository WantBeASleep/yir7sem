package main

import (
	"log/slog"
	"net"
	"os"

	"yirv2/auth/internal/config"
	pkgconfig "yirv2/pkg/config"
	"yirv2/pkg/grpclib"
	"yirv2/pkg/loglib"

	"yirv2/auth/internal/repository"

	loginsrv "yirv2/auth/internal/services/login"
	passwordsrv "yirv2/auth/internal/services/password"
	refreshsrv "yirv2/auth/internal/services/refresh"
	registersrv "yirv2/auth/internal/services/register"
	tokenizersrv "yirv2/auth/internal/services/tokenizer"

	pb "yirv2/auth/internal/generated/grpc/service"
	grpchandler "yirv2/auth/internal/grpc"

	loginhadnler "yirv2/auth/internal/grpc/login"
	refreshhadnler "yirv2/auth/internal/grpc/refresh"
	registerhadnler "yirv2/auth/internal/grpc/register"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	defaultCfgPath = "./service.yml"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	loglib.InitLogger(loglib.WithDevEnv())
	cfg, err := pkgconfig.Load[config.Config](defaultCfgPath)
	if err != nil {
		slog.Error("init config", "err", err)
		return failExitCode
	}

	pubKey, privKey, err := cfg.ParseRsaKeys()
	if err != nil {
		slog.Error("parse rsa keys", "err", err)
		return failExitCode
	}

	db, err := sqlx.Open("postgres", cfg.DB.Dsn)
	if err != nil {
		slog.Error("init db", "err", err)
		return failExitCode
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("ping db", "err", err)
		return failExitCode
	}

	dao := repository.NewRepository(db)

	tokenizerSrv := tokenizersrv.New(
		cfg.JWT.AccessTokenTime,
		cfg.JWT.RefreshTokenTime,
		privKey,
		pubKey,
	)
	passwordSrv := passwordsrv.New()

	loginSrv := loginsrv.New(dao, passwordSrv, tokenizerSrv)
	refreshSrv := refreshsrv.New(dao, tokenizerSrv)
	registerSrv := registersrv.New(dao, passwordSrv)

	loginHadnler := loginhadnler.New(loginSrv)
	refreshHadnler := refreshhadnler.New(refreshSrv)
	registerHadnler := registerhadnler.New(registerSrv)

	handler := grpchandler.New(
		loginHadnler,
		refreshHadnler,
		registerHadnler,
	)

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(grpclib.ServerCallLoggerInterceptor))
	pb.RegisterAuthSrvServer(server, handler)

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	slog.Info("start serve", slog.String("app url", cfg.App.Url))
	if err := server.Serve(lis); err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	return successExitCode
}
