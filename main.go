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

	rawText, err := processor.ReadFile(arg[0])
	if err != nil {
		log.Fatal("Unable to read file:", err)
	}

	processedText := processor.ProcessContent(rawText)
	processor.WriteToFile(processedText, arg[1])

}
