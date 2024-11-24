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

func (r *UziRepo) InsertFormation(ctx context.Context, formation *entity.Formation) error {
	formationDB := mappers.MustTransformObj[entity.Formation, models.Formation](formation)
	if err := db.CreateRecord[models.Formation](ctx, r.db, formationDB); err != nil {
		return fmt.Errorf("create formation: %w", err)
	}

	return nil
}

func (r *UziRepo) GetFormationByID(ctx context.Context, id uuid.UUID) (*entity.Formation, error) {
	resp, err := db.GetSingleMappedRecord[entity.Formation, models.Formation](ctx, r.db, db.WithWhere("id = ?", id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// получить все formation этого узи
func (r *UziRepo) GetFormationsByUziID(ctx context.Context, uziID uuid.UUID) ([]entity.Formation, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Distinct("formations.id", "formations.ai", "formations.tirads_id").
		Joins("inner join images on uzis.id = images.uzi_id").
		Joins("inner join segments on images.id = segments.image_id").
		Joins("inner join formations on segments.formation_id = formations.id").
		Where("uzis.id = ?", uziID)

	resp := []models.Formation{}
	if err := query.Find(&resp).Error; err != nil {
		return nil, fmt.Errorf("get uzi formations: %w", err)
	}

	return mappers.MustTransformSlice[models.Formation, entity.Formation](resp), nil
}

// получить все formation на этой картинке
func (r *UziRepo) GetFormationsByImageID(ctx context.Context, imageID uuid.UUID) ([]entity.Formation, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Image{}).
		Distinct("formations.id", "formations.ai", "formations.tirads_id").
		Joins("inner join segments on images.id = segments.image_id").
		Joins("inner join formations on segments.formation_id = formations.id").
		Where("images.id = ?", imageID)

	resp := []models.Formation{}
	if err := query.Find(&resp).Error; err != nil {
		return nil, fmt.Errorf("get image formations: %w", err)
	}

	return mappers.MustTransformSlice[models.Formation, entity.Formation](resp), nil
}

func (r *UziRepo) UpdateFormation(ctx context.Context, id uuid.UUID, formation *entity.Formation) (*entity.Formation, error) {
	formationDB := mappers.MustTransformObj[entity.Formation, models.Formation](formation)

	if err := r.db.WithContext(ctx).
		Model(&models.Formation{}).
		Where("id = ?", id).
		Updates(formationDB).
		Error; err != nil {
		return nil, fmt.Errorf("update formation: %w", err)
	}

	return r.GetFormationByID(ctx, id)
}
