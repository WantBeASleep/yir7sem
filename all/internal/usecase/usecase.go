package usecase

import (
	"context"
	"service/all/internal/entity"
)

type Card interface {
	GetCards(ctx context.Context, limit, offset int) (*entity.PatientCardList, error)
	PostCard(ctx context.Context, Card *entity.PatientInformation) error
	GetCardByID(ctx context.Context, ID uint64) (*entity.PatientInformation, error)
	PutCard(ctx context.Context, Card *entity.PatientCard) error
	DeleteCard(ctx context.Context, ID uint64) error
}

type Patient interface {
	AddPatient(ctx context.Context, PatientInfo *entity.PatientInformation) error
	UpdatePatient(ctx context.Context, PatientInfo *entity.PatientInformation) error
	GetPatientList(ctx context.Context) ([]*entity.Patient, error)
	GetPatientInfoByID(ctx context.Context, ID uint64) (*entity.PatientInformation, error)
}

type MedicalWorker interface {
	GetMedWorkers(ctx context.Context, limit, offset int) (*entity.MedicalWorkerList, error)
	GetMedWorkerByID(ctx context.Context, ID int) (*entity.MedicalWorker, error)
	UpdateMedWorker(ctx context.Context, ID int, updateData *entity.MedicalWorkerUpdateRequest) (*entity.MedicalWorker, error)
	AddMedWorker(ctx context.Context, createData *entity.AddMedicalWorkerRequest) (*entity.MedicalWorker, error)
	GetPatientsByMedWorker(ctx context.Context, medWorkerID uint64) ([]*entity.PatientCard, uint64, error)
}
