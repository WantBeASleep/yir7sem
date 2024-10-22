package usecases

import (
	"context"
	"service/auth/internal/entity"
)

type Auth interface {
	Login(ctx context.Context, request *entity.RequestLogin) (*entity.TokensPair, error)
	Register(ctx context.Context, request *entity.RequestRegister) (uint64, error)
	TokenRefresh(ctx context.Context, request string) (*entity.TokensPair, error)
}
