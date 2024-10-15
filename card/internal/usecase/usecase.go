package usecase

import (
	"context"
	"service/internal/entity"
)

type Card interface {
	GetCards(ctx context.Context, limit, offset int) (*entity.PatientCardList, error)
	PostCard(ctx context.Context, Card *entity.PatientCard) error
	GetCardByID(ctx context.Context, ID uint64) (*entity.PatientCard, error)
	PutCard(ctx context.Context, Card *entity.PatientCard) error
	DeleteCard(ctx context.Context, ID uint64) error
}
