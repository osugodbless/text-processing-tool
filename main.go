package main

import (
	"fmt"
	"os"
	"text-processing-tool/internal/processor"
)

func main() {
	arg := os.Args[1:]
	lenOfArgs := len(arg)

	if lenOfArgs != 2 {
		fmt.Println("Not enough or too many argument! This program requires two CLI arguments to run properly.")
		os.Exit(1)
	}
	rawText := processor.ReadFile(arg[0])
	processedText := processor.ProcessContent(rawText)
	fmt.Println(processedText)
	fmt.Println(processor.MatchPatternLower("(low, 3)"))
}
