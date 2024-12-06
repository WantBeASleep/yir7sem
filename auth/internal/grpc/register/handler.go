package register

import (
	"context"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"
	"auth/internal/services/register"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterHandler interface {
	Register(ctx context.Context, in *pb.RegisterIn) (*pb.RegisterOut, error)
}

type handler struct {
	registerSrv register.Service
}

func New(
	registerSrv register.Service,
) RegisterHandler {
	return &handler{
		registerSrv: registerSrv,
	}
}

func (h *handler) Register(ctx context.Context, in *pb.RegisterIn) (*pb.RegisterOut, error) {
	id, err := h.registerSrv.Register(ctx, domain.User{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.RegisterOut{Id: id.String()}, nil
}
