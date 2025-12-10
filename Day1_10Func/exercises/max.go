package exercises

import "fmt"

// Bai 3 : Tim gia tri max cua array
func Max(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("no elements in the array")
	}
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max, nil
}
