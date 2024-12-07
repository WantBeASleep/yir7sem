package patient

import (
	"github.com/WantBeASleep/goooool/gtclib"

	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
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
