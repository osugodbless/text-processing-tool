package processor

import (
	"regexp"
	"unicode"
)

func Capitalize(s *string) {
	newS := *s
	r := []rune(newS)
	r[0] = unicode.ToUpper(r[0])
	*s = string(r)
}

func MatchPatternCap(s string) string {
	re := regexp.MustCompile(`\(cap, \d+\)`)
	matches := re.FindAll([]byte(s), -1)
	if matches == nil {
		return ""
	}
	return s
}
