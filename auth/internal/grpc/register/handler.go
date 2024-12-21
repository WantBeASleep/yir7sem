package register

import (
	"context"
	"fmt"

	"auth/internal/domain"
	pb "auth/internal/generated/grpc/service"
	"auth/internal/services/register"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterHandler interface {
	Register(ctx context.Context, in *pb.RegisterIn) (*pb.RegisterOut, error)
}

type handler struct {
	registerSrv register.Service
	validator   *protovalidate.Validator
}

func New(
	registerSrv register.Service,
) (RegisterHandler, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.RegisterIn{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("init validator: %w", err)
	}
	return &handler{
		registerSrv: registerSrv,
		validator:   validator,
	}, nil
}

func (h *handler) Register(ctx context.Context, in *pb.RegisterIn) (*pb.RegisterOut, error) {
	if err := h.validator.Validate(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
	}

	id, err := h.registerSrv.Register(ctx, domain.User{
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.RegisterOut{Id: id.String()}, nil
}
