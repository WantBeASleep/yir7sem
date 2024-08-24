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

	logger *zap.Logger
}

func NewAuthUseCase(
	authRepo repositories.Auth,
	jwtService core.JWTService,
	logger *zap.Logger,
) *AuthUseCase {
	return &AuthUseCase{
		authRepo:   authRepo,
		jwtService: jwtService,
		logger:     logger,
	}
}

// покурить что может это сразу в уровень INFO логов
// ДОБАВИТЬ КОНТЕКСТ
// Дебаги здесь потом менять на инфу вырезать приватное из логов
func (a *AuthUseCase) Login(ctx context.Context, request *enity.RequestLogin) (*enity.TokenPair, error) {
	a.logger.Debug("Login usecase started", zap.Any("Request", request))

	a.logger.Debug("[Request] Get user by login", zap.String("Request", request.Email))
	user, err := a.authRepo.GetUserByLogin(ctx, request.Email)
	if err != nil {
		a.logger.Error("Get user by login", zap.Error(err))
		return nil, fmt.Errorf("get user by login: %w", err)
	}
	a.logger.Debug("[Response] Get user by login", zap.Any("Response", user))

	salt := user.PasswordHash[64:]
	passHash, err := enity.HashByScrypt(request.Password, salt)
	if err != nil {
		a.logger.Error("Password hashing", zap.Error(err))
		return nil, fmt.Errorf("password hashing: %w", err)
	}
	a.logger.Debug("pass params", zap.String("salt", salt), zap.String("pass hash", passHash))

	if user.PasswordHash != passHash+salt {
		return nil, enity.ErrPassHashNotEqual
	}

	a.logger.Info("[Request] Generate new JWT pair tokens")
	refreshTokenWord := gofakeit.MinecraftVillagerJob()
	tokenPair, err := a.jwtService.Generate(
		map[string]any{
			"id":        user.ID,
			"likesFood": gofakeit.MinecraftFood(),
		},
		refreshTokenWord,
	)
	if err != nil {
		a.logger.Error("Generate tokens", zap.Error(err))
		return nil, fmt.Errorf("generate tokens: %w", err)
	}
	a.logger.Info("[Response] Get Tokens")

	a.logger.Info("[Request] Update JWT tokens in DB")
	if err = a.authRepo.UpdateRefreshTokenByID(ctx, user.ID, tokenPair.RefreshToken); err != nil {
		a.logger.Error("Update refresh token", zap.Error(err))
		return nil, fmt.Errorf("update refresh token: %w", err)
	}
	a.logger.Info("[Response] Update JWT tokens in DB")

	return tokenPair, nil
}

// Register(ctx context.Context, request *enity.RequestRegister) (uint64, error)
// TokenRefresh(ctx context.Context, refreshToken string) (*enity.TokenPair, error)
func (a *AuthUseCase) Register(ctx context.Context, request *enity.RequestRegister) (uint64, error) {
	panic("unimplemented")
}

func (a *AuthUseCase) TokenRefresh(ctx context.Context, refreshToken string) (*enity.TokenPair, error) {
	panic("unimplemented")
}
