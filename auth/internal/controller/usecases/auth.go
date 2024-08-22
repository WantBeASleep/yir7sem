package usecases

import (
	"context"
	"yir/auth/internal/enity"
)

type Auth interface {
	// подумать над неймингом переменных, выглядит как кринжатина
	Login(ctx context.Context, loginData *enity.RequestLogin) (*enity.TokenPair, error)
	Register(ctx context.Context, regData *enity.RequestRegister) (uint64, error)
}