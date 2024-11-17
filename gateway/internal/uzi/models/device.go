package models

import (
	pb "yir/gateway/rpc/uzi"
	
)


type Device struct {
	ID   int `json:"id"`
	Name string `json:"name"`
}

func PBGetDeviceListToDevices(in *pb.GetDeviceListResponse) []Device {
	resp := make([]Device, len(in.Devices))
	for _, d := range resp {
		resp = append(resp, Device{
			ID: d.ID,
			Name: d.Name,
		})
	}
	return resp
} 