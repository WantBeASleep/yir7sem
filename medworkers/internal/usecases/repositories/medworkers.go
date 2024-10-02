package repositories

import (
	"context"
	"service/internal/entity" //структуры из entity только в use case
)

type MedWorker interface {
	GetMedicalWorkerByID(ctx context.Context, ID int) (*entity.MedicalWorker, error)
	AddMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) (int, error)
	UpdateMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) error
	PatchMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) error
	ListMedicalWorkers(ctx context.Context, limit, offset int) ([]*entity.MedicalWorker, int, error)
	DeleteMedicalWorker(ctx context.Context, ID int) error
}

type PatientService interface {
	GetPatientsByMedWorker(ctx context.Context, medWorkerId uint64) ([]*entity.PatientCardDTO, error)
}
