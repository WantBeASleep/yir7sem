package utils

func SliceToMap[T comparable](src []T) map[T]struct{} {
	dst := make(map[T]struct{}, len(src))

	for _, v := range src {
		dst[v] = struct{}{}
	}

	return dst
}

func MapToSlice[T comparable](src map[T]struct{}) []T {
	dst := make([]T, 0, len(src))
	for k := range src {
		dst = append(dst, k)
	}

	return dst
}