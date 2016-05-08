package html2rst

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

func isScriptElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "script"
}

func isImgElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "img"
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
	if isImgElement(n.FirstChild) {
		rstText := img2rst(n.FirstChild)
		href, ok := getAttribute(n, "href")
		if ok {
			rstText += "   :target: "
			rstText += href
			rstText += "\n"
		}
		return rstText
	}

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

func img2rst(n *html.Node) string {
	rstText := ".. image:: "

	src, ok := getAttribute(n, "src")
	if ok {
		rstText += src
		rstText += "\n"
	} else {
		rstText += "\n"
	}

	alt, ok := getAttribute(n, "alt")
	if ok {
		rstText += "   :alt: "
		rstText += alt
		rstText += "\n"
	}

	class, ok := getAttribute(n, "class")
	if ok {
		if class == "align-center" {
			rstText += "   :align: center\n"
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
	if isImgElement(n) {
		return img2rst(n)
	}
	if isUlElement(n) {
		return ul2rst(n)
	}
	if isScriptElement(n) {
		return ""
	}

	rstText := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		rstText += traverse(c)
	}

	return rstText
}

func HtmlToRst(r io.Reader) string {
	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}

	return traverse(doc)
}
