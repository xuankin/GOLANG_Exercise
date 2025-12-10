package exercises

import (
	"strings"
	"unicode"
)

// bai so 9 tinh so tu trong 1 cau
func CountWords(text string) map[string]int {
	counts := make(map[string]int)
	cleanText := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || unicode.IsSpace(r) || unicode.IsLetter(r) {
			return r
		}
		return ' '
	}, text)
	words := strings.Fields(cleanText)
	for _, word := range words {
		if word != "" {
			counts[word]++
		}
	}
	return counts
}
