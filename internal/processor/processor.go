package processor

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func ReadFile(str string) []string {
	file, err := os.ReadFile("testdata/" + str)
	if err != nil {
		log.Fatal("Unable to read file")
	}
	fileStr := string(file)
	re := regexp.MustCompile(`\([^)]+\)|\S+`)
	matches := re.FindAll([]byte(fileStr), -1)

	var sliceFileStr []string

	for _, item := range matches {
		if len(item) > 0 {
			sliceFileStr = append(sliceFileStr, string(item))
		}
	}

	return sliceFileStr
}

func ExtractDigit(s string) int {
	r := []rune(s)
	n := ""
	for _, ch := range r {
		if unicode.IsDigit(ch) {
			n += string(ch)
		}
	}
	if n == "" {
		return 1
	}
	num, _ := strconv.Atoi(n)
	return num
}

func ProcessContent(s []string) []string {
	stringToDel := make(map[int]bool)

	for i, str := range s {
		switch str {
		case "(hex)":
			HexToDecimal(&s[i-1])
			stringToDel[i] = true
		case "(bin)":
			BinToDecimal(&s[i-1])
			stringToDel[i] = true
		case "(up)", MatchPatternUpper(str):
			num := ExtractDigit(str)
			n := i - 1
			for num > 0 {
				Uppercase(&s[n])
				n--
				num--
			}
			stringToDel[i] = true
		case "(low)", MatchPatternLower(str):
			num := ExtractDigit(str)
			n := i - 1
			for num > 0 {
				Lowercase(&s[n])
				n--
				num--
			}
			stringToDel[i] = true
		case "(cap)", MatchPatternCap(str):
			num := ExtractDigit(str)
			n := i - 1
			for num > 0 {
				Capitalize(&s[n])
				n--
				num--
			}
			stringToDel[i] = true
		case ".", ",", "!", "?", ":", ";":

		}
	}

	var result []string
	for i, str := range s {
		if !stringToDel[i] {
			result = append(result, str)
		}
	}

	return result
}
