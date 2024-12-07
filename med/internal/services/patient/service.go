package patient

import (
	"context"
	"errors"
	"fmt"

	"med/internal/domain"
	"med/internal/repository"
	"med/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	CreatePatient(ctx context.Context, patient domain.Patient) (uuid.UUID, error)
	GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error)
	GetDoctorPatients(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error)
	UpdatePatient(ctx context.Context, doctorID uuid.UUID, patientID uuid.UUID, update UpdatePatient) (domain.Patient, error)
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}

// TODO: нужно ввожить слой entity
func (s *service) CreatePatient(ctx context.Context, patient domain.Patient) (uuid.UUID, error) {
	patient.Id = uuid.New()
	if err := s.dao.NewPatientQuery(ctx).InsertPatient(entity.Patient{}.FromDomain(patient)); err != nil {
		return uuid.Nil, fmt.Errorf("insert patient: %w", err)
	}

	return patient.Id, nil
}

func (s *service) GetPatient(ctx context.Context, id uuid.UUID) (domain.Patient, error) {
	patient, err := s.dao.NewPatientQuery(ctx).GetPatientByPK(id)
	if err != nil {
		return domain.Patient{}, fmt.Errorf("get patient: %w", err)
	}

	return patient.ToDomain(), err
}

func (s *service) GetDoctorPatients(ctx context.Context, doctorID uuid.UUID) ([]domain.Patient, error) {
	patients, err := s.dao.NewPatientQuery(ctx).GetPatientsByDoctorID(doctorID)
	if err != nil {
		return nil, fmt.Errorf("get doctor patients: %w", err)
	}

	res := make([]domain.Patient, 0, len(patients))
	for _, v := range patients {
		res = append(res, v.ToDomain())
	}
	return res, nil
}

func (s *service) UpdatePatient(
	ctx context.Context,
	doctorID uuid.UUID,
	patientID uuid.UUID,
	update UpdatePatient,
) (domain.Patient, error) {
	exists, err := s.dao.NewCardQuery(ctx).CheckCardExists(doctorID, patientID)
	if err != nil {
		return domain.Patient{}, fmt.Errorf("check existing card")
	}
	if !exists {
		return domain.Patient{}, errors.New("card doctor-patient doesnt exists")
	}

	patient, err := s.GetPatient(ctx, patientID)
	if err != nil {
		return domain.Patient{}, fmt.Errorf("get patient: %w", err)
	}
	update.Update(&patient)

	if _, err := s.dao.NewPatientQuery(ctx).UpdatePatient(entity.Patient{}.FromDomain(patient)); err != nil {
		return domain.Patient{}, fmt.Errorf("update patient: %w", err)
	}

	return patient, nil
}
