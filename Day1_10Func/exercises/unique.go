package exercises

// Bai 2 Unique Array

func UniqueArray(arr []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0)
	for _, v := range arr {
		_, ok := seen[v]
		if !ok {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}
