package doctor

import (
	"context"
	"fmt"

	"med/internal/domain"
	"med/internal/repository"
	"med/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	RegisterDoctor(ctx context.Context, doctor domain.Doctor) error
	GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error)
	UpdateDoctor(ctx context.Context, id uuid.UUID, update UpdateDoctor) (domain.Doctor, error)
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

func (s *service) RegisterDoctor(ctx context.Context, doctor domain.Doctor) error {
	if err := s.dao.NewDoctorQuery(ctx).InsertDoctor(entity.Doctor{}.FromDomain(doctor)); err != nil {
		return fmt.Errorf("insert doctor: %w", err)
	}

	return nil
}

func (s *service) GetDoctor(ctx context.Context, id uuid.UUID) (domain.Doctor, error) {
	doctor, err := s.dao.NewDoctorQuery(ctx).GetDoctorByPK(id)
	if err != nil {
		return domain.Doctor{}, fmt.Errorf("get doctor by pk: %w", err)
	}

	return doctor.ToDomain(), nil
}

func (s *service) UpdateDoctor(ctx context.Context, id uuid.UUID, update UpdateDoctor) (domain.Doctor, error) {
	doctor, err := s.GetDoctor(ctx, id)
	if err != nil {
		return domain.Doctor{}, fmt.Errorf("get doctor: %w", err)
	}
	update.Update(&doctor)

	if _, err := s.dao.NewDoctorQuery(ctx).UpdateDoctor(entity.Doctor{}.FromDomain(doctor)); err != nil {
		return domain.Doctor{}, fmt.Errorf("update doctor: %w", err)
	}

	return doctor, nil
}
