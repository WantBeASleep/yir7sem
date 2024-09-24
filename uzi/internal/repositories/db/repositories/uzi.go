package repositories

import (
	"context"
	"fmt"
	"yir/uzi/internal/config"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/repositories/db/models"
	"yir/uzi/internal/repositories/db/utils"
	mapper "yir/uzi/internal/utils"

	"errors"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UziRepo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.DB) (*UziRepo, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: utils.GetDSN(cfg),
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create db gorm obj: %w", err)
	}

	// https://popovza.kaiten.ru/space/420777/card/37587888
	db.AutoMigrate(
		&models.Device{},
		&models.Uzi{},
		&models.Image{},
		&models.Tirads{},
		&models.Formation{},
		&models.ImageFormation{},
	)

	return &UziRepo{
		db: db,
	}, nil
}

func (r *UziRepo) GetDevice(ctx context.Context, id int) (*entity.Device, error) {
	var resp models.Device

	query := r.db.WithContext(ctx).
		Model(&models.Device{}).
		Where("id = ?", id)

	if err := query.Take(&resp).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}

	return mapper.MustTransformObj[models.Device, entity.Device](&resp), nil
}

func (r *UziRepo) GetDevicesList(ctx context.Context) ([]entity.Device, error) {
	var resp []models.Device

	err := r.db.WithContext(ctx).
		Model(&models.Device{}).
		Find(&resp).
		Error

	if err != nil {
		return nil, err
	}

	return mapper.MustTransformSlice[models.Device, entity.Device](resp), nil
}

func (r *UziRepo) CreateTirads(ctx context.Context, tirads *entity.Tirads) (int, error) {
	tiradsDB := mapper.MustTransformObj[entity.Tirads, models.Tirads](tirads)

	if err := r.db.WithContext(ctx).
		Model(&models.Tirads{}).
		Create(tiradsDB).
		Error; err != nil {
		return 0, err
	}

	return tiradsDB.Id, nil
}

func (r *UziRepo) UpdateTirads(ctx context.Context, id int, tirads *entity.Tirads) error {
	tiradsDB := mapper.MustTransformObj[entity.Tirads, models.Tirads](tirads)

	err := r.db.WithContext(ctx).
		Model(&models.Tirads{}).
		Where("id = ?", id).
		Updates(tiradsDB).
		Error

	return err
}

func (r *UziRepo) InsertUzi(ctx context.Context, uzi *entity.Uzi) error {
	uziDB := mapper.MustTransformObj[entity.Uzi, models.Uzi](uzi)

	err := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Create(uziDB).
		Error

	return err
}

func (r *UziRepo) UpdateUzi(ctx context.Context, uzi *entity.Uzi) error {
	uziDB := mapper.MustTransformObj[entity.Uzi, models.Uzi](uzi)

	err := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Where("id = ?", uziDB.Id).
		Updates(uziDB).
		Error

	return err
}

func (r *UziRepo) GetUzi(ctx context.Context, id uuid.UUID) (*entity.Uzi, error) {
	var resp models.Uzi

	query := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Where("id = ?", id)

	if err := query.Take(&resp).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}

	return mapper.MustTransformObj[models.Uzi, entity.Uzi](&resp), nil
}

func (r *UziRepo) InsertImages(ctx context.Context, images []entity.Image) error {
	imagesDB := mapper.MustTransformSlice[entity.Image, models.Image](images)

	err := r.db.WithContext(ctx).
		Model(&models.Image{}).
		Create(imagesDB).
		Error

	return err
}

func (r *UziRepo) GetImage(ctx context.Context, id uuid.UUID) (*entity.Image, error) {
	var resp models.Image

	query := r.db.WithContext(ctx).
		Model(&models.Image{}).
		Where("id = ?", id)

	if err := query.Take(&resp).Error; err != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}

	return mapper.MustTransformObj[models.Image, entity.Image](&resp), nil
}
