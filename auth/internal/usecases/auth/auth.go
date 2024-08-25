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
	medRepo    repositories.MedRepo
	jwtService core.JWTService

	logger *zap.Logger
}

func NewAuthUseCase(
	authRepo repositories.Auth,
	medRepo repositories.MedRepo,
	jwtService core.JWTService,
	logger *zap.Logger,
) *AuthUseCase {
	return &AuthUseCase{
		authRepo:   authRepo,
		medRepo:    medRepo,
		jwtService: jwtService,
		logger:     logger,
	}
}

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
	tokensPair, err := a.generateTokenPair(user.ID, refreshTokenWord)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}

	a.logger.Info("[Request] Update JWT tokens in DB")
	if err = a.authRepo.UpdateRefreshTokenByID(ctx, user.ID, refreshTokenWord); err != nil {
		a.logger.Error("Update refresh token", zap.Error(err))
		return nil, fmt.Errorf("update refresh token: %w", err)
	}
	a.logger.Info("[Response] Updated JWT tokens in DB")

	return tokensPair, nil
}

func (a *AuthUseCase) Register(ctx context.Context, request *enity.RequestRegister) (uint64, error) {
	a.logger.Debug("Register usecases started", zap.Any("Request", request))

	// заглушка https://popovza.kaiten.ru/space/420777/card/37360398
	err := a.medRepo.AddMed()
	if err != nil {
		a.logger.Warn("Med Repo not implemented")
	}
	medWorkerID := gofakeit.Number(0, 1<<10)

	salt := gofakeit.MinecraftBiome()
	passHash, err := enity.HashByScrypt(request.Password, salt)
	if err != nil {
		a.logger.Error("Password hashing", zap.Error(err))
		return 0, fmt.Errorf("password hashing: %w", err)
	}

	a.logger.Info("[Request] Add new user")
	user := enity.User{
		Login:        request.Email,
		PasswordHash: passHash + salt,
		MedWorkerID:  medWorkerID,
	}
	userID, err := a.authRepo.CreateUser(ctx, &user)
	if err != nil {
		a.logger.Error("Create new user")
		return 0, fmt.Errorf("create new user: %w", err)
	}
	a.logger.Info("[Response] Added new user")

	refreshTokenWord := gofakeit.MinecraftVillagerJob()
	_, err = a.generateTokenPair(userID, refreshTokenWord)
	if err != nil {
		return 0, fmt.Errorf("generate tokens: %w", err)
	}

	a.logger.Info("[Request] Update refresh token")
	if err := a.authRepo.UpdateRefreshTokenByID(ctx, userID, refreshTokenWord); err != nil {
		a.logger.Error("Update refresh token in DB")
		return 0, fmt.Errorf("update refresh token: %w", err)
	}
	a.logger.Error("[Response] Updated refresh token")

	return uint64(userID), nil
}

func (a *AuthUseCase) TokenRefresh(ctx context.Context, request string) (*enity.TokensPair, error) {
	a.logger.Debug("Token refresh usecases started", zap.Any("Request", request))

	userData, err := a.jwtService.ParseUserData(request)
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
	tokensPair, err := a.generateTokenPair(user.ID, refreshTokenWord)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}

	a.logger.Info("[Request] Update user refresh token in repo")
	if err := a.authRepo.UpdateRefreshTokenByID(ctx, user.ID, refreshTokenWord); err != nil {
		return nil, fmt.Errorf("update refresh token in repo: %w", err)
	}
	a.logger.Info("[Reponse] Updated user refresh token in repo")

	return tokensPair, nil
}
