package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "gateway/docs"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"

	"gateway/internal/config"
	"gateway/internal/repository"
	"gateway/pkg/brokerlib"
	pkgconfig "gateway/pkg/config"
	"gateway/pkg/grpclib"
	"gateway/pkg/loglib"

	brokeradapters "gateway/internal/adapters/broker"

	grpcadapters "gateway/internal/adapters/grpc"
	medgrpcadapter "gateway/internal/adapters/grpc/med"
	uzigrpcadapter "gateway/internal/adapters/grpc/uzi"

	medhandler "gateway/internal/api/med"
	uzihandler "gateway/internal/api/uzi"

	"gateway/internal/middleware"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	defaultCfgPath = "service.yml"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

// @title			Example API
// @version		1.0
// @description	This is a sample API for demonstration.
// @host			localhost:8080
// @BasePath		/api/v1
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

	pubKey, err := cfg.ParseRsaKeys()
	if err != nil {
		slog.Error("parse rsa key", "err", err)
		return failExitCode
	}

	client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Secure: false,
		Creds:  credentials.NewStaticV4(cfg.S3.Access_Token, cfg.S3.Secret_Token, ""),
	})
	if err != nil {
		slog.Error("init s3", "err", err)
		return failExitCode
	}

	dao := repository.NewRepository(client, "uzi")

	// TODO: обернуть в интерфейсы продьюсера/консьюмера

	producer, err := brokerlib.NewProducer(cfg.Broker.Addrs)
	if err != nil {
		slog.Error("init broker producer", "err", err)
		return failExitCode
	}

	brokeradapter := brokeradapters.New(producer)

	// TODO: поновыносить по папкам весь этот мусор
	medConn, err := grpc.NewClient(
		cfg.Adapters.MedUrl,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(grpclib.ClientCallLogger),
	)
	if err != nil {
		slog.Error("init medConn", "err", err)
		return failExitCode
	}

	uziConn, err := grpc.NewClient(
		cfg.Adapters.UziUrl,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(grpclib.ClientCallLogger),
	)
	if err != nil {
		slog.Error("init uziConn", "err", err)
		return failExitCode
	}

	medAdapter := medgrpcadapter.New(medConn)
	uziAdapter := uzigrpcadapter.New(uziConn)

	grpcadapter := grpcadapters.New(
		nil,
		medAdapter,
		uziAdapter,
	)

	medHandler := medhandler.New(grpcadapter)
	uziHandler := uzihandler.New(grpcadapter, brokeradapter, dao)
	// TODO: пробросить ошибки с логированием на верхнем уровне
	mdlwrs := middleware.New(pubKey)

	r := mux.NewRouter()

	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	apiRouter := r.PathPrefix("/api").Subrouter()

	medRouter := apiRouter.PathPrefix("/med").Subrouter()
	uziRouter := apiRouter.PathPrefix("/uzi").Subrouter()

	medRouter.Use(mdlwrs.Log, mdlwrs.Jwt)
	uziRouter.Use(mdlwrs.Log, mdlwrs.Jwt)

	medRouter.HandleFunc("/doctors", medHandler.GetDoctor).Methods("GET")

	uziRouter.HandleFunc("/echographics/{id}", uziHandler.PatchEchographics).Methods("PATCH")

	uziRouter.HandleFunc("/images/{id}/nodes-segments", uziHandler.GetUziNodeSegments).Methods("GET")
	uziRouter.HandleFunc("/uzis/{id}/images", uziHandler.GetUziImages).Methods("GET")
	uziRouter.HandleFunc("/uzis/{id}", uziHandler.GetUzi).Methods("GET")
	uziRouter.HandleFunc("/uzis/{id}", uziHandler.PatchUzi).Methods("PATCH")
	uziRouter.HandleFunc("/uzis", uziHandler.PostUzi).Methods("POST")

	slog.Info("start serve", slog.String("url", cfg.App.Url))
	if err := http.ListenAndServe(cfg.App.Url, r); err != nil {
		slog.Error("serve", "err", err)
		return failExitCode
	}

	return successExitCode
}
