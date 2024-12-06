package grpc

import (
	"yir/auth/internal/generated/grpc/service"
	"yir/auth/internal/grpc/login"
	"yir/auth/internal/grpc/refresh"
	"yir/auth/internal/grpc/register"
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
