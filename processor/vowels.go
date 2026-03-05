package processor

import (
	"strings"
	"unicode"
)

func IsVowel(s *string) bool {
	if *s == "" {
		return false
	}
	newS := []rune(*s)
	return strings.ContainsRune("aeiouh", unicode.ToLower(newS[0]))
}
