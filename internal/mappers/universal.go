package mappers

import (
	"github.com/jinzhu/copier"
)

// можно ужать в одну функцию, но писать утверждения типа на каждый маппер такая себе идея

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