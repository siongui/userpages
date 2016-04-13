package table2rst

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
)

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

func rstListTablePrefixOfEachLine(indexOfTd, indexOfLine int) string {
	if indexOfTd == 0 {
		if indexOfLine == 0 {
			return "   * - "
		} else {
			return "       "
		}
	} else {
		if indexOfLine == 0 {
			return "     - "
		} else {
			return "       "
		}
	}
}

func processTr(tr *goquery.Selection, fRstOutput *os.File) {
	tr.Find("td").Each(func(indexOfTd int, td *goquery.Selection) {
		lines := StringToLines(td.Text())
		for indexOfLine, line := range lines {
			line = strings.TrimSpace(line)
			fmt.Fprintf(fRstOutput, rstListTablePrefixOfEachLine(indexOfTd, indexOfLine))
			fmt.Fprintf(fRstOutput, line)
			fmt.Fprintf(fRstOutput, "\n")
		}
	})
}

func htmlTableToRst(url, outputFilePath string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	fRstOutput, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}

	doc.Find("table").Each(func(_ int, table *goquery.Selection) {
		fmt.Fprintf(fRstOutput, ".. list-table:: \n\n")
		table.Find("tr").Each(func(_ int, tr *goquery.Selection) {
			processTr(tr, fRstOutput)
		})
		fmt.Fprintf(fRstOutput, "\n\n")
		fmt.Fprintf(fRstOutput, "|\n|\n\n")
	})
}
