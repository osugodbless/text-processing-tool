package processor

import (
	"strings"
	"unicode"
)

func IsVowel(s *string) bool {
	newS := []rune(*s)
	return strings.ContainsRune("aeiouh", unicode.ToLower(newS[0]))
}
