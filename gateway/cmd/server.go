package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	_ "yir/gateway/docs"
	"yir/pkg/kafka"

	auth "yir/gateway/internal/auth"
	"yir/gateway/internal/med"
	authpb "yir/gateway/rpc/auth"

	s3api "yir/s3upload/api"
	s3client "yir/s3upload/pkg/client"

	uzi "yir/gateway/internal/uzi"
	uzipb "yir/gateway/rpc/uzi"

	medpb "yir/gateway/rpc/med"

	"flag"
	"yir/gateway/internal/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpswag "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultCfgPath = "/home/wantbeasleep/yir/gateway/config/config.yml"
	shorthand      = " (shorthand)"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", defaultCfgPath, "set config path")
	flag.StringVar(&configPath, "c", defaultCfgPath, "set config path"+shorthand)
}

// @title		1000-7 dev cursed
// @version	1.0
// @schemes	http
func main() {
	flag.Parse()
	cfg := config.MustLoad(configPath)

	s := grpc.NewServer()
	mux := runtime.NewServeMux()

	medConn, err := grpc.Dial(cfg.Med.Url, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("med conn: %w", err))
	}
	log.Println("Connected to MED")

	medCtrl := med.NewCtrl(medpb.NewMedCardClient(medConn), medpb.NewMedPatientClient(medConn), medpb.NewMedWorkersClient(medConn))

	medpb.RegisterMedCardServer(s, medCtrl)    // Пример регистрации MedCardService
	medpb.RegisterMedPatientServer(s, medCtrl) // Пример регистрации MedPatientService
	medpb.RegisterMedWorkersServer(s, medCtrl) // Пример регистрации MedWorkersService
	if err := medpb.RegisterMedCardHandlerFromEndpoint(context.Background(), mux, cfg.Gateway.Host+":"+cfg.Gateway.GRPCport, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		panic(fmt.Errorf("register http med card: %w", err))
	}
	if err := medpb.RegisterMedPatientHandlerFromEndpoint(context.Background(), mux, cfg.Gateway.Host+":"+cfg.Gateway.GRPCport, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		panic(fmt.Errorf("register http med card: %w", err))
	}
	if err := medpb.RegisterMedWorkersHandlerFromEndpoint(context.Background(), mux, cfg.Gateway.Host+":"+cfg.Gateway.GRPCport, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		panic(fmt.Errorf("register http med card: %w", err))
	}
	log.Println("Registered MED domens and set endpoints")

	authConn, err := grpc.Dial(cfg.Auth.Url, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("auth conn: %w", err))
	}
	log.Println("Connected to Auth")

	authCtrl := auth.NewCtrl(authpb.NewAuthClient(authConn))

	authpb.RegisterAuthServer(s, authCtrl)
	if err := authpb.RegisterAuthHandlerFromEndpoint(context.Background(), mux, cfg.Gateway.Host+":"+cfg.Gateway.GRPCport, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		panic(fmt.Errorf("register http auth: %w", err))
	}
	log.Println("Registered Auth and set endpoint")

	s3conn, err := grpc.Dial(cfg.S3.Url, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("uzi conn: %w", err))
	}
	s3clientGRPC := s3api.NewS3Client(s3conn)
	s3clientPKG := s3client.NewS3Client(s3clientGRPC)

	prod := kafka.New([]string{"localhost:9092"}, "uziUpload")

	uziConn, err := grpc.Dial(cfg.Uzi.Url, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("uzi conn: %w", err))
	}
	log.Println("Connected to Uzi")
	uziCtrl := uzi.NewCtrl(uzipb.NewUziAPIClient(uziConn), s3clientPKG, prod)

	uzipb.RegisterUziAPIServer(s, uziCtrl)
	if err := uzipb.RegisterUziAPIHandlerFromEndpoint(context.Background(), mux, cfg.Gateway.Host+":"+cfg.Gateway.GRPCport, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		panic(fmt.Errorf("register http auth: %w", err))
	}
	log.Println("Registered Uzi and set endpoint")

	GRPCLis, err := net.Listen("tcp", cfg.Gateway.Host+":"+cfg.Gateway.GRPCport)
	if err != nil {
		panic(fmt.Errorf("lister grpc host:port: %w", err))
	}

	mux.HandlePath("GET", "/docs/*", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		httpswag.WrapHandler.ServeHTTP(w, r)
	})

	mux.HandlePath("POST", "/uzi/create", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		uziCtrl.CreateUziHTTP(w, r)
	})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if err := s.Serve(GRPCLis); err != nil {
			panic(fmt.Errorf("start serve grpc: %w", err))
		}
		log.Println("gRPC started listen")
		wg.Done()
	}()

	go func() {
		if err := http.ListenAndServe(cfg.Gateway.Host+":"+cfg.Gateway.HTTPport, mux); err != nil {
			panic(fmt.Errorf("start serve http: %w", err))
		}
		log.Println("HTTP started listen")
		wg.Done()
	}()

	wg.Wait()
}
