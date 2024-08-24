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
func (a *AuthUseCase) Login(ctx context.Context, request *enity.RequestLogin) (*enity.TokensPair, error) {
	a.logger.Debug("Login usecase started", zap.Any("Request", request))

	a.logger.Info("[Request] Get user by login")
	user, err := a.authRepo.GetUserByLogin(ctx, request.Email)
	if err != nil {
		a.logger.Error("Get user by login", zap.Error(err))
		return nil, fmt.Errorf("get user by login: %w", err)
	}
	a.logger.Debug("[Response] Got user by login")

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

	refreshTokenWord := gofakeit.MinecraftVillagerJob()
	tokensPair, err := a.generateTokenPair(ctx, user.ID, refreshTokenWord)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}
	// a.logger.Debug("Tokens pair", zap.Any("Pair", tokensPair))

	a.logger.Info("[Request] Update JWT tokens in DB")
	if err = a.authRepo.UpdateRefreshTokenByID(ctx, user.ID, refreshTokenWord); err != nil {
		a.logger.Error("Update refresh token", zap.Error(err))
		return nil, fmt.Errorf("update refresh token: %w", err)
	}
	a.logger.Info("[Response] Updated JWT tokens in DB")

	return tokensPair, nil
}

func (a *AuthUseCase) Register(ctx context.Context, request *enity.RequestRegister) (uint64, error) {
	panic("unimplemented")
}

func (a *AuthUseCase) TokenRefresh(ctx context.Context, refreshToken string) (*enity.TokensPair, error) {
	a.logger.Debug("Token refresh usecases started", zap.Any("Request", refreshToken))

	userData, err := a.jwtService.ParseUserData(refreshToken)
	if err != nil {
		a.logger.Error("Parse token err", zap.Error(err))
		return nil, fmt.Errorf("parse token: %w", err)
	}

	a.logger.Info("[Request] Get user by ID")
	user, err := a.authRepo.GetUserByID(ctx, userData.UserID)
	if err != nil {
		a.logger.Error("Get user by id err", zap.Error(err))
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	a.logger.Debug("[Response] Got user by ID")

	if user.RefreshTokenWord != userData.RefreshTokenWord {
		a.logger.Warn("Tokens RTW divergents", zap.Int("userID", user.ID))
		return nil, enity.ErrExpiredSession
	}

	refreshTokenWord := gofakeit.MinecraftVillagerStation()
	tokensPair, err := a.generateTokenPair(ctx, user.ID, refreshTokenWord)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}

	// a.logger.Debug("Tokens pair", zap.Any("Pair", tokensPair))

	a.logger.Info("[Request] Update user refresh token in repo")
	if err := a.authRepo.UpdateRefreshTokenByID(ctx, user.ID, refreshTokenWord); err != nil {
		return nil, fmt.Errorf("update refresh token in repo: %w", err)
	}
	a.logger.Info("[Reponse] Updated user refresh token in repo")

	return tokensPair, nil
}

// внедрить контекст в jwt
func (a *AuthUseCase) generateTokenPair(_ context.Context, userID int, refreshTokenWord string) (*enity.TokensPair, error) {
	a.logger.Info("[Request] Generate new JWT pair tokens")
	tokenPair, err := a.jwtService.GeneratePair(
		map[string]any{
			"id":        userID,
			"likesFood": gofakeit.MinecraftFood(),
		},
		refreshTokenWord,
	)
	if err != nil {
		a.logger.Error("Generate tokens", zap.Error(err))
		return nil, err
	}
	a.logger.Info("[Response] Generated JWT Tokens")

	return tokenPair, nil
}
