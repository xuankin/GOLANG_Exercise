package exercises

// Bai so 6 :ham map
func Map[T any, R any](input []T, fn func(T) R) []R {
	result := make([]R, len(input))
	for i, value := range input {
		result[i] = fn(value)
	}
	return result
}
