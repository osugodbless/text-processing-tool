package processor

import (
	"log"
	"os"
	"regexp"
)

func ReadFile(str string) []string {
	file, err := os.ReadFile("testdata/" + str)
	if err != nil {
		log.Fatal("Unable to read file")
	}
	fileStr := string(file)
	re := regexp.MustCompile(`\(+[^)]+\)|\S+`)
	matches := re.FindAll([]byte(fileStr), -1)

	var sliceFileStr []string

	for _, item := range matches {
		sliceFileStr = append(sliceFileStr, string(item))
	}

	return sliceFileStr
}
