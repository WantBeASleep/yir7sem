package services

import (
	"context"
	"service/internal/entity"
)

type PatientService interface {
	GetPatientsByMedWorker(ctx context.Context, medWorkerID uint64) ([]*entity.PatientCardDTO, error)
}
