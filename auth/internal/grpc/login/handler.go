package login

import (
	"context"
	"fmt"

	pb "auth/internal/generated/grpc/service"
	"auth/internal/services/login"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginHandler interface {
	Login(ctx context.Context, in *pb.LoginIn) (*pb.LoginOut, error)
}

type handler struct {
	loginSrv  login.Service
	validator *protovalidate.Validator
}

func New(
	loginSrv login.Service,
) (LoginHandler, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.LoginIn{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("init validator: %v", err)
	}
	return &handler{
		loginSrv:  loginSrv,
		validator: validator,
	}, nil
}

func (h *handler) Login(ctx context.Context, in *pb.LoginIn) (*pb.LoginOut, error) {
	if err := h.validator.Validate(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
	}

	access, refresh, err := h.loginSrv.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.LoginOut{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
