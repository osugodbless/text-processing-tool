package main

import (
	"fmt"
	"go-reloaded/processor"
	"log"
	"os"
)

func main() {
	arg := os.Args[1:]
	lenOfArgs := len(arg)

	if lenOfArgs != 2 {
		fmt.Println("Not enough or too many argument! This program requires two CLI arguments to run properly.")
		os.Exit(1)
	}

	rawText := processor.ReadFile(arg[0])
	fmt.Println(rawText)
	processedText := processor.ProcessContent(rawText)

	var outputStr string
	for i, item := range processedText {
		outputStr += item
		if i == 0 || i%(len(processedText)-1) != 0 {
			outputStr += " "
		}
	}

	err := os.WriteFile("testdata/"+arg[1], []byte(outputStr), 0644)
	if err != nil {
		log.Fatal("Unable to write to file:", err)
	}

}
