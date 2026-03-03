package processor

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func ReadFile(str string) []string {
	file, err := os.ReadFile("testdata/" + str)
	if err != nil {
		log.Fatal("Unable to read file")
	}
	fileStr := string(file)
	re := regexp.MustCompile(`\([^)]+\)|[.,!?:;]+|'|\w+`)
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
		case "(up)", "(low)", "(cap)", MatchPatternCase(str):
			r := []rune(str)
			num := ExtractDigit(str)
			n := i - 1
			for num > 0 {
				if str == "(up)" || (r[1] == 'u' && r[2] == 'p') {
					Uppercase(&s[n])
				}
				if str == "(low)" || (r[1] == 'l' && r[2] == 'o') {
					Lowercase(&s[n])
				}
				if str == "(cap)" || (r[1] == 'c' && r[2] == 'a') {
					Capitalize(&s[n])
				}

				n--
				num--
			}
			stringToDel[i] = true
		case MatchPatternPunctuation(str):
			if s[i-1] == "(cap)" || s[i-1] == "(low)" || s[i-1] == "(up)" || s[i-1] == MatchPatternCase(s[i-1]) {
				s[i-2] += s[i]
			} else {
				s[i-1] += s[i]
			}
			stringToDel[i] = true

		case "a", "A":
			if i < len(s)-1 {
				if IsVowel(&s[i+1]) {
					if s[i] == "a" {
						s[i] = "an"
					}
					if s[i] == "A" {
						s[i] = "An"
					}
				}
			}
		}
	}

	var result []string
	openQuote := false
	tempStr := ""

	for i, str := range s {
		if !stringToDel[i] {
			if str != "'" && !openQuote {
				result = append(result, str)
			} else if str == "'" && !openQuote {
				tempStr += str
				openQuote = true
			} else if str != "'" && openQuote {
				tempStr = tempStr + str + " "
			} else if str == "'" && openQuote {
				tempStr = strings.TrimSpace(tempStr)
				tempStr += str
				result = append(result, tempStr)
				tempStr = ""
				openQuote = false
			}

		}

	}

	return result
}
