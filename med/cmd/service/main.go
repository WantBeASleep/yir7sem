package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/WantBeASleep/goooool/grpclib"
	"github.com/WantBeASleep/goooool/loglib"

	"med/internal/config"

	pkgconfig "github.com/WantBeASleep/goooool/config"

	"med/internal/repository"

	cardsrv "med/internal/services/card"
	doctorsrv "med/internal/services/doctor"
	patientsrv "med/internal/services/patient"

	pb "med/internal/generated/grpc/service"
	grpchandler "med/internal/grpc"

	cardhandler "med/internal/grpc/card"
	doctorhandler "med/internal/grpc/doctor"
	patienthandler "med/internal/grpc/patient"

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

	patientSrv := patientsrv.New(dao)
	doctorSrv := doctorsrv.New(dao)
	cardSrv := cardsrv.New(dao)

	patientHandler := patienthandler.New(patientSrv)
	doctorHandler := doctorhandler.New(doctorSrv)
	cardHandler := cardhandler.New(cardSrv)

	handler := grpchandler.New(
		patientHandler,
		doctorHandler,
		cardHandler,
	)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpclib.ServerCallPanicRecoverInterceptor,
			grpclib.ServerCallLoggerInterceptor,
		),
	)
	pb.RegisterMedSrvServer(server, handler)

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
