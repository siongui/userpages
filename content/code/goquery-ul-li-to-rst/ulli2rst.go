package ulli2rst

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
)

var liMark = []string{"-", "*"}

func StringToLines(s string) []string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return lines
}

func processUl(ul *goquery.Selection, depth int) {
	ul.Find("li").Each(func(_ int, li *goquery.Selection) {
		li.Find("ul").Each(func(_ int, childUl *goquery.Selection) {
			processUl(childUl, depth+1)
		})

		lines := StringToLines(li.Text())
		var indentedLines []string
		for i, line := range lines {
			if i == 0 {
				liMarkIndex := depth % 2
				mark := liMark[liMarkIndex]
				indentedLines = append(indentedLines, "\n"+mark+" "+line)
			} else {
				indentedLines = append(indentedLines, "  "+line)
			}
		}
		li.ReplaceWithHtml(strings.Join(indentedLines, "\n"))
	})

	ul.ReplaceWithHtml(ul.Text())
}

func HtmlUlLiToRst(doc *goquery.Document) *goquery.Document {
	for ul := doc.Find("ul").First(); ul.Length() != 0; ul = doc.Find("ul").First() {
		processUl(ul, 0)
	}

	return doc
}
