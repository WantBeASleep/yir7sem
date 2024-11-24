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

func (r *UziRepo) InsertSegments(ctx context.Context, segments []entity.Segment) error {
	segmentsDB := mappers.MustTransformSlice[entity.Segment, models.Segment](segments)
	for _, seg := range segmentsDB {
		if err := db.CreateRecord[models.Segment](ctx, r.db, &seg); err != nil {
			return fmt.Errorf("create segment: %w", err)
		}
	}

	return nil
}

func (r *UziRepo) getSegmentByID(ctx context.Context, id uuid.UUID) (*entity.Segment, error) {
	resp, err := db.GetSingleMappedRecord[entity.Segment, models.Segment](ctx, r.db, db.WithWhere("id = ?", id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) GetSegmentsByUziID(ctx context.Context, uziID uuid.UUID) ([]entity.Segment, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Distinct("segments.id", "segments.image_id", "segments.formation_id", "segments.contor", "segments.tirads_id").
		Joins("inner join images on uzis.id = images.uzi_id").
		Joins("inner join segments on images.id = segments.image_id").
		Where("uzis.id = ?", uziID)

	resp := []models.Segment{}
	if err := query.Find(&resp).Error; err != nil {
		return nil, fmt.Errorf("get uzi segments: %w", err)
	}

	return mappers.MustTransformSlice[models.Segment, entity.Segment](resp), nil
}

func (r *UziRepo) GetSegmentsByImageID(ctx context.Context, imageID uuid.UUID) ([]entity.Segment, error) {
	resp, err := db.GetMultiMappedRecord[entity.Segment, models.Segment](ctx, r.db, db.WithWhere("image_id = ?", imageID))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) GetSegmentsByFormationID(ctx context.Context, formationID uuid.UUID) ([]entity.Segment, error) {
	resp, err := db.GetMultiMappedRecord[entity.Segment, models.Segment](ctx, r.db, db.WithWhere("formation_id = ?", formationID))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) UpdateSegment(ctx context.Context, id uuid.UUID, segment *entity.Segment) (*entity.Segment, error) {
	segmentDB := mappers.MustTransformObj[entity.Segment, models.Segment](segment)

	if err := r.db.WithContext(ctx).
		Model(&models.Segment{}).
		Where("id = ?", id).
		Updates(segmentDB).
		Error; err != nil {
		return nil, fmt.Errorf("update segment: %w", err)
	}

	return r.getSegmentByID(ctx, id)
}
