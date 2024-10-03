package repositories

import (
	"context"
	"yir/auth/internal/entity"
)

type Auth interface {
	GetUserByID(ctx context.Context, ID int) (*entity.User, error)
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (int, error)
	UpdateRefreshTokenByID(ctx context.Context, ID int, refreshTokenWord string) error
}
