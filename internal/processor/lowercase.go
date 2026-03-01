package processor

import "strings"

func Lowercase(s *string) {
	*s = strings.ToLower(*s)
}
