package utils

import (
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
