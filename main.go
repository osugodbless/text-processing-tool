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
	text := processor.ReadFile(arg[0])
	result := processor.ProcessContent(text)
	fmt.Println(result)
}
