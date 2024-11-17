package models

import (
	pb "yir/gateway/rpc/uzi"
	"github.com/google/uuid"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func ContorToPB(in []Point) []*pb.Point {
	res := make([]*pb.Point, 0, len(in))
	for _, p := range in {
		res = append(res, &pb.Point{
			X: int64(p.X),
			Y: int64(p.Y),
		})
	}

	return res
}

type Id struct {
	Id string
}

func IdToPbId(in Id) *pb.Id {
	return &pb.Id{
		Id: in.Id,
	}
}