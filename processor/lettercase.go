package processor

import (
	"strings"
	"unicode"
)

func Lowercase(s *string) {
	*s = strings.ToLower(*s)
}

func Uppercase(s *string) {
	*s = strings.ToUpper(*s)
}

func Capitalize(s *string) {
	if *s == "" {
		return
	}
	newS := *s
	r := []rune(newS)
	r[0] = unicode.ToUpper(r[0])
	*s = string(r)
}
