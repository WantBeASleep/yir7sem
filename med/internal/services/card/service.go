package card

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
	CreateCard(ctx context.Context, card domain.Card) error
	GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error)
	UpdateCard(ctx context.Context, doctorID, patientID uuid.UUID, update UpdateCard) (domain.Card, error)
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

func (s *service) CreateCard(ctx context.Context, card domain.Card) error {
	if err := s.dao.NewCardQuery(ctx).InsertCard(entity.Card{}.FromDomain(card)); err != nil {
		return fmt.Errorf("insert card: %w", err)
	}

	return nil
}

// TODO: вынести в CARD_PK
// TODO: перенести все на entity
func (s *service) GetCard(ctx context.Context, doctorID, patientID uuid.UUID) (domain.Card, error) {
	card, err := s.dao.NewCardQuery(ctx).GetCardByPK(doctorID, patientID)
	if err != nil {
		return domain.Card{}, fmt.Errorf("get card by id: %w", err)
	}

	return card.ToDomain(), nil
}

func (s *service) UpdateCard(ctx context.Context, doctorID, patientID uuid.UUID, update UpdateCard) (domain.Card, error) {
	cardQuery := s.dao.NewCardQuery(ctx)

	exists, err := cardQuery.CheckCardExists(doctorID, patientID)
	if err != nil {
		return domain.Card{}, fmt.Errorf("check card exists: %w", err)
	}
	if !exists {
		return domain.Card{}, errors.New("card doesn't exists")
	}

	card, err := s.GetCard(ctx, doctorID, patientID)
	if err != nil {
		return domain.Card{}, fmt.Errorf("get card: %w", err)
	}
	update.Update(&card)

	if _, err := cardQuery.UpdateCard(entity.Card{}.FromDomain(card)); err != nil {
		return domain.Card{}, fmt.Errorf("update card: %w", err)
	}

	return card, nil
}
