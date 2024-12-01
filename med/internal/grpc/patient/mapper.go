package patient

import (
	"yir/med/internal/domain"
	pb "yir/med/internal/generated/grpc/service"
	"yir/pkg/gtclib"
)

func domainToPb(d *domain.Patient) *pb.Patient {
	return &pb.Patient{
		Id:          d.Id.String(),
		Fullname:    d.FullName,
		Email:       d.Email,
		Policy:      d.Policy,
		Active:      d.Active,
		Malignancy:  d.Malignancy,
		LastUziDate: gtclib.Timestamp.TimePointerTo(d.LastUziDate),
	}
}
