package login

import (
	"context"

	pb "auth/internal/generated/grpc/service"
	"auth/internal/services/login"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginHandler interface {
	Login(ctx context.Context, in *pb.LoginIn) (*pb.LoginOut, error)
}

type handler struct {
	loginSrv login.Service
}

func New(
	loginSrv login.Service,
) LoginHandler {
	return &handler{
		loginSrv: loginSrv,
	}
}

func (h *handler) Login(ctx context.Context, in *pb.LoginIn) (*pb.LoginOut, error) {
	access, refresh, err := h.loginSrv.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.LoginOut{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
