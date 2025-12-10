package exercises

func Reduce[T any, R any](arr []T, init R, fn func(R, T) R) R {
	acc := init
	for _, val := range arr {
		acc = fn(acc, val)
	}
	return acc
}
