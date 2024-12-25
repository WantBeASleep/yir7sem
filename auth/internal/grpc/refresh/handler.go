package refresh

import (
	"context"

	pb "auth/internal/generated/grpc/service"
	"auth/internal/services/refresh"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RefreshHandler interface {
	Refresh(ctx context.Context, in *pb.RefreshIn) (*pb.RefreshOut, error)
}

type handler struct {
	refreshSrv refresh.Service
}

func New(
	refreshSrv refresh.Service,
) RefreshHandler {
	return &handler{
		refreshSrv: refreshSrv,
	}
}

func (h *handler) Refresh(ctx context.Context, in *pb.RefreshIn) (*pb.RefreshOut, error) {
	access, refresh, err := h.refreshSrv.Refresh(ctx, in.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.RefreshOut{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
