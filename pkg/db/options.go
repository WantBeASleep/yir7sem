package db

import (
	"gorm.io/gorm"
)

type QueryOption func(db *gorm.DB) *gorm.DB

var (
	WithWhere = func(condition string, args ...any) QueryOption {
		return QueryOption(func(db *gorm.DB) *gorm.DB {
			return db.Where(condition, args...)
		})
	}
)
