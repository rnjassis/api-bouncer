package utils

func RemoveIndex[T any](orig []T, idx int) []T {
	res := make([]T, 0)
	res = append(res, orig[:idx]...)
	return append(res, orig[idx+1:]...)
}
