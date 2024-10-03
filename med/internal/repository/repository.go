package repository

import (
	"context"
	"yir/med/internal/entity"
)

type Patient interface {
	CreatePatient(ctx context.Context, patient *entity.PatientInformation) error
	UpdatePatient(ctx context.Context, patient *entity.PatientInformation) error
	GetListPatient(ctx context.Context) ([]*entity.Patient, error)
	GetPatientInfoByID(ctx context.Context, ID int) (*entity.Patient, error)
	// TODO: SHOTS, сюда запрос будет возвращать PatientInformation
}
