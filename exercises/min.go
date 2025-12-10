package exercises

import "fmt"

// Bai 4 Tim gia tri min cua 1 array
func Min(arr []int) (int, error) {
	if len(arr) < 1 {
		return 0, fmt.Errorf("no elements in the array")
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, nil
}
