package app

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type Gateway struct {
	httpServer *http.Server
}

func New(addr, port string, r *mux.Router) *Gateway {
	return &Gateway{
		httpServer: &http.Server{
			Addr:    addr + ":" + port,
			Handler: r,
			// доп. параметры сервера
		},
	}
}

func (g *Gateway) Run() error {
	return g.httpServer.ListenAndServe() // запускаем сервер, в случае неудачи вернет ошибку
}

func (g *Gateway) Shutdown(ctx context.Context) error {
	return g.httpServer.Shutdown(ctx)
}
