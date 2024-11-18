package repository

import (
	"context"
	"yir/med/internal/entity"
)

type Card interface {
	ListCards(ctx context.Context, limit, offset int) ([]*entity.PatientCard, int, error)
	CreateCard(ctx context.Context, Card *entity.PatientCard) error
	CardByID(ctx context.Context, ID string) (*entity.PatientCard, error)
	UpdateCardInfo(ctx context.Context, Card *entity.PatientCard) error
	PatchCardInfo(ctx context.Context, Card *entity.PatientCard) error
	DeleteCardInfo(ctx context.Context, ID string) error
}

type MedWorker interface {
	GetMedicalWorkerByID(ctx context.Context, ID string) (*entity.MedicalWorker, error)
	AddMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) (string, error)
	UpdateMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) error
	PatchMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) error
	ListMedicalWorkers(ctx context.Context, limit, offset int) ([]*entity.MedicalWorker, int, error)
	DeleteMedicalWorker(ctx context.Context, ID string) error
	GetPatientsByMedWorker(ctx context.Context, medWorkerID string) ([]*entity.PatientCard, error)
}

type Patient interface {
	CreatePatient(ctx context.Context, patient *entity.PatientInformation) error
	UpdatePatient(ctx context.Context, patient *entity.PatientInformation) error
	GetListPatient(ctx context.Context) ([]*entity.Patient, error)
	GetPatientInfoByID(ctx context.Context, ID string) (*entity.Patient, error)
	// TODO: SHOTS, сюда запрос будет возвращать PatientInformation
}
