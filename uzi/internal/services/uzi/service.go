package uzi

import (
	"context"
	"errors"
	"fmt"
	"time"

	"yirv2/uzi/internal/domain"
	"yirv2/uzi/internal/repository"
	"yirv2/uzi/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	CreateUzi(ctx context.Context, uzi domain.Uzi) (uuid.UUID, error)
	GetUziByID(ctx context.Context, id uuid.UUID) (domain.Uzi, domain.Echographic, error)
	UpdateUzi(ctx context.Context, id uuid.UUID, update UpdateUzi) (domain.Uzi, error)
	UpdateEchographic(ctx context.Context, id uuid.UUID, update UpdateEchographic) (domain.Echographic, error)
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

func (s *service) CreateUzi(ctx context.Context, uzi domain.Uzi) (uuid.UUID, error) {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("start transaction: %w", err)
	}

	uzi.Id = uuid.New()
	uzi.Checked = false
	uzi.CreateAt = time.Now()

	if err := s.dao.NewUziQuery(ctx).InsertUzi(entity.Uzi{}.FromDomain(uzi)); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return uuid.Nil, fmt.Errorf("insert uzi: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.NewEchographicQuery(ctx).InsertEchographic(entity.Echographic{Id: uzi.Id}); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return uuid.Nil, fmt.Errorf("insert echographic: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("commit transaction: %w", err)
	}

	return uzi.Id, nil
}

func (s *service) GetUziByID(ctx context.Context, id uuid.UUID) (domain.Uzi, domain.Echographic, error) {
	uzi, err := s.dao.NewUziQuery(ctx).GetUziByPK(id)
	if err != nil {
		return domain.Uzi{}, domain.Echographic{}, fmt.Errorf("get uzi: %w", err)
	}

	echographic, err := s.dao.NewEchographicQuery(ctx).GetEchographicByPK(id)
	if err != nil {
		return domain.Uzi{}, domain.Echographic{}, fmt.Errorf("get echographic: %w", err)
	}

	return uzi.ToDomain(), echographic.ToDomain(), nil
}

func (s *service) UpdateUzi(ctx context.Context, id uuid.UUID, update UpdateUzi) (domain.Uzi, error) {
	uzi, _, err := s.GetUziByID(ctx, id)
	if err != nil {
		return domain.Uzi{}, fmt.Errorf("get uzi: %w", err)
	}
	update.Update(&uzi)

	if _, err := s.dao.NewUziQuery(ctx).UpdateUzi(entity.Uzi{}.FromDomain(uzi)); err != nil {
		return domain.Uzi{}, fmt.Errorf("update uzi: %w", err)
	}

	return uzi, nil
}

func (s *service) UpdateEchographic(ctx context.Context, id uuid.UUID, update UpdateEchographic) (domain.Echographic, error) {
	_, echographic, err := s.GetUziByID(ctx, id)
	if err != nil {
		return domain.Echographic{}, fmt.Errorf("get uzi: %w", err)
	}
	update.Update(&echographic)

	if _, err := s.dao.NewEchographicQuery(ctx).UpdateEchographic(entity.Echographic{}.FromDomain(echographic)); err != nil {
		return domain.Echographic{}, fmt.Errorf("update echographic: %w", err)
	}

	return echographic, nil
}
