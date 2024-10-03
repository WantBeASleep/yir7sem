package usecase

import (
	"context"
	"yir/med/internal/entity"
)

type Patient interface {
	AddPatient(ctx context.Context, PatientInfo *entity.PatientInformation) error
	UpdatePatient(ctx context.Context, PatientInfo *entity.PatientInformation) error
	GetPatientList(ctx context.Context) ([]*entity.Patient, error)
	GetPatientInfoByID(ctx context.Context, ID uint64) (*entity.PatientInformation, error)
}

/*
type WorkerUsecase interface {}
*/
