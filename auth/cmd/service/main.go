package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/WantBeASleep/goooool/grpclib"
	"github.com/WantBeASleep/goooool/loglib"

	"auth/internal/config"

	pkgconfig "github.com/WantBeASleep/goooool/config"

	"auth/internal/repository"

	loginsrv "auth/internal/services/login"
	passwordsrv "auth/internal/services/password"
	refreshsrv "auth/internal/services/refresh"
	registersrv "auth/internal/services/register"
	tokenizersrv "auth/internal/services/tokenizer"

	pb "auth/internal/generated/grpc/service"
	grpchandler "auth/internal/grpc"

	loginhadnler "auth/internal/grpc/login"
	refreshhadnler "auth/internal/grpc/refresh"
	registerhadnler "auth/internal/grpc/register"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
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
	cfg, err := pkgconfig.Load[config.Config]()
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

	loginHadnler, err := loginhadnler.New(loginSrv)
	if err != nil {
		slog.Error("init loginHandler: %v", err) //TODO: проверить правильно ли вообще тут начинать возвращать ошибку+ slog - это какая-то модификация обычного log
		return exitCode
	}
	refreshHadnler, err := refreshhadnler.New(refreshSrv)
	if err != nil {
		slog.Error("init refreshHandler: %v", err)
		return exitCode
	}
	registerHadnler, err := registerhadnler.New(registerSrv)
	if err != nil {
		slog.Error("init registerHandler: %v", err)
		return exitCode
	}

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
