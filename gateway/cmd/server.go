package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	auth "yir/gateway/internal/auth"
	authpb "yir/gateway/rpc/auth"

	"flag"
	"yir/gateway/internal/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

//	@title		API-Gateway
//	@version	1.0
//	@schemes	http
func main() {
	flag.Parse()
	cfg := config.MustLoad(configPath)

	s := grpc.NewServer()
	mux := runtime.NewServeMux()

	authConn, err := grpc.Dial(cfg.Auth.Url, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("auth conn: %w", err))
	}

	authCtrl := auth.NewCtrl(authpb.NewAuthClient(authConn))

	authpb.RegisterAuthServer(s, authCtrl)
	if err := authpb.RegisterAuthHandlerFromEndpoint(context.Background(), mux, cfg.Gateway.Host+":"+cfg.Gateway.GRPCport, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
		panic(fmt.Errorf("register http auth: %w", err))
	}

	GRPCLis, err := net.Listen("tcp", cfg.Gateway.Host+":"+cfg.Gateway.GRPCport)
	if err != nil {
		panic(fmt.Errorf("lister grpc host:port: %w", err))
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if err := s.Serve(GRPCLis); err != nil {
			panic(fmt.Errorf("start serve grpc: %w", err))
		}
		wg.Done()
	}()

	go func() {
		if err := http.ListenAndServe(cfg.Gateway.Host+":"+cfg.Gateway.HTTPport, mux); err != nil {
			panic(fmt.Errorf("start serve http: %w", err))
		}
		wg.Done()
	}()

	wg.Wait()
}
