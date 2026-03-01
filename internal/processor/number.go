package processor

import (
	"log"
	"strconv"
)

func HexToDecimal(s *string) {
	n, err := strconv.ParseInt(*s, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	*s = strconv.Itoa(int(n))
}

func BinToDecimal(s string) {

}
