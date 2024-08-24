package usecases

import (
	"context"
	"yir/auth/internal/enity"
)

type Auth interface {
	// подумать над неймингом переменных, выглядит как кринжатина
	Login(ctx context.Context, request *enity.RequestLogin) (*enity.TokensPair, error)
	Register(ctx context.Context, request *enity.RequestRegister) (uint64, error)
	TokenRefresh(ctx context.Context, refreshToken string) (*enity.TokensPair, error)
}
