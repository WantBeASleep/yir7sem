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
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(cfg config.DBConfig) (*Repository, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: mappers.GetDSN(cfg),
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create db gorm obj: %w", err)
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetUserByID(ctx context.Context, ID uint) (*enity.DomainUser, error) {
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

	user, err := mappers.AuthInfoToDomainUser(resp)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByLogin(ctx context.Context, login string) (*enity.DomainUser, error) {
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

	user, err := mappers.AuthInfoToDomainUser(resp)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// не уверен что стоит возвращать ID, а не DomainUser. (покурить на этот счет)
func (r *Repository) Create(ctx context.Context, user enity.DomainUser) (uint, error) {
	auth, err := mappers.DomainUserToAuthInfo(user)
	if err != nil {
		return 0, err
	}

	if err := r.db.WithContext(ctx).Create(&auth).Error; err != nil {
		return 0, err
	}
	return auth.ID, nil
}

func (r *Repository) UpdateRefreshTokenByID(ctx context.Context, ID uint, newToken string) error {
	err := r.db.WithContext(ctx).
		Where("id = ?", ID).
		Update("refresh_token", newToken).
		Error

	return err
}
