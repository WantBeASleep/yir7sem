package device

import (
	"yir/uzi/internal/domain"
	pb "yir/uzi/internal/generated/grpc/service"
)

func domainDeviceToPbDevice(d *domain.Device) *pb.Device {
	return &pb.Device{
		Id:   int64(d.Id),
		Name: d.Name,
	}
}
