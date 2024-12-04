package grpc

import (
	"yirv2/auth/internal/generated/grpc/service"
	"yirv2/auth/internal/grpc/login"
	"yirv2/auth/internal/grpc/refresh"
	"yirv2/auth/internal/grpc/register"
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
