package repositories

import (
	"context"
	"yir/auth/internal/enity"
)

type Auth interface {
	GetUserByID(ctx context.Context, ID int) (*enity.User, error)
	GetUserByLogin(ctx context.Context, login string) (*enity.User, error)
	CreateUser(ctx context.Context, user *enity.User) (int, error)
	UpdateRefreshTokenByID(ctx context.Context, ID int, newToken string) error
}
