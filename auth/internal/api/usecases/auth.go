package usecases

import (
	"context"
	"yir/auth/internal/entity"

	"github.com/google/uuid"
)

type Auth interface {
	Login(ctx context.Context, mail string, password string) (*entity.TokensPair, error)
	Register(ctx context.Context, req *entity.RequestRegister) (uuid.UUID, error)
	TokenRefresh(ctx context.Context, request string) (*entity.TokensPair, error)
}
