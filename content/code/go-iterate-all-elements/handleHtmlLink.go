package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func isAnchorElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func isTextNode(n *html.Node) bool {
	return n.Type == html.TextNode
}

func isHasOnlyOneChild(n *html.Node) bool {
	return n.FirstChild == n.LastChild
}

func getAttribute(n *html.Node, key string) (string, error) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, nil
		}
	}
	return "", errors.New(key + " not exist in attribute!")
}

func printRstLink(n *html.Node) {
	if !isHasOnlyOneChild(n) {
		fmt.Fprintf(os.Stderr, "Child number of anchor is not 1\n")
		return
	}

	if !isTextNode(n.FirstChild) {
		fmt.Fprintf(os.Stderr, "Child of anchor is not TextNode\n")
		return
	}

	text := strings.TrimSpace(n.FirstChild.Data)

	href, err := getAttribute(n, "href")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	rstLink := "`" + text + " <" + href + ">`__"
	fmt.Println(rstLink)
}
