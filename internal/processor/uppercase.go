package processor

import "strings"

func Uppercase(s *string) {
	*s = strings.ToUpper(*s)
}
