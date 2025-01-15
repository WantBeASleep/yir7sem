package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/WantBeASleep/goooool/brokerlib"
	"github.com/WantBeASleep/goooool/grpclib"
	"github.com/WantBeASleep/goooool/loglib"

	_ "gateway/docs"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"

	"gateway/internal/config"
	"gateway/internal/repository"

	pkgconfig "github.com/WantBeASleep/goooool/config"

	adapters "gateway/internal/adapters"
	brokeradapters "gateway/internal/adapters/broker"
	authgrpcadapter "gateway/internal/adapters/grpc/auth"
	medgrpcadapter "gateway/internal/adapters/grpc/med"
	uzigrpcadapter "gateway/internal/adapters/grpc/uzi"

	authhandler "gateway/internal/api/auth"
	downloadhandler "gateway/internal/api/download"
	medhandler "gateway/internal/api/med"
	uzihandler "gateway/internal/api/uzi"

	"gateway/internal/middleware"

	"github.com/minio/minio-go/v7/pkg/credentials"
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
	cfg, err := pkgconfig.Load[config.Config]()
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

	authConn, err := grpc.NewClient(
		cfg.Adapters.AuthUrl,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(grpclib.ClientCallLogger),
	)
	if err != nil {
		slog.Error("init uziConn", "err", err)
		return failExitCode
	}

	medAdapter := medgrpcadapter.New(medConn)
	uziAdapter := uzigrpcadapter.New(uziConn)
	authAdapter := authgrpcadapter.New(authConn)

	adapter := adapters.New(
		authAdapter,
		medAdapter,
		uziAdapter,
		brokeradapter,
	)

	authHandler := authhandler.New(adapter)
	medHandler := medhandler.New(adapter)
	uziHandler := uzihandler.New(adapter, dao)
	downloadHandler := downloadhandler.New(dao)

	// TODO: пробросить ошибки с логированием на верхнем уровне
	mdlwrs := middleware.New(pubKey)

	r := mux.NewRouter()

	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	apiRouter := r.PathPrefix("/api").Subrouter()

	downloadRouter := apiRouter.PathPrefix("/download").Subrouter()
	downloadRouter.Use(mdlwrs.Log, mdlwrs.Jwt)

	downloadRouter.HandleFunc("/uzi/{id}", downloadHandler.GetUzi).Methods("GET")
	downloadRouter.HandleFunc("/uzi/{uzi_id}/image/{image_id}", downloadHandler.GetImage).Methods("GET")

	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.Use(mdlwrs.Log)

	authRouter.HandleFunc("/register", authHandler.Register).Methods("POST")
	authRouter.HandleFunc("/login", authHandler.Login).Methods("POST")
	authRouter.HandleFunc("/refresh", authHandler.Refresh).Methods("POST")

	medRouter := apiRouter.PathPrefix("/med").Subrouter()
	medRouter.Use(mdlwrs.Log, mdlwrs.Jwt)

	medRouter.HandleFunc("/card/{id}", medHandler.UpdateCard).Methods("PATCH")
	medRouter.HandleFunc("/card/{id}", medHandler.GetCard).Methods("GET")
	medRouter.HandleFunc("/card", medHandler.PostCard).Methods("POST")

	medRouter.HandleFunc("/patient/{id}/uzis", medHandler.GetDoctorPatients).Methods("GET")
	medRouter.HandleFunc("/patient/{id}", medHandler.UpdatePatient).Methods("PATCH")
	medRouter.HandleFunc("/patient/{id}", medHandler.GetPatient).Methods("GET")
	medRouter.HandleFunc("/patient", medHandler.PostPatient).Methods("POST")

	medRouter.HandleFunc("/doctors", medHandler.UpdateDoctor).Methods("PATCH")
	medRouter.HandleFunc("/doctors/patient", medHandler.GetDoctorPatients).Methods("GET")
	medRouter.HandleFunc("/doctors", medHandler.GetDoctor).Methods("GET")

	uziRouter := apiRouter.PathPrefix("/uzi").Subrouter()
	uziRouter.Use(mdlwrs.Log, mdlwrs.Jwt)

	uziRouter.HandleFunc("/echographics/{id}", uziHandler.PatchEchographics).Methods("PATCH")
	uziRouter.HandleFunc("/echographics/{id}", uziHandler.GetEchographics).Methods("GET")

	uziRouter.HandleFunc("/segments/{id}", uziHandler.PatchSegment).Methods("PATCH")
	uziRouter.HandleFunc("/segments/{id}", uziHandler.DeleteSegment).Methods("DELETE")
	uziRouter.HandleFunc("/segments", uziHandler.PostSegment).Methods("POST")

	uziRouter.HandleFunc("/nodes/{id}", uziHandler.PatchNode).Methods("PATCH")
	uziRouter.HandleFunc("/nodes/{id}", uziHandler.DeleteNode).Methods("DELETE")
	uziRouter.HandleFunc("/nodes", uziHandler.PostNodes).Methods("POST")

	uziRouter.HandleFunc("/images/{id}/nodes-segments", uziHandler.GetUziNodeSegments).Methods("GET")

	uziRouter.HandleFunc("/patient/{id}/uzis", uziHandler.GetPatientUzi).Methods("GET")
	uziRouter.HandleFunc("/uzis/{id}/images", uziHandler.GetUziImages).Methods("GET")
	uziRouter.HandleFunc("/uzis/{id}/nodes", uziHandler.GetAllNodes).Methods("GET")
	uziRouter.HandleFunc("/uzis/{id}", uziHandler.GetUzi).Methods("GET")
	uziRouter.HandleFunc("/uzis/{id}", uziHandler.PatchUzi).Methods("PATCH")
	uziRouter.HandleFunc("/uzis", uziHandler.PostUzi).Methods("POST")

	uziRouter.HandleFunc("/devices", uziHandler.GetUziDevices).Methods("GET")

	slog.Info("start serve", slog.String("url", cfg.App.Url))
	if err := http.ListenAndServe(cfg.App.Url, r); err != nil {
		slog.Error("serve", "err", err)
		return failExitCode
	}

	return successExitCode
}
