package db

import (
	"context"

	"gorm.io/gorm"
	"yir/pkg/mappers"
)

func GetSingleMappedRecord[Entity, Model any](ctx context.Context, db *gorm.DB, opts ...QueryOption) (*Entity, error) {
	resp, err := GetSingleRecord[Model](ctx, db, opts...)
	if err != nil {
		return nil, err
	}

	return mappers.MustTransformObj[Model, Entity](resp), nil
}

func GetMultiMappedRecord[Entity, Model any](ctx context.Context, db *gorm.DB, opts ...QueryOption) ([]Entity, error) {
	resp, err := GetMultiRecord[Model](ctx, db, opts...)
	if err != nil {
		return nil, err
	}

	return mappers.MustTransformSlice[Model, Entity](resp), nil
}