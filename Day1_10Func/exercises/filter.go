package exercises

func Filter[T any](input []T, fn func(T) bool) []T {
	var result []T
	for _, v := range input {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
