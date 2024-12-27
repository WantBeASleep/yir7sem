// TODO: убрать мусор отсюда сделать нормальную инициализацию
package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"github.com/WantBeASleep/goooool/brokerlib"
	"github.com/WantBeASleep/goooool/grpclib"
	"github.com/WantBeASleep/goooool/loglib"

	pkgconfig "github.com/WantBeASleep/goooool/config"

	"uzi/internal/config"

	"uzi/internal/repository"

	devicesrv "uzi/internal/services/device"
	imagesrv "uzi/internal/services/image"
	nodesrv "uzi/internal/services/node"
	segmentsrv "uzi/internal/services/segment"
	uzisrv "uzi/internal/services/uzi"

	pb "uzi/internal/generated/grpc/service"
	grpchandler "uzi/internal/grpc"

	devicehandler "uzi/internal/grpc/device"
	imagehandler "uzi/internal/grpc/image"
	nodehandler "uzi/internal/grpc/node"
	segmenthandler "uzi/internal/grpc/segment"
	uzihandler "uzi/internal/grpc/uzi"

	uziprocessedsubscriber "uzi/internal/subs/uziprocessed"
	uziuploadsubscriber "uzi/internal/subs/uziupload"

	adapters "uzi/internal/adapters"
	brokeradapter "uzi/internal/adapters/broker"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

	producer, err := brokerlib.NewProducer(cfg.Broker.Addrs)
	if err != nil {
		slog.Error("init broker producer", slog.Any("err", err))
	}

	dao := repository.NewRepository(db, client, "uzi")
	adapter := adapters.New(brokeradapter.New(producer))

	deviceSrv := devicesrv.New(dao)
	uziSrv := uzisrv.New(dao)
	imageSrv := imagesrv.New(dao, adapter)
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

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpclib.ServerCallPanicRecoverInterceptor,
			grpclib.ServerCallLoggerInterceptor,
		),
	)
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
