package taobaoitem2rst

import (
	"golang.org/x/net/html"
	"io"
)

type ItemInfo struct {
	Title  string
	URL    string
	ImgURL string
}

func GetAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func isLinkElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "link"
}

func isImgElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "img"
}

func isMetaElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "meta"
}

func traverse(n *html.Node, ii *ItemInfo) {
	if isTitleElement(n) {
		ii.Title = n.FirstChild.Data
	}
	if isLinkElement(n) {
		rel, ok := GetAttribute(n, "rel")
		if ok && rel == "canonical" {
			ii.URL, _ = GetAttribute(n, "href")
		}
	}
	if isImgElement(n) {
		// item.taobao.com
		id, ok := GetAttribute(n, "id")
		if ok && id == "J_ImgBooth" {
			ii.ImgURL, _ = GetAttribute(n, "src")
		}
	}
	if isMetaElement(n) {
		// world.taobao.com
		property, ok := GetAttribute(n, "property")
		if ok && property == "og:image" {
			ii.ImgURL, _ = GetAttribute(n, "content")
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, ii)
	}
}

func getTaobaoItemInfo(r io.Reader) ItemInfo {
	ii := ItemInfo{}

	doc, err := html.Parse(r)
	if err != nil {
		panic("Fail to parse html")
	}
	traverse(doc, &ii)

	return ii
}
