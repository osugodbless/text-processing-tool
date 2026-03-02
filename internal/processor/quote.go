package processor

import "regexp"

func Quote(s *string) {
	newS := ""
	for i, ch := range *s {
		if (i == 1 && ch == 32) || (i == len(*s)-2 && ch == 32) {
			continue
		} else {
			newS += string(ch)
		}
	}
	*s = newS
}

func MatchPatternQuote(s string) string {
	re := regexp.MustCompile(`'[^']+'`)
	matches := re.FindAll([]byte(s), -1)
	if matches == nil {
		return ""
	}
	return s
}
