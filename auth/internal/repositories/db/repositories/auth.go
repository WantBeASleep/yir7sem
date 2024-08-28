package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"yir/auth/internal/config"
	"yir/auth/internal/entity"
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
	db.AutoMigrate(&models.UserCreditals{})

	return &AuthRepo{
		db: db,
	}, nil
}

func (r *AuthRepo) GetUserByID(ctx context.Context, ID int) (*entity.UserCreditals, error) {
	var resp models.UserCreditals

	query := r.db.WithContext(ctx).
		Model(&models.UserCreditals{}).
		Where("id = ?", ID)

	if err := query.Take(&resp).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}

	user, err := mappers.ModelUserCreditalsToEntityUserCreditals(&resp)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepo) GetUserByUUID(ctx context.Context, UUID uuid.UUID) (*entity.UserCreditals, error) {
	var resp models.UserCreditals

	query := r.db.WithContext(ctx).
		Model(&models.UserCreditals{}).
		Where("uuid = ?", UUID.String())

	if err := query.Take(&resp).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}

	user, err := mappers.ModelUserCreditalsToEntityUserCreditals(&resp)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepo) GetUserByLogin(ctx context.Context, login string) (*entity.UserCreditals, error) {
	var resp models.UserCreditals

	query := r.db.WithContext(ctx).
		Model(&models.UserCreditals{}).
		Where("login = ?", login)

	if err := query.Take(&resp).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}

	user, err := mappers.ModelUserCreditalsToEntityUserCreditals(&resp)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepo) CreateUser(ctx context.Context, user *entity.UserCreditals) (int, error) {
	auth, err := mappers.EntityUserCreditalsToModelUserCreditals(user)
	if err != nil {
		return 0, err
	}

	if err := r.db.WithContext(ctx).
		Model(&models.UserCreditals{}).
		Create(&auth).
		Error; err != nil {
		return 0, err
	}
	return int(auth.ID), nil
}

func (r *AuthRepo) UpdateRefreshTokenByID(ctx context.Context, ID int, refreshTokenWord string) error {
	err := r.db.WithContext(ctx).
		Model(&models.UserCreditals{}).
		Where("id = ?", ID).
		Update("refresh_token_word", refreshTokenWord).
		Error

	return err
}
