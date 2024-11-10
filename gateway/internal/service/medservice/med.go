package medservice

import (
	"yir/gateway/internal/pb/medpb"
)

type MedService struct {
	CardClient    medpb.MedCardClient
	WorkerClient  medpb.MedWorkersClient
	PatientClient medpb.MedPatientClient
}
