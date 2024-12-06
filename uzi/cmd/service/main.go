// TODO: убрать мусор отсюда сделать нормальную инициализацию
package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"yir/pkg/brokerlib"
	pkgconfig "yir/pkg/config"
	"yir/pkg/grpclib"
	"yir/pkg/loglib"
	"yir/uzi/internal/config"

	"yir/uzi/internal/repository"

	devicesrv "yir/uzi/internal/services/device"
	imagesrv "yir/uzi/internal/services/image"
	nodesrv "yir/uzi/internal/services/node"
	segmentsrv "yir/uzi/internal/services/segment"
	uzisrv "yir/uzi/internal/services/uzi"

	pb "yir/uzi/internal/generated/grpc/service"
	grpchandler "yir/uzi/internal/grpc"

	devicehandler "yir/uzi/internal/grpc/device"
	imagehandler "yir/uzi/internal/grpc/image"
	nodehandler "yir/uzi/internal/grpc/node"
	segmenthandler "yir/uzi/internal/grpc/segment"
	uzihandler "yir/uzi/internal/grpc/uzi"

	uziprocessedsubscriber "yir/uzi/internal/broker/uziprocessed"
	uziuploadsubscriber "yir/uzi/internal/broker/uziupload"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
)

const (
	defaultCfgPath = "service.yml"
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

	db, err := sqlx.Open("postgres", cfg.DB.Dsn)
	if err != nil {
		slog.Error("init db", "err", err)
		return failExitCode
	}
	defer db.Close()

	client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Secure: false,
		Creds:  credentials.NewStaticV4(cfg.S3.Access_Token, cfg.S3.Secret_Token, ""),
	})
	if err != nil {
		slog.Error("init s3", "err", err)
		return failExitCode
	}

	if err := db.Ping(); err != nil {
		slog.Error("ping db", "err", err)
		return failExitCode
	}

	dao := repository.NewRepository(db, client, "uzi")

	deviceSrv := devicesrv.New(dao)
	uziSrv := uzisrv.New(dao)
	imageSrv := imagesrv.New(dao)
	nodeSrv := nodesrv.New(dao)
	serviceSrv := segmentsrv.New(dao)

	// grpc
	deviceHandler := devicehandler.New(deviceSrv)
	uziHandler := uzihandler.New(uziSrv)
	imageHandler := imagehandler.New(imageSrv)
	nodeHandler := nodehandler.New(nodeSrv)
	serviceHandler := segmenthandler.New(serviceSrv)

	handler := grpchandler.New(
		deviceHandler,
		uziHandler,
		imageHandler,
		nodeHandler,
		serviceHandler,
	)

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(grpclib.ServerCallLoggerInterceptor))
	pb.RegisterUziSrvServer(server, handler)

	// broker
	uziuploadSubscriber := uziuploadsubscriber.New(imageSrv)
	uziprocessedSubscriber := uziprocessedsubscriber.New(nodeSrv)

	uziuploadHandler, err := brokerlib.GetSubscriberHandler(
		uziuploadSubscriber,
		cfg.Broker.Addrs,
		nil,
	)
	if err != nil {
		slog.Error("create uzipload sub", "err", err)
		return failExitCode
	}

	uziprocessedHandler, err := brokerlib.GetSubscriberHandler(
		uziprocessedSubscriber,
		cfg.Broker.Addrs,
		nil,
	)
	if err != nil {
		slog.Error("create uziprocesse sub", "err", err)
		return failExitCode
	}

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	close := make(chan struct{})
	// ЛЮТОЕ MVP
	slog.Info("start serve", slog.String("app url", cfg.App.Url))
	go func() {
		if err := server.Serve(lis); err != nil {
			slog.Error("take port", "err", err)
			panic("serve grpc")
		}
		close <- struct{}{}
	}()
	go func() {
		// пока без DI
		if err := uziuploadHandler.Start(context.Background()); err != nil {
			slog.Error("start uziupload handler", "err", err)
		}
	}()
	go func() {
		if err := uziprocessedHandler.Start(context.Background()); err != nil {
			slog.Error("start uziprocessedHandler handler", "err", err)
		}
	}()

	<-close

	return successExitCode
}
