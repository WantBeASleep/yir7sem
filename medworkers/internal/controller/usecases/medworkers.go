package usecases

import (
	"context"
	"yir/medworkers/internal/entity"
)

type MedicalWorker interface {
	GetMedWorkers(ctx context.Context, limit, offset int) (*entity.MedicalWorkerList, error)
	GetMedWorkerByID(ctx context.Context, ID int) (*entity.MedicalWorker, error)
	UpdateMedWorker(ctx context.Context, ID int, updateData *entity.MedicalWorkerUpdateRequest) (*entity.MedicalWorker, error)
}
