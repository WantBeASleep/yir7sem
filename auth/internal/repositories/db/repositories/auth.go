package repositories

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"yir/auth/internal/config"
	"yir/auth/internal/enity"
	"yir/auth/internal/repositories/db/mappers"
	"yir/auth/internal/repositories/db/models"
	"yir/auth/internal/repositories/db/utils"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.DB) (*AuthRepo, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: utils.GetDSN(cfg),
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create db gorm obj: %w", err)
	}

	// https://popovza.kaiten.ru/space/420777/card/37587888
	db.AutoMigrate(&models.AuthInfo{})

	return &AuthRepo{
		db: db,
	}, nil
}

func (r *AuthRepo) GetUserByID(ctx context.Context, ID int) (*enity.User, error) {
	var resp models.AuthInfo

	query := r.db.WithContext(ctx).
		Model(&models.AuthInfo{}).
		Where("id = ?", ID)

	if err := query.Take(&resp).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, enity.ErrNotFound
		}
		return nil, err
	}

	user, err := mappers.AuthInfoToUser(&resp)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepo) GetUserByLogin(ctx context.Context, login string) (*enity.User, error) {
	var resp models.AuthInfo

	query := r.db.WithContext(ctx).
		Model(&models.AuthInfo{}).
		Where("login = ?", login)

	if err := query.Take(&resp).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, enity.ErrNotFound
		}
		return nil, err
	}

	user, err := mappers.AuthInfoToUser(&resp)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepo) CreateUser(ctx context.Context, user *enity.User) (int, error) {
	auth, err := mappers.UserToAuthInfo(user)
	if err != nil {
		return 0, err
	}

	if err := r.db.WithContext(ctx).
		Model(&models.AuthInfo{}).
		Create(&auth).
		Error; err != nil {
		return 0, err
	}
	return int(auth.ID), nil
}

func (r *AuthRepo) UpdateRefreshTokenByID(ctx context.Context, ID int, refreshTokenWord string) error {
	err := r.db.WithContext(ctx).
		Model(&models.AuthInfo{}).
		Where("id = ?", ID).
		Update("refresh_token_word", refreshTokenWord).
		Error

	return err
}
