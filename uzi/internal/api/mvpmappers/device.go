package mvpmappers

import (
	pb "yir/uzi/api/grpcapi"
	"yir/uzi/internal/entity"
)

func DeviceToPBDevice(device entity.Device) *pb.Device {
	return &pb.Device{
		Id:   int64(device.Id),
		Name: device.Name,
	}
}

func DevicesToPBDevices(devices []entity.Device) []*pb.Device {
	PBDevices := make([]*pb.Device, 0, len(devices))
	for _, device := range devices {
		PBDevices = append(PBDevices, DeviceToPBDevice(device))
	}

	return PBDevices
}
