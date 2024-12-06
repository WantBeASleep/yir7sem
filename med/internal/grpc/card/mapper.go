package card

import (
	"med/internal/domain"
	pb "med/internal/generated/grpc/service"
)

func domainToPb(d *domain.Card) *pb.Card {
	return &pb.Card{
		DoctorId:  d.DoctorID.String(),
		PatientId: d.PatientID.String(),
		Diagnosis: d.Diagnosis,
	}
}
