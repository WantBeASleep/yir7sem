package repositories

import (
	"context"
	"yir/auth/internal/entity"

	"github.com/google/uuid"
)

type Auth interface {
	GetUserByMail(ctx context.Context, mail string) (*entity.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)

	CreateUser(ctx context.Context, user *entity.User) (uuid.UUID, error)

	UpdateRefreshTokenByUserID(ctx context.Context, id uuid.UUID, refreshTokenWord string) error
}
