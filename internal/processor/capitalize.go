package processor

import "strings"

func Capitalize(s *string) {
	*s = strings.ToUpper(*s)
}
