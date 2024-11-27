package repository

import (
	"context"
	"service/all/internal/entity"
)

type Card interface {
	ListCards(ctx context.Context, limit, offset int) ([]*entity.PatientInformation, int, error)
	CreateCard(ctx context.Context, Card *entity.PatientInformation) error
	CardByID(ctx context.Context, ID uint64) (*entity.PatientInformation, error)
	UpdateCardInfo(ctx context.Context, Card *entity.PatientCard) error
	PatchCardInfo(ctx context.Context, Card *entity.PatientCard) error
	DeleteCardInfo(ctx context.Context, ID int) error
}

type MedWorker interface {
	GetMedicalWorkerByID(ctx context.Context, ID int) (*entity.MedicalWorker, error)
	AddMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) (int, error)
	UpdateMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) error
	PatchMedicalWorker(ctx context.Context, medworker *entity.MedicalWorker) error
	ListMedicalWorkers(ctx context.Context, limit, offset int) ([]*entity.MedicalWorker, int, error)
	DeleteMedicalWorker(ctx context.Context, ID int) error
	GetPatientsByMedWorker(ctx context.Context, medWorkerID uint64) ([]*entity.PatientCard, error)
}

type Patient interface {
	CreatePatient(ctx context.Context, patient *entity.PatientInformation) error
	UpdatePatient(ctx context.Context, patient *entity.PatientInformation) error
	GetListPatient(ctx context.Context) ([]*entity.Patient, error)
	GetPatientInfoByID(ctx context.Context, ID int) (*entity.Patient, error)
	// TODO: SHOTS, сюда запрос будет возвращать PatientInformation
}
