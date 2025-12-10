package exercises

func Flatten(input []any) []int {
	var result []int
	for _, value := range input {
		switch v := value.(type) {
		case int:
			result = append(result, v)
		case []any:
			flat := Flatten(v)
			result = append(result, flat...)
		}
	}
	return result
}
