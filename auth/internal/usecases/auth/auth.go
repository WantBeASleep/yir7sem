package usecases

import (
	"context"
	"fmt"
	"yir/auth/internal/entity"
	"yir/auth/internal/usecases/core"
	"yir/auth/internal/usecases/repositories"

	"github.com/google/uuid"

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

func (a *AuthUseCase) Login(ctx context.Context, request *entity.RequestLogin) (*entity.TokensPair, error) {
	a.logger.Debug("Login usecase started", zap.Any("Request", request))

	a.logger.Info("[Request] Get user by login")
	user, err := a.authRepo.GetUserByLogin(ctx, request.Email)
	if err != nil {
		a.logger.Error("Get user by login", zap.Error(err))
		return nil, fmt.Errorf("get user by login: %w", err)
	}
	a.logger.Debug("[Response] Got user by login")

	salt := user.PasswordHash[64:]
	passHash, err := entity.HashByScrypt(request.Password, salt)
	if err != nil {
		a.logger.Error("Password hashing", zap.Error(err))
		return nil, fmt.Errorf("password hashing: %w", err)
	}
	a.logger.Debug("pass params", zap.String("salt", salt), zap.String("pass hash", passHash))

	if user.PasswordHash != passHash+salt {
		return nil, entity.ErrPassHashNotEqual
	}

	refreshTokenWord := gofakeit.MinecraftVillagerJob()
	tokensPair, err := a.generateUserTokenPair(user.UUID, refreshTokenWord)
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

// заглушка https://popovza.kaiten.ru/space/420777/card/37360398
// err := a.medRepo.AddMed()
//
//	if err != nil {
//		a.logger.Warn("Med Repo not implemented")
//	}
//
// medWorkerID := gofakeit.Number(0, 1<<10)
func (a *AuthUseCase) Register(ctx context.Context, request *entity.RequestRegister) (uint64, error) {
	a.logger.Debug("Register usecase started", zap.Any("Request", request))

	medWorkerID, err := a.medRepo.AddMed(ctx, request)
	if err != nil {
		a.logger.Warn("Failed to register med worker via med service", zap.Error(err))
		return 0, fmt.Errorf("failed to add med worker: %w", err)
	}
	a.logger.Info("Med worker registered successfully", zap.Int("MedWorkerID", medWorkerID))

	salt := gofakeit.MinecraftBiome()
	passHash, err := entity.HashByScrypt(request.Password, salt)
	if err != nil {
		a.logger.Error("Password hashing", zap.Error(err))
		return nil, fmt.Errorf("password hashing: %w", err)
	}

	a.logger.Info("[Request] Add new user")
	user := entity.UserCreditals{
		UUID:          UUID,
		Login:         request.Email,
		PasswordHash:  passHash + salt,
		MedWorkerUUID: MedWorkerUUID,
	}
	if _, err := a.authRepo.CreateUser(ctx, &user); err != nil {
		a.logger.Error("Create new user", zap.Error(err))
		return nil, fmt.Errorf("create new user: %w", err)
	}
	a.logger.Info("[Response] Added new user")

	return &entity.ResponseRegister{
		UUID: UUID,
	}, nil
}

func (a *AuthUseCase) TokenRefresh(ctx context.Context, request string) (*entity.TokensPair, error) {
	a.logger.Debug("Token refresh usecases started", zap.Any("Request", request))

	userData, err := a.jwtService.ParseUserData(request)
	if err != nil {
		a.logger.Error("Parse token err", zap.Error(err))
		return nil, fmt.Errorf("parse token: %w", err)
	}

	a.logger.Info("[Request] Get user by UUID")
	user, err := a.authRepo.GetUserByUUID(ctx, userData.UserUUID)
	if err != nil {
		a.logger.Error("Get user by uuid err", zap.Error(err))
		return nil, fmt.Errorf("get user by uuid: %w", err)
	}
	a.logger.Debug("[Response] Got user by UUID")

	if user.RefreshTokenWord != userData.RefreshTokenWord {
		a.logger.Warn("Tokens RTW divergents", zap.Int("userID", user.ID))
		return nil, entity.ErrExpiredSession
	}

	refreshTokenWord := gofakeit.MinecraftVillagerStation()
	tokensPair, err := a.generateUserTokenPair(user.UUID, refreshTokenWord)
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
