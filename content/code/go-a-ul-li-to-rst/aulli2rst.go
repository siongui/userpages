package aulli2rst

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io"
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

func indentEachLine(s string) string {
	lines := StringToLines(s)
	var indentedLines []string
	for _, line := range lines {
		indentedLines = append(indentedLines, "  "+line)
	}
	return strings.Join(indentedLines, "\n")
}

func isAnchorElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func isUlElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "ul"
}

func isLiElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "li"
}

func isTextNode(n *html.Node) bool {
	return n.Type == html.TextNode
}

func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func textNode2rst(n *html.Node) string {
	text := strings.TrimSpace(n.Data)
	if text == "" {
		return "\n"
	}
	return n.Data
}

func a2rst(n *html.Node) string {
	text := strings.TrimSpace(n.FirstChild.Data)

	href, ok := getAttribute(n, "href")
	if ok {
		return "`" + text + " <" + href + ">`__"
	}

	return ""
}

func li2rst(n *html.Node) string {
	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isTextNode(c) {
			rstText += textNode2rst(c)
		}
		if isAnchorElement(c) {
			rstText += a2rst(c)
		}
		if isUlElement(c) {
			rstText += "\n"
			rstText += indentEachLine(ul2rst(c))
		}
	}

	return "- " + rstText + "\n"
}

func ul2rst(n *html.Node) string {
	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isLiElement(c) {
			rstText += li2rst(c)
		}
	}

	return rstText
}

func traverse(n *html.Node) string {
	if isTextNode(n) {
		return textNode2rst(n)
	}
	if isAnchorElement(n) {
		return a2rst(n)
	}
	if isUlElement(n) {
		return ul2rst(n)
	}

	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		rstText += traverse(c)
	}

	return rstText
}

func HtmlAUlLiToRst(r io.Reader) string {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}
