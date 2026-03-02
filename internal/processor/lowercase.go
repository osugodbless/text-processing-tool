package processor

import (
	"regexp"
	"strings"
)

func Lowercase(s *string) {
	*s = strings.ToLower(*s)
}

func MatchPatternLower(s string) string {
	re := regexp.MustCompile(`\(low, \d+\)`)
	matches := re.FindAll([]byte(s), -1)
	if matches == nil {
		return ""
	}
	return s
}
