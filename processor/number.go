package processor

import (
	"strconv"
)

func HexToDecimal(s *string) {
	n, err := strconv.ParseInt(*s, 16, 64)
	if err != nil {
		return
	}
	*s = strconv.Itoa(int(n))
}

func BinToDecimal(s *string) {
	n, err := strconv.ParseInt(*s, 2, 64)
	if err != nil {
		return
	}
	*s = strconv.Itoa(int(n))
}
