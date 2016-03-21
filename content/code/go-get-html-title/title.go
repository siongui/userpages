package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func processHTML(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(err)
	}

	title := doc.Find("title").Text()
	fmt.Println(title)
}

func main() {
	pPath := flag.String("input", "", "Path of file to be processed")
	flag.Parse()
	path := *pPath
	if path == "" {
		fmt.Fprintf(os.Stderr, "Error: empty path!\n")
		return
	}

	processHTML(path)
}
