package mappers

import (
	"fmt"

	"github.com/jinzhu/copier"
)

func MustTransformObj[S any, T any](src *S) *T {
	var dst T
	copier.Copy(&dst, src)

	return &dst
}

func MustTransformSlice[S any, T any](src []S) []T {
	var dst []T
	copier.Copy(&dst, &src)

	return dst
}

func TransformObj[S any, T any](src *S) (*T, error) {
	var dst T
	if err := copier.Copy(&dst, src); err != nil {
		return nil, fmt.Errorf("copier: %w", err)
	}

	return &dst, nil
}

func TransformSlice[S any, T any](src []S) ([]T, error) {
	var dst []T
	if err := copier.Copy(&dst, src); err != nil {
		return nil, fmt.Errorf("copier: %w", err)
	}

	return dst, nil
}
