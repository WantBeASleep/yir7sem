package uzirepo

import (
	"fmt"
	"yir/uzi/internal/config"
	"yir/uzi/internal/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ctrl struct{}

var uziRepoCtrl ctrl

func (r *ctrl) init(cfg *config.DB) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.GetDSN(),
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("create db gorm obj: %w", err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// https://popovza.kaiten.ru/space/420777/card/37587888
	db.AutoMigrate(
		&models.Device{},
		&models.Uzi{},
		&models.Image{},
		&models.Tirads{},
		&models.Formation{},
		&models.Segment{},
	)

	return db, nil
}

type UziRepo struct {
	db *gorm.DB
}

func NewRepository(cfg *config.DB) (*UziRepo, error) {
	db, err := uziRepoCtrl.init(cfg)
	if err != nil {
		return nil, fmt.Errorf("init repo layer: %w", err)
	}

	return &UziRepo{
		db: db,
	}, nil
}
