// покурить над фильтрами при get

package auth

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"yir/auth/internal/config"
	"yir/auth/internal/enity"
	"yir/auth/internal/repository/db/mappers"
	"yir/auth/internal/repository/db/models"
	"yir/auth/internal/repository/db/utils"
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

	//возможно это лишнее
	db.AutoMigrate(&models.AuthInfo{})

	return &AuthRepo{
		db: db,
	}, nil
}

func (r *AuthRepo) GetUserByID(ctx context.Context, ID uint) (*enity.User, error) {
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

// не уверен что стоит возвращать ID, а не User. (покурить на этот счет)
func (r *AuthRepo) CreateUser(ctx context.Context, user *enity.User) (uint, error) {
	auth, err := mappers.UserToAuthInfo(user)
	if err != nil {
		return 0, err
	}

	if err := r.db.WithContext(ctx).Create(&auth).Error; err != nil {
		return 0, err
	}
	return uint(auth.ID), nil
}

func (r *AuthRepo) UpdateRefreshTokenByID(ctx context.Context, ID uint, newToken string) error {
	err := r.db.WithContext(ctx).
		Where("id = ?", ID).
		Update("refresh_token", newToken).
		Error

	return err
}
