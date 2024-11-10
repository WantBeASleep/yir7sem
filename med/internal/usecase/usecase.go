package usecase

import (
	"context"
	"yir/med/internal/entity"
)

type Card interface {
	GetCards(ctx context.Context, limit, offset int) (*entity.PatientCardList, error)
	PostCard(ctx context.Context, Card *entity.PatientInformation) error
	GetCardByID(ctx context.Context, ID string) (*entity.PatientInformation, error)
	PutCard(ctx context.Context, Card *entity.PatientCard) error
	DeleteCard(ctx context.Context, ID string) error
}

type Patient interface {
	AddPatient(ctx context.Context, PatientInfo *entity.PatientInformation) error
	UpdatePatient(ctx context.Context, PatientInfo *entity.PatientInformation) error
	GetPatientList(ctx context.Context) ([]*entity.Patient, error)
	GetPatientInfoByID(ctx context.Context, ID string) (*entity.PatientInformation, error)
}

type MedicalWorker interface {
	GetMedWorkers(ctx context.Context, limit, offset int) (*entity.MedicalWorkerList, error)
	GetMedWorkerByID(ctx context.Context, ID string) (*entity.MedicalWorker, error)
	UpdateMedWorker(ctx context.Context, ID string, updateData *entity.MedicalWorkerUpdateRequest) (*entity.MedicalWorker, error)
	AddMedWorker(ctx context.Context, createData *entity.AddMedicalWorkerRequest) (*entity.MedicalWorker, error)
	GetPatientsByMedWorker(ctx context.Context, medWorkerID string) ([]*entity.PatientCard, uint64, error)
}
