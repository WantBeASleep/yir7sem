package uzirepo

import (
	"fmt"
	// "os"

	"yir/uzi/internal/config"
	"yir/uzi/internal/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "time"
	// "gorm.io/gorm/logger"
)

type ctrl struct{}

var uziRepoCtrl ctrl

// type gormLog struct {
// }

// func (l *gormLog) Printf(format string, args ...any) {
// 	fmt.Fprintf(os.Stdout, format, args...)
// }

func (r *ctrl) init(cfg *config.DB) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: cfg.GetDSN(),
	}), &gorm.Config{
		// Logger: logger.New(
		// 	&gormLog{},
		// 	logger.Config{
		// 		SlowThreshold: time.Second,
		// 		LogLevel:      logger.Info,
		// 		Colorful:      true,
		// 	},
		// ),
	})
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
