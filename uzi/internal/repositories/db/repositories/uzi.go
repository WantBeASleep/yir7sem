package repositories

import (
	"context"
	"fmt"
	"yir/uzi/internal/config"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/repositories/db/mappers"
	"yir/uzi/internal/repositories/db/models"
	"yir/uzi/internal/repositories/db/utils"

	"errors"
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

var entityMapper = mappers.EntityToModel{}
var modelMapper = mappers.ModelToEntity{}

func (r *UziRepo) CreateDevice(ctx context.Context, device *entity.Device) error {
	deviceDB := entityMapper.Device(device)

	err := r.db.WithContext(ctx).
		Model(&models.Device{}).
		Create(deviceDB).
		Error; 
	
	return err
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

	device := modelMapper.Device(&resp)
	return device, nil
}

func (r *UziRepo) DeleteDevice(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).
		Delete(&models.Device{}, id).
		Error;

	return err
}

func (r *UziRepo) InsertTirads(ctx context.Context, tirads *entity.Tirads) (int, error) {
	tiradsDB := entityMapper.Tirads(tirads)

	if err := r.db.WithContext(ctx).
		Model(&models.Tirads{}).
		Create(tiradsDB).
		Error; 
	err != nil {
		return 0, err
	}

	return int(tiradsDB.Id), nil
}

func (r *UziRepo) UpdateTirads(ctx context.Context, id int, tirads *entity.Tirads) error {
	tiradsDB := entityMapper.Tirads(tirads)

	err := r.db.WithContext(ctx).
		Model(&models.Tirads{}).
		Where("id = ?", id).
		Updates(tiradsDB).
		Error
	
	return err
}

// func (r *UziRepo) InsertUzi(ctx context.Context, )
