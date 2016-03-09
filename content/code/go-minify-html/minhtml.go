package minhtml

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"
)

func RemoveHTMLComments(content []byte) []byte {
	// https://www.google.com/search?q=regex+html+comments
	// http://stackoverflow.com/a/1084759
	htmlcmt := regexp.MustCompile(`<!--[^>]*-->`)
	return htmlcmt.ReplaceAll(content, []byte(""))
}

func MinifyHTML(html []byte) string {
	// read line by line
	minifiedHTML := ""
	scanner := bufio.NewScanner(bytes.NewReader(RemoveHTMLComments(html)))
	for scanner.Scan() {
		// all leading and trailing white space of each line are removed
		lineTrimmed := strings.TrimSpace(scanner.Text())
		minifiedHTML += lineTrimmed
		if len(lineTrimmed) > 0 {
			// in case of following trimmed line:
			// <div id="foo"
			minifiedHTML += " "
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return minifiedHTML
}
