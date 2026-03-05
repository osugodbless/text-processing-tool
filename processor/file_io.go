package processor

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func ReadFile(str string) ([]string, error) {
	file, err := os.ReadFile("testdata/" + str)
	if err != nil {
		return []string{}, err
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

	return sliceFileStr, nil
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

func WriteToFile(text []string, filename string) {
	var outputStr string
	for i, item := range text {
		outputStr += item
		if i == 0 || i%(len(text)-1) != 0 {
			outputStr += " "
		}
	}

	err := os.WriteFile("testdata/"+filename, []byte(outputStr), 0644)
	if err != nil {
		log.Fatal("Unable to write to file:", err)
	}
}
