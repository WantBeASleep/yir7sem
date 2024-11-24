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

func (r *UziRepo) InsertImages(ctx context.Context, images []entity.Image) error {
	imagesDB := mappers.MustTransformSlice[entity.Image, models.Image](images)
	for _, img := range imagesDB {
		if err := db.CreateRecord[models.Image](ctx, r.db, &img); err != nil {
			return fmt.Errorf("create image: %w", err)
		}
	}

	return nil
}

func (r *UziRepo) GetImagesByUziID(ctx context.Context, uziID uuid.UUID) ([]entity.Image, error) {
	resp, err := db.GetMultiMappedRecord[entity.Image, models.Image](ctx, r.db, db.WithWhere("uzi_id = ?", uziID))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) GetImageByID(ctx context.Context, id uuid.UUID) (*entity.Image, error) {
	resp, err := db.GetSingleMappedRecord[entity.Image, models.Image](ctx, r.db, db.WithWhere("id = ?", id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
