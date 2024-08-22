package usecases

import (
	"context"
	"fmt"
	"yir/auth/internal/enity"
	"yir/auth/internal/usecases/core"
	"yir/auth/internal/usecases/repositories"

	"github.com/brianvoe/gofakeit/v7"
	"go.uber.org/zap"
)

type AuthUseCase struct {
	authRepo   repositories.Auth
	jwtService core.JWTService
	logger     *zap.Logger
}

func NewAuthUseCase(
	authRepo repositories.Auth,
	jwtService core.JWTService,
	logger *zap.Logger,
) *AuthUseCase {
	return &AuthUseCase{
		authRepo: authRepo,
		logger:   logger,
	}
}

// покурить что может это сразу в уровень INFO логов
// ДОБАВИТЬ КОНТЕКСТ
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
		return nil, fmt.Errorf("password hash: %w", err)
	}

	if user.PasswordHash != passHash {
		return nil, enity.ErrPassHashNotEqual
	}

	a.logger.Debug("[Request] Generate new JWT pair tokens")
	refreshTokenWord := gofakeit.MinecraftVillagerJob()
	tokenPair, err := a.jwtService.Generate(
		map[string]any{
			"id":        user.ID,
			"likesFood": gofakeit.MinecraftFood(),
		},
		refreshTokenWord,
	)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}
	a.logger.Debug("[Response] Get Tokens")

	if err = a.authRepo.UpdateRefreshTokenByID(ctx, user.ID, tokenPair.RefreshToken); err != nil {
		return nil, fmt.Errorf("update refresh token: %w", err)
	}

	return tokenPair, nil
}
