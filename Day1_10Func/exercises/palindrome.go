package exercises

import "unicode"

// bai so 10 kiem tra chuoi doi xung
func IsPalindrome(text string) bool {
	var cleanedRunes []rune
	for _, r := range text {
		if !unicode.IsSpace(r) {
			cleanedRunes = append(cleanedRunes, unicode.ToLower(r))
		}
	}
	n := len(cleanedRunes)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if cleanedRunes[i] != cleanedRunes[j] {
			return false
		}
	}
	return true
}
