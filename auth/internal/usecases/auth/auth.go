package usecases

import (
	"context"
	"fmt"
	"yir/auth/internal/enity"
	"yir/auth/internal/usecases/repositories"

	"go.uber.org/zap"
)

type AuthUseCase struct {
	authRepo repositories.Auth
	logger   *zap.Logger
}

func NewAuthUseCase(
	authRepo repositories.Auth,
	logger *zap.Logger,
) *AuthUseCase {
	return &AuthUseCase{
		authRepo: authRepo,
		logger:   logger,
	}
}

func (a *AuthUseCase) Login(ctx context.Context, request *enity.RequestLogin) (*enity.TokenPair, error) {
	a.logger.Debug("Login usecase started", zap.Any("Request", request))

	a.logger.Debug("[Request] Get user by login", zap.String("Request", request.Email))
	user, err := a.authRepo.GetUserByLogin(ctx, request.Email)
	if err != nil {
		return nil, fmt.Errorf("get user by login: %w", err)
	}
	a.logger.Debug("[Response] Get user by login", zap.Any("Response", user))

	passHash, err := enity.HashByScrypt(request.Password)
	if err != nil {
		return nil, fmt.Errorf("password hash: %v", err)
	}

}
