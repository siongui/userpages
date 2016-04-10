package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func traverse(n *html.Node) {
	if isAnchorElement(n) {
		printRstLink(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}
}

func parseCommandLineArguments() string {
	pPath := flag.String("input", "", "Path of HTML file to be processed")
	flag.Parse()
	path := *pPath
	if path == "" {
		fmt.Fprintf(os.Stderr, "Error: empty path!\n")
	}

	return path
}

func main() {
	inputFile := parseCommandLineArguments()

	fin, err := os.Open(inputFile)
	if err != nil {
		panic("Fail to open " + inputFile)
	}
	defer fin.Close()

	doc, err := html.Parse(fin)
	if err != nil {
		panic("Fail to parse " + inputFile)
	}

	traverse(doc)
}
