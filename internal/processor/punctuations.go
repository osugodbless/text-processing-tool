package processor

import (
	"regexp"
)

func MatchPatternPunctuation(s string) string {
	re := regexp.MustCompile(`[.,!?:;]+`)
	matches := re.FindAll([]byte(s), -1)
	if matches == nil {
		return ""
	}
	return s
}
