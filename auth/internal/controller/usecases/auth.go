package usecases

import (
	"context"
	"yir/auth/internal/enity"
)

type Auth interface {
	Login(ctx context.Context, request *enity.RequestLogin) (*enity.TokensPair, error)
	Register(ctx context.Context, request *enity.RequestRegister) (uint64, error)
	TokenRefresh(ctx context.Context, request string) (*enity.TokensPair, error)
}
