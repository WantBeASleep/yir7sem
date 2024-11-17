package models

import (
pb "yir/gateway/rpc/uzi"
)

type Tirads struct {
	Tirads23 float64 `json:"tirads_23"`
	Tirads4  float64 `json:"tirads_4"`
	Tirads5  float64 `json:"tirads_5"`
}

func PBTiradsToTirads(in *pb.Tirads) Tirads {
	return Tirads{
		Tirads23: in.Tirads_23,
		Tirads4: in.Tirads_4,
		Tirads5: in.Tirads_5,
	}
}

func TiradsToPB(in *Tirads) *pb.Tirads {
	return &pb.Tirads{
		Tirads_23: in.Tirads23,
		Tirads_4: in.Tirads4,
		Tirads_5: in.Tirads5,
	}
}