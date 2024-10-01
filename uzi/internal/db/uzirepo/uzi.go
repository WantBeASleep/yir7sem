package uzirepo

import (
	"context"
	"errors"
	"fmt"
	"yir/pkg/db"
	"yir/pkg/mappers"
	"yir/uzi/internal/config"
	"yir/uzi/internal/db/models"
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (r *UziRepo) GetDevicesList(ctx context.Context) ([]entity.Device, error) {
	resp, err := db.GetMultiMappedRecord[entity.Device, models.Device](ctx, r.db)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

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

func (r *UziRepo) UpdateTirads(ctx context.Context, id int, tirads *entity.Tirads) error {
	tiradsDB := mappers.MustTransformObj[entity.Tirads, models.Tirads](tirads)

	err := r.db.WithContext(ctx).
		Model(&models.Tirads{}).
		Where("id = ?", id).
		Updates(tiradsDB).
		Error

	return err
}

func (r *UziRepo) InsertUzi(ctx context.Context, uzi *entity.Uzi) error {
	return db.CreateRecord[models.Uzi](ctx, r.db, uzi)
}

func (r *UziRepo) GetUziByID(ctx context.Context, id uuid.UUID) (*entity.Uzi, error) {
	resp, err := db.GetSingleMappedRecord[entity.Uzi, models.Uzi](ctx, r.db, db.WithWhere("id = ?", id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) UpdateUzi(ctx context.Context, id uuid.UUID, uzi *entity.Uzi) error {
	uziDB := mappers.MustTransformObj[entity.Uzi, models.Uzi](uzi)

	err := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Where("id = ?", id).
		Updates(uziDB).
		Error

	return err
}

func (r *UziRepo) InsertImages(ctx context.Context, images []entity.Image) error {
	return db.CreateRecord[models.Image](ctx, r.db, images)
}

func (r *UziRepo) GetUziImages(ctx context.Context, uziID uuid.UUID) ([]entity.Image, error) {
	resp, err := db.GetMultiMappedRecord[entity.Image, models.Image](ctx, r.db, db.WithWhere("uzi_id = ?", uziID))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) InsertFormations(ctx context.Context, formations []entity.Formation) error {
	return db.CreateRecord[models.Formation](ctx, r.db, formations)
}

func (r *UziRepo) GetUziFormations(ctx context.Context, uziID uuid.UUID) ([]entity.Formation, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Distinct("formations.id").
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

func (r *UziRepo) GetFormationByID(ctx context.Context, id uuid.UUID) (*entity.Formation, error) {
	resp, err := db.GetSingleMappedRecord[entity.Formation, models.Formation](ctx, r.db, db.WithWhere("id = ?", id))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UziRepo) InsertSegments(ctx context.Context, segments []entity.Segment) error {
	return db.CreateRecord[models.Segment](ctx, r.db, segments)
}

func (r *UziRepo) GetUziSegments(ctx context.Context, uziID uuid.UUID) ([]entity.Segment, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Uzi{}).
		Distinct("segments.id").
		Joins("inner join images on uzis.id = images.uzi_id").
		Joins("inner join segments on images.id = segments.image_id").
		Where("uzis.id = ?", uziID)

	resp := []models.Segment{}
	if err := query.Find(&resp).Error; err != nil {
		return nil, fmt.Errorf("get uzi segments: %w", err)
	}

	return mappers.MustTransformSlice[models.Segment, entity.Segment](resp), nil
}

func (r *UziRepo) GetFormationSegments(ctx context.Context, formationID uuid.UUID) ([]entity.Segment, error) {
	query := r.db.WithContext(ctx).
		Model(&models.Segment{}).
		Distinct("segments.id").
		Joins("inner join formations on segments.formation_id = formations.id").
		Where("formations.id = ?", formationID)

	resp := []models.Segment{}
	if err := query.Find(&resp).Error; err != nil {
		return nil, fmt.Errorf("get uzi segments: %w", err)
	}

	return mappers.MustTransformSlice[models.Segment, entity.Segment](resp), nil
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

func (r *UziRepo) CheckImagesIDsExist(ctx context.Context, ids uuid.UUIDs) error {
	var imagesIDs uuid.UUIDs
	if err := r.db.WithContext(ctx).Model(&models.Image{}).Pluck("id", &imagesIDs).Error; err != nil {
		return fmt.Errorf("get images ids: %w", err)
	}

	var notFoundIds uuid.UUIDs
	imagesIDsMap := mapper.SliceToMap(imagesIDs)
	for _, id := range ids {
		if _, ok := imagesIDsMap[id]; !ok {
			notFoundIds = append(notFoundIds, id)
		}
	}

	if notFoundIds != nil {
		return &entity.ImagesNotFoundError{Ids: notFoundIds}
	}

	return nil
}
