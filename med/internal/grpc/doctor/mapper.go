package doctor

import (
	"yir/med/internal/domain"
	pb "yir/med/internal/generated/grpc/service"
)

func domainToPb(d *domain.Doctor) *pb.Doctor {
	return &pb.Doctor{
		Id:       d.Id.String(),
		Fullname: d.FullName,
		Org:      d.Org,
		Job:      d.Job,
		Desc:     d.Desc,
	}
}
