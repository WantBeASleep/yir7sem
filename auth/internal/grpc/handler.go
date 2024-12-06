package grpc

import (
	"auth/internal/generated/grpc/service"
	"auth/internal/grpc/login"
	"auth/internal/grpc/refresh"
	"auth/internal/grpc/register"
)

type Handler struct {
	login.LoginHandler
	refresh.RefreshHandler
	register.RegisterHandler

	service.UnsafeAuthSrvServer
}

func New(
	loginHandler login.LoginHandler,
	refreshHandler refresh.RefreshHandler,
	registerHandler register.RegisterHandler,
) *Handler {
	return &Handler{
		LoginHandler:    loginHandler,
		RefreshHandler:  refreshHandler,
		RegisterHandler: registerHandler,
	}
}
