package app

import (
	"context"
	"net/http"
	"yir/gateway/internal/config"

	"github.com/gorilla/mux"
)

type Gateway struct {
	httpServer *http.Server
}

func New(cfg config.Gateway, r *mux.Router) *Gateway {
	return &Gateway{
		httpServer: &http.Server{
			Addr:         cfg.Host + ":" + cfg.Port,
			Handler:      r,
			ReadTimeout:  cfg.Timeout,
			WriteTimeout: cfg.Timeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}

func (g *Gateway) Run() error {
	return g.httpServer.ListenAndServe() // запускаем сервер, в случае неудачи вернет ошибку
}

func (g *Gateway) Shutdown(ctx context.Context) error {
	return g.httpServer.Shutdown(ctx)
}
