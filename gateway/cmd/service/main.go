package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "yirv2/gateway/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"

	"yirv2/gateway/internal/config"
	pkgconfig "yirv2/pkg/config"
	"yirv2/pkg/grpclib"
	"yirv2/pkg/loglib"

	"yirv2/gateway/internal/adapters"
	medadapter "yirv2/gateway/internal/adapters/med"

	medhandler "yirv2/gateway/internal/api/med"

	"yirv2/gateway/internal/middleware"
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

	medConn, err := grpc.NewClient(
		cfg.Adapters.MedUrl,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(grpclib.ClientCallLogger),
	)
	if err != nil {
		slog.Error("init authConn", "err", err)
		return failExitCode
	}

	medAdapter := medadapter.New(medConn)

	adapter := adapters.New(
		nil,
		medAdapter,
		nil,
	)

	medHandler := medhandler.New(adapter)

	jwtMdlwr := middleware.NewJWTMiddleware(pubKey)

	r := mux.NewRouter()

	r.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusMovedPermanently)
	})
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	apiRouter := r.PathPrefix("/api").Subrouter()

	medRouter := apiRouter.PathPrefix("/med").Subrouter()
	medRouter.Use(jwtMdlwr.JwtMiddleware)

	medRouter.HandleFunc("/doctors", medHandler.GetDoctor).Methods("GET")

	slog.Info("start serve", slog.String("url", cfg.App.Url))
	if err := http.ListenAndServe(cfg.App.Url, r); err != nil {
		slog.Error("serve", "err", err)
		return failExitCode
	}

	return successExitCode
}
