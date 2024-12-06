package slicer

func Flatten2DArray[T any](slice [][]T) []T {
	cnt := 0
	for _, v := range slice {
		cnt += len(v)
	}

	res := make([]T, 0, cnt)
	for _, v := range slice {
		res = append(res, v...)
	}

	return res
}
