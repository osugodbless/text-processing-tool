package processor

import (
	"regexp"
	"strings"
)

func Uppercase(s *string) {
	*s = strings.ToUpper(*s)
}

func MatchPatternUpper(s string) string {
	re := regexp.MustCompile(`\(up, \d+\)`)
	matches := re.FindAll([]byte(s), -1)
	if matches == nil {
		return ""
	}
	return s
}
