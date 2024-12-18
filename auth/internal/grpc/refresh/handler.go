package refresh

import (
	"context"
	"fmt"

	pb "auth/internal/generated/grpc/service"
	"auth/internal/services/refresh"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RefreshHandler interface {
	Refresh(ctx context.Context, in *pb.RefreshIn) (*pb.RefreshOut, error)
}

type handler struct {
	refreshSrv refresh.Service
	validator  *protovalidate.Validator
}

func New(
	refreshSrv refresh.Service,
) (RefreshHandler, error) {
	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			&pb.RefreshIn{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("init validator: %v", err)
	}
	return &handler{
		refreshSrv: refreshSrv,
		validator:  validator,
	}, nil
}

func (h *handler) Refresh(ctx context.Context, in *pb.RefreshIn) (*pb.RefreshOut, error) {
	if err := h.validator.Validate(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("validation failed: %v", err))
	}

	access, refresh, err := h.refreshSrv.Refresh(ctx, in.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.RefreshOut{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
