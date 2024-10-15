package repository

import (
	"context"
	"service/internal/entity"
)

type Card interface {
	ListCards(ctx context.Context, limit, offset int) ([]*entity.PatientCard, int, error)
	CreateCard(ctx context.Context, card *entity.PatientCard) error
	CardByID(ctx context.Context, ID uint64) (*entity.PatientCard, error)
	UpdateCardInfo(ctx context.Context, Card *entity.PatientCard) error
	PatchCardInfo(ctx context.Context, Card *entity.PatientCard) error
	DeleteCardInfo(ctx context.Context, ID int) error
}
