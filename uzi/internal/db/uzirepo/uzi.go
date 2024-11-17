package uzirepo

import (
	"context"
	"fmt"
	"yir/pkg/db"
	"yir/pkg/mappers"
	"yir/uzi/internal/db/models"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

func (r *UziRepo) GetDevicesList(ctx context.Context) ([]entity.Device, error) {
	resp, err := db.GetMultiMappedRecord[entity.Device, models.Device](ctx, r.db)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) CreateUzi(ctx context.Context, uzi *entity.Uzi) (uuid.UUID, error) {
	uziDB := mappers.MustTransformObj[entity.Uzi, models.Uzi](uzi)
	uziID := uuid.New()
	uziDB.Id = uziID
	if err := db.CreateRecord[models.Uzi](ctx, r.db, uziDB); err != nil {
		return uuid.Nil, err
	}

	return uziID, nil
}

func (r *UziRepo) GetUziByID(ctx context.Context, id uuid.UUID) (*entity.Uzi, error) {
	resp, err := db.GetSingleMappedRecord[entity.Uzi, models.Uzi](ctx, r.db, db.WithWhere("id = ?", id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) UpdateUzi(ctx context.Context, id uuid.UUID, uzi *entity.Uzi) (*entity.Uzi, error) {
	uziDB := mappers.MustTransformObj[entity.Uzi, models.Uzi](uzi)
	var updatedUzi models.Uzi

	if err := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Where("id = ?", id).
		Updates(uziDB).
		Find(&updatedUzi).
		Error; err != nil {
		return nil, fmt.Errorf("update uzi: %w", err)
	}

	return mappers.MustTransformObj[models.Uzi, entity.Uzi](&updatedUzi), nil
}
