package usecases

import (
	"context"
	"service/internal/entity"
)

type MedicalWorker interface {
	GetMedWorkers(ctx context.Context, limit, offset int) (*entity.MedicalWorkerList, error)
	GetMedWorkerByID(ctx context.Context, ID int) (*entity.MedicalWorker, error)
	UpdateMedWorker(ctx context.Context, ID int, updateData *entity.MedicalWorkerUpdateRequest) (*entity.MedicalWorker, error)
	AddMedWorker(ctx context.Context, createData *entity.AddMedicalWorkerRequest) (*entity.MedicalWorker, error)
	GetPatientsByMedWorker(ctx context.Context, medWorkerID uint64) (*entity.MedicalWorkerWithPatients, error)
}