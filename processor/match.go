package processor

import (
	"regexp"
)

func MatchPatternCase(s string) string {
	re := regexp.MustCompile(`\(up, \d+\)|\(low, \d+\)|\(cap, \d+\)`)
	matches := re.FindAll([]byte(s), -1)
	if matches != nil {
		return s
	}

	return ""
}

func MatchPatternQuote(s string) string {
	re := regexp.MustCompile(`'[^']+'`)
	matches := re.FindAll([]byte(s), -1)
	if matches == nil {
		return ""
	}
	return s
}

func MatchPatternPunctuation(s string) string {
	re := regexp.MustCompile(`[.,!?:;]+`)
	match := re.FindAll([]byte(s), -1)
	if match == nil {
		return ""
	}
	return s
}
