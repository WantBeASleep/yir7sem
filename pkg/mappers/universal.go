package mappers

import (
	"fmt"

	"github.com/jinzhu/copier"
)

// Если это поход в БД используйте не MUST опции!
type MustTransformOpts[S any, T any] func(src *S, dst *T)

// можно ужать в одну функцию, но писать утверждения типа на каждый маппер такая себе идея

func MustTransformObj[S any, T any](src *S, opts ...MustTransformOpts[S, T]) *T {
	var dst T
	copier.Copy(&dst, src)

	for _, opt := range opts {
		opt(src, &dst)
	}

	return &dst
}

// опция применяется к каждому конкретному значению, а не к всему слайсу
func MustTransformSlice[S any, T any](src []S, opts ...MustTransformOpts[S, T]) []T {
	var dst []T
	copier.Copy(&dst, &src)

	for i := range src {
		for _, opt := range opts {
			opt(&src[i], &dst[i])
		}
	}

	return dst
}

type TransformOpts[S any, T any] func(src *S, dst *T) error

func TransformObj[S any, T any](src *S, opts ...TransformOpts[S, T]) (*T, error) {
	var dst T
	if err := copier.Copy(&dst, src); err != nil {
		return nil, fmt.Errorf("copier: %w", err)
	}

	for i, opt := range opts {
		if err := opt(src, &dst); err != nil {
			return nil, fmt.Errorf("use %d opt: %w", i, err)
		}
	}

	return &dst, nil
}

func TransformSlice[S any, T any](src []S, opts ...TransformOpts[S, T]) ([]T, error) {
	var dst []T
	if err := copier.Copy(&dst, src); err != nil {
		return nil, fmt.Errorf("copier: %w", err)
	}

	for i := range src {
		for _, opt := range opts {
			if err := opt(&src[i], &dst[i]); err != nil {
				return nil, fmt.Errorf("use %d opt: %w", i, err)
			}
		}
	}

	return dst, nil
}