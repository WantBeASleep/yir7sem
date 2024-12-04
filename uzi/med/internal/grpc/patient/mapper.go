package patient

import (
	"yirv2/med/internal/domain"
	pb "yirv2/med/internal/generated/grpc/service"
	"yirv2/pkg/gtclib"
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
