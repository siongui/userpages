package mincss

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"
)

func RemoveCStyleComments(content []byte) []byte {
	// http://blog.ostermiller.org/find-comment
	ccmt := regexp.MustCompile(`/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/`)
	return ccmt.ReplaceAll(content, []byte(""))
}

func RemoveCppStyleComments(content []byte) []byte {
	cppcmt := regexp.MustCompile(`//.*`)
	return cppcmt.ReplaceAll(content, []byte(""))
}

func concatenateCSS(cssPathes []string) []byte {
	var cssAll []byte
	for _, cssPath := range cssPathes {
		println("concatenating " + cssPath + " ...")
		b, err := ioutil.ReadFile(cssPath)
		if err != nil {
			panic(err)
		}
		cssAll = append(cssAll, b...)
	}
	return cssAll
}

func MinifyCSS(cssPathes []string) string {
	cssAll := concatenateCSS(cssPathes)
	cssAllNoComments := RemoveCppStyleComments(RemoveCStyleComments(cssAll))

	// read line by line
	minifiedCss := ""
	scanner := bufio.NewScanner(bytes.NewReader(cssAllNoComments))
	for scanner.Scan() {
		// all leading and trailing white space of each line are removed
		minifiedCss += strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return minifiedCss
}
