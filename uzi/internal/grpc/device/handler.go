package device

import (
	"context"

	"uzi/internal/services/device"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "uzi/internal/generated/grpc/service"
)

type DeviceHandler interface {
	GetDeviceList(ctx context.Context, _ *empty.Empty) (*pb.GetDeviceListOut, error)
}

type handler struct {
	deviceSrv device.Service
}

func New(
	deviceSrv device.Service,
) DeviceHandler {
	return &handler{
		deviceSrv: deviceSrv,
	}
}

func (h *handler) GetDeviceList(ctx context.Context, _ *empty.Empty) (*pb.GetDeviceListOut, error) {
	devices, err := h.deviceSrv.GetDeviceList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	out := new(pb.GetDeviceListOut)
	for _, d := range devices {
		pbDevice := domainDeviceToPbDevice(&d)
		out.Devices = append(out.Devices, pbDevice)
	}

	return out, nil
}
