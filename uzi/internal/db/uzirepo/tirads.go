package uzirepo

import (
	"context"
	"fmt"
	"yir/pkg/db"
	"yir/pkg/mappers"
	"yir/uzi/internal/db/models"
	"yir/uzi/internal/entity"
)

// генериться ID на стороне БД потому что не uuid
func (r *UziRepo) CreateTirads(ctx context.Context, tirads *entity.Tirads) (int, error) {
	tiradsDB := mappers.MustTransformObj[entity.Tirads, models.Tirads](tirads)

	if err := db.CreateRecord[models.Tirads](ctx, r.db, tiradsDB); err != nil {
		return 0, err
	}

	return tiradsDB.Id, nil
}

func (r *UziRepo) GetTiradsByID(ctx context.Context, id int) (*entity.Tirads, error) {
	resp, err := db.GetSingleMappedRecord[entity.Tirads, models.Tirads](ctx, r.db, db.WithWhere("id = ?", id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) UpdateTirads(ctx context.Context, id int, tirads *entity.Tirads) (*entity.Tirads, error) {
	tiradsDB := mappers.MustTransformObj[entity.Tirads, models.Tirads](tirads)
	var updateTirads models.Tirads

	if err := r.db.WithContext(ctx).
		Model(&models.Tirads{}).
		Where("id = ?", id).
		Updates(tiradsDB).
		Find(&updateTirads).
		Error; err != nil {
		return nil, fmt.Errorf("get tirads: %w", err)
	}

	return mappers.MustTransformObj[models.Tirads, entity.Tirads](&updateTirads), nil
}
