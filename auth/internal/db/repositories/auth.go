package repositories

import (
	"context"
	"fmt"

	"yir/pkg/db"
	"yir/pkg/mappers"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"yir/auth/internal/config"
	"yir/auth/internal/db/models"
	"yir/auth/internal/entity"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.DB) (*AuthRepo, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.GetDSN(),
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create db gorm obj: %w", err)
	}

	// https://popovza.kaiten.ru/space/420777/card/37587888
	db.AutoMigrate(&models.User{})

	return &AuthRepo{
		db: db,
	}, nil
}

func (r *AuthRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	resp, err := db.GetSingleMappedRecord[entity.User, models.User](ctx, r.db, db.WithWhere("id = ?", id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *AuthRepo) GetUserByMail(ctx context.Context, mail string) (*entity.User, error) {
	resp, err := db.GetSingleMappedRecord[entity.User, models.User](ctx, r.db, db.WithWhere("mail = ?", mail))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *AuthRepo) CreateUser(ctx context.Context, user *entity.User) (uuid.UUID, error) {
	userDB := mappers.MustTransformObj[entity.User, models.User](user)
	if err := db.CreateRecord[models.User](ctx, r.db, userDB); err != nil {
		return uuid.Nil, err
	}

	return user.Id, nil
}

func (r *AuthRepo) UpdateRefreshTokenByUserID(ctx context.Context, id uuid.UUID, refreshTokenWord string) error {
	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("refresh_token_word", refreshTokenWord).
		Error

	return err
}
