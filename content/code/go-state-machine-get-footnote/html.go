package main

import (
	"flag"
	"fmt"
	"os"
)

func parseCommandLineArguments() string {
	pPath := flag.String("input", "", "Path of HTML file to be processed")
	flag.Parse()
	path := *pPath
	if path == "" {
		fmt.Fprintf(os.Stderr, "Error: empty input file path!\n")
	}

	return path
}

func main() {
	inputFilePath := parseCommandLineArguments()

	f, err := os.Open(inputFilePath)
	if err != nil {
		panic("Fail to open " + inputFilePath)
	}
	defer f.Close()

	footnoteBody := extractFootnote(f)
	fmt.Println(footnoteBody)
}
