package node

import (
	"yir/uzi/internal/domain"
	pb "yir/uzi/internal/generated/grpc/service"
)

func DomainNodeToPb(d *domain.Node) *pb.Node {
	return &pb.Node{
		Id:        d.Id.String(),
		Ai:        d.Ai,
		Tirads_23: d.Tirads23,
		Tirads_4:  d.Tirads4,
		Tirads_5:  d.Tirads5,
	}
}
