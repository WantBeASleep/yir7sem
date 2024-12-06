package segment

import (
	"yir/uzi/internal/domain"
	pb "yir/uzi/internal/generated/grpc/service"
)

func DomainSegmentToPb(d *domain.Segment) *pb.Segment {
	return &pb.Segment{
		Id:        d.Id.String(),
		ImageId:   d.ImageID.String(),
		NodeId:    d.NodeID.String(),
		Contor:    d.Contor,
		Tirads_23: d.Tirads23,
		Tirads_4:  d.Tirads4,
		Tirads_5:  d.Tirads5,
	}
}
