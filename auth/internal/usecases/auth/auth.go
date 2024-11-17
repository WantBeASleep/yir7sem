package usecases

import (
	"context"
	"fmt"
	"yir/auth/internal/entity"
	"yir/auth/internal/usecases/repositories"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthUseCase struct {
	authRepo repositories.Auth
	medRepo  repositories.MedRepo
	jwt      *entity.JWT

	logger *zap.Logger
}

func NewAuthUseCase(
	authRepo repositories.Auth,
	medRepo repositories.MedRepo,
	jwt *entity.JWT,
	logger *zap.Logger,
) *AuthUseCase {
	return &AuthUseCase{
		authRepo: authRepo,
		medRepo:  medRepo,
		jwt:      jwt,
		logger:   logger,
	}
}

func (u *AuthUseCase) Login(ctx context.Context, mail string, password string) (*entity.TokensPair, error) {
	u.logger.Info("[Request] Get user by login")
	user, err := u.authRepo.GetUserByMail(ctx, mail)
	if err != nil {
		u.logger.Error("Get user by login", zap.Error(err))
		return nil, fmt.Errorf("get user by login: %w", err)
	}
	u.logger.Debug("[Response] Got user by login")

	salt := user.PasswordHash[64:]
	passHash, err := entity.HashByScrypt(password, salt)
	if err != nil {
		u.logger.Error("Password hashing", zap.Error(err))
		return nil, fmt.Errorf("password hashing: %w", err)
	}
	u.logger.Debug("pass params", zap.String("salt", salt))

	if user.PasswordHash != passHash+salt {
		return nil, entity.ErrPassHashNotEqual
	}

	tokensPair, refreshTokenWord, err := u.generateUserTokenPair(user.Id)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}

	u.logger.Info("[Request] Update JWT tokens in DB")
	if err = u.authRepo.UpdateRefreshTokenByUserID(ctx, user.Id, refreshTokenWord); err != nil {
		u.logger.Error("Update refresh token", zap.Error(err))
		return nil, fmt.Errorf("update refresh token: %w", err)
	}
	u.logger.Info("[Response] Updated JWT tokens in DB")

	return tokensPair, nil
}

func (a *AuthUseCase) Register(ctx context.Context, req *entity.RequestRegister) (uuid.UUID, error) {
	a.logger.Debug("req", zap.Any("req", req))
	medWorkerID, _ := uuid.NewRandom()

	salt := gofakeit.MinecraftBiome()
	passHash, err := entity.HashByScrypt(req.Password, salt)
	if err != nil {
		a.logger.Error("Password hashing", zap.Error(err))
		return uuid.Nil, fmt.Errorf("password hashing: %w", err)
	}

	a.logger.Info("[Request] Add new user")
	user := entity.User{
		Mail:         req.Mail,
		PasswordHash: passHash + salt,
		MedWorkerID:  medWorkerID,
	}

	userID, err := a.authRepo.CreateUser(ctx, &user)
	if err != nil {
		a.logger.Error("Create new user", zap.Error(err))
		return uuid.Nil, fmt.Errorf("create new user: %w", err)
	}
	a.logger.Info("[Response] Added new user", zap.Any("ID", userID))

	return userID, nil
}

func (a *AuthUseCase) TokenRefresh(ctx context.Context, request string) (*entity.TokensPair, error) {
	userID, refreshTokenWord, err := a.jwt.ParseUserIDAndRW(request)
	if err != nil {
		a.logger.Error("Parse token err", zap.Error(err))
		return nil, fmt.Errorf("parse token: %w", err)
	}

	a.logger.Info("[Request] Get user by id")
	user, err := a.authRepo.GetUserByID(ctx, userID)
	if err != nil {
		a.logger.Error("Get user by uuid err", zap.Error(err))
		return nil, fmt.Errorf("get user by uuid: %w", err)
	}
	a.logger.Debug("[Response] Got user by UUID")

	if user.RefreshTokenWord != refreshTokenWord {
		a.logger.Warn("Tokens RTW divergents", zap.Any("userID", userID))
		return nil, entity.ErrExpiredSession
	}

	tokensPair, refreshTokenWord, err := a.generateUserTokenPair(userID)
	if err != nil {
		return nil, fmt.Errorf("generate tokens: %w", err)
	}

	a.logger.Info("[Request] Update user refresh token in repo")
	if err := a.authRepo.UpdateRefreshTokenByUserID(ctx, userID, refreshTokenWord); err != nil {
		return nil, fmt.Errorf("update refresh token in repo: %w", err)
	}
	a.logger.Info("[Reponse] Updated user refresh token in repo")

	return tokensPair, nil
}
