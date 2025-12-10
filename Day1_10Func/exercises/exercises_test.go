package exercises

import (
	"reflect"
	"testing"
)

// 1. Test cho ReverseString
func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Normal string", "hello", "olleh"},
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
		{"Unicode string", "xin chào", "oàhc nix"},
		{"Palindrome", "racecar", "racecar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseString(tt.input)
			if got != tt.expected {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// 2. Test cho UniqueArray
func TestUniqueArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"No duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"Has duplicates", []int{1, 2, 2, 3, 1}, []int{1, 2, 3}},
		{"All same elements", []int{5, 5, 5}, []int{5}},
		{"Empty array", []int{}, []int{}},
		{"Unsorted mixed", []int{10, 1, 10, 2}, []int{10, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueArray(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("UniqueArray(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// 3. Test cho Max
func TestMax(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expected  int
		expectErr bool
	}{
		{"Positive numbers", []int{1, 5, 3}, 5, false},
		{"Negative numbers", []int{-10, -5, -20}, -5, false},
		{"Mixed numbers", []int{-5, 0, 5}, 5, false},
		{"Single element", []int{100}, 100, false},
		{"Empty array", []int{}, 0, true}, // Error expected
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Max(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("Max() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && got != tt.expected {
				t.Errorf("Max() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 4. Test cho Min
func TestMin(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expected  int
		expectErr bool
	}{
		{"Positive numbers", []int{5, 1, 3}, 1, false},
		{"Negative numbers", []int{-10, -5, -20}, -20, false},
		{"Mixed numbers", []int{-5, 0, 5}, -5, false},
		{"Single element", []int{42}, 42, false},
		{"Empty array", []int{}, 0, true}, // Error expected
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Min(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("Min() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && got != tt.expected {
				t.Errorf("Min() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 5. Test cho Map
func TestMap(t *testing.T) {
	// Sub-test 1: Int -> Int
	t.Run("Map Integers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected []int
		}{
			{"Double values", []int{1, 2, 3}, []int{2, 4, 6}},
			{"Negative values", []int{-1, -2}, []int{-2, -4}},
			{"Empty slice", []int{}, []int{}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Map(tt.input, func(x int) int { return x * 2 })
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Got %v, want %v", got, tt.expected)
				}
			})
		}
	})

	// Sub-test 2: String -> Int (đảm bảo đủ 3 case tổng thể)
	t.Run("Map String to Int", func(t *testing.T) {
		input := []string{"a", "ab", "abc"}
		expected := []int{1, 2, 3}
		got := Map(input, func(s string) int { return len(s) })
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, want %v", got, expected)
		}
	})
}

// 6. Test cho Filter
func TestFilter(t *testing.T) {
	// Table-driven cho Int
	t.Run("Filter Integers", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected []int
		}{
			{"Keep Evens", []int{1, 2, 3, 4}, []int{2, 4}},
			{"All Match", []int{2, 4, 6}, []int{2, 4, 6}},
			{"None Match", []int{1, 3, 5}, nil}, // Filter thường trả về nil nếu không append gì vào slice rỗng
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Filter(tt.input, func(n int) bool { return n%2 == 0 })
				// Xử lý trường hợp nil vs empty slice để so sánh chính xác
				if len(got) == 0 && len(tt.expected) == 0 {
					return
				}
				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("Got %v, want %v", got, tt.expected)
				}
			})
		}
	})

	// Test thêm với String cho đa dạng
	t.Run("Filter Strings", func(t *testing.T) {
		input := []string{"hi", "hello", "go"}
		expected := []string{"hello"}
		got := Filter(input, func(s string) bool { return len(s) > 3 })
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v, want %v", got, expected)
		}
	})
}

// 7. Test cho Reduce
func TestReduce(t *testing.T) {
	// Table-driven chung
	// Lưu ý: Do generic T và R thay đổi, ta chia thành các t.Run con nhưng vẫn giữ cấu trúc 3 cases

	t.Run("Sum Integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		expected := 10
		got := Reduce(input, 0, func(acc, val int) int { return acc + val })
		if got != expected {
			t.Errorf("Sum = %v, want %v", got, expected)
		}
	})

	t.Run("Concat Strings", func(t *testing.T) {
		input := []string{"Go", "Lang"}
		expected := "GoLang"
		got := Reduce(input, "", func(acc, val string) string { return acc + val })
		if got != expected {
			t.Errorf("Concat = %v, want %v", got, expected)
		}
	})

	t.Run("Reduce with Initial Value", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := 16 // 10 + 1 + 2 + 3
		got := Reduce(input, 10, func(acc, val int) int { return acc + val })
		if got != expected {
			t.Errorf("WithInit = %v, want %v", got, expected)
		}
	})
}

// 8. Test cho Flatten
func TestFlatten(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		expected []int
	}{
		{"Flat list", []any{1, 2, 3}, []int{1, 2, 3}},
		{"Nested list", []any{1, []any{2, 3}, 4}, []int{1, 2, 3, 4}},
		{"Deep nested", []any{1, []any{2, []any{3}}, 4}, []int{1, 2, 3, 4}},
		{"Empty sub-arrays", []any{1, []any{}, 2}, []int{1, 2}},
		{"Mixed ignored types", []any{1, "string", 2}, []int{1, 2}}, // Giả sử hàm chỉ lấy int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Flatten(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Flatten() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 9. Test cho CountWords
func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{"Simple sentence", "hello world hello", map[string]int{"hello": 2, "world": 1}},
		{"With punctuation", "hello, world!", map[string]int{"hello": 1, "world": 1}},
		{"Empty string", "", map[string]int{}},
		{"Multiple spaces", "a   b  a", map[string]int{"a": 2, "b": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("CountWords() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 10. Test cho IsPalindrome
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Palindrome simple", "madam", true},
		{"Not Palindrome", "hello", false},
		{"Case insensitive", "Madam", true},
		{"With spaces", "nurses run", true},
		{"Empty string", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
