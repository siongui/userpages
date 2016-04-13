package table2rst

import (
	"golang.org/x/net/html"
	"strings"
)

type ElementType int

const (
	TextNode ElementType = iota
	tableElementNode
	tbodyElementNode
	trElementNode
	tdElementNode
	NoNeedToKnow
)

func getElementType(n *html.Node) ElementType {
	if n.Type == html.TextNode {
		return TextNode
	}
	if n.Type == html.ElementNode && n.Data == "td" {
		return tdElementNode
	}
	if n.Type == html.ElementNode && n.Data == "tr" {
		return trElementNode
	}
	if n.Type == html.ElementNode && n.Data == "tbody" {
		return tbodyElementNode
	}
	if n.Type == html.ElementNode && n.Data == "table" {
		return tableElementNode
	}
	return NoNeedToKnow
}

func getTextNodeRst(text *html.Node) string {
	return strings.TrimSpace(text.Data)
}

func getTdRst(td *html.Node) string {
	s := ""
	for c := td.FirstChild; c != nil; c = c.NextSibling {
		if getElementType(c) == TextNode {
			s += (getTextNodeRst(c) + "\n")
			continue
		}
		panic("cannot convert this td")
	}
	return s
}

func getTrRst(tr *html.Node) string {
	s := ""
	isFirstTd := true
	for c := tr.FirstChild; c != nil; c = c.NextSibling {
		if getElementType(c) == tdElementNode {
			if isFirstTd {
				s += ("  * - " + getTdRst(c))
				isFirstTd = false
			} else {
				s += ("    - " + getTdRst(c))
			}
			continue
		}
		if getElementType(c) == TextNode {
			s += getTextNodeRst(c)
			continue
		}
		panic("cannot convert this tr")
	}
	return s
}

func getTbodyRst(tbody *html.Node) string {
	s := ""
	for c := tbody.FirstChild; c != nil; c = c.NextSibling {
		if getElementType(c) == trElementNode {
			s += getTrRst(c)
			continue
		}
		if getElementType(c) == TextNode {
			s += getTextNodeRst(c)
			continue
		}
		panic("cannot convert this tbody")
	}
	return s
}

func getTableRst(table *html.Node) string {
	s := ".. list-table::\n\n"
	for c := table.FirstChild; c != nil; c = c.NextSibling {
		if getElementType(c) == tbodyElementNode {
			s += getTbodyRst(c)
			continue
		}
		if getElementType(c) == TextNode {
			s += getTextNodeRst(c)
			continue
		}
		panic("cannot convert this table")
	}
	return s
}

func traverse(n *html.Node) string {
	s := ""

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if getElementType(c) == tableElementNode {
			s += getTableRst(c)
			continue
		} else {
			s += traverse(c)
		}
	}

	return s
}

func HtmlTableToRstListTable(doc *html.Node) string {
	return traverse(doc)
}
