package db

import (
	"context"

	"gorm.io/gorm"
)

func GetSingleRecord[T any](ctx context.Context, db *gorm.DB, opts ...QueryOption) (*T, error) {
	var resp T

	query := db.WithContext(ctx).Model(&resp)
	for _, opt := range opts {
		query = opt(query)
	}

	if err := query.Take(&resp).Error; err != nil {
		return nil, err
	}

	return &resp, nil
}

func GetMultiRecord[T any](ctx context.Context, db *gorm.DB, opts ...QueryOption) ([]T, error) {
	var model T
	var resp []T

	query := db.WithContext(ctx).Model(&model)
	for _, opt := range opts {
		query = opt(query)
	}

	if err := query.Find(&resp).Error; err != nil {
		return nil, err
	}

	return resp, nil
}

// single --> value = &T
// multi --> value = []T
func CreateRecord[T any](ctx context.Context, db *gorm.DB, value any, opts ...QueryOption) error {
	var model T

	query := db.WithContext(ctx).Model(&model)
	for _, opt := range opts {
		query = opt(query)
	}

	return query.Create(&value).Error
}