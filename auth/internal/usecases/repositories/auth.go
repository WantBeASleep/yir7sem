package repositories

import (
	"context"
	"yir/auth/internal/entity"

	"github.com/google/uuid"
)

type Auth interface {
	GetUserByID(ctx context.Context, ID int) (*entity.UserCreditals, error)
	GetUserByUUID(ctx context.Context, UUID uuid.UUID) (*entity.UserCreditals, error)
	GetUserByLogin(ctx context.Context, login string) (*entity.UserCreditals, error)
	CreateUser(ctx context.Context, user *entity.UserCreditals) (int, error)
	UpdateRefreshTokenByID(ctx context.Context, ID int, refreshTokenWord string) error
}
