package device

import (
	"yirv2/uzi/internal/domain"
	pb "yirv2/uzi/internal/generated/grpc/service"
)

func domainDeviceToPbDevice(d *domain.Device) *pb.Device {
	return &pb.Device{
		Id:   int64(d.Id),
		Name: d.Name,
	}
}
