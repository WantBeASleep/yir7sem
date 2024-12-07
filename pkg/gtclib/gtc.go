package gtclib

func ValueToPointer[T any](v T) *T {
	p := new(T)
	*p = v
	return p
}
