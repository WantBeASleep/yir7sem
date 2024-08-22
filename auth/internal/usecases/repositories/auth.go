package repositories

import (
	"context"
	"yir/auth/internal/enity"
)

type Auth interface {
	GetUserByID(ctx context.Context, ID uint) (*enity.DomainUser, error)
	GetUserByLogin(ctx context.Context, login string) (*enity.DomainUser, error)
	CreateUser(ctx context.Context, user enity.DomainUser) (uint, error)
	UpdateRefreshTokenByID(ctx context.Context, ID uint, newToken string) error
}