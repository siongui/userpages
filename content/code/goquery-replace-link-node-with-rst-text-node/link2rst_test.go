package link2rst

import (
	"github.com/PuerkitoBio/goquery"
	"testing"
)

func TestHtmlAnchorNodeToRstTextNode(t *testing.T) {
	url := "http://nanda.online-dhamma.net/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	doc2 := HtmlAnchorNodeToRstTextNode(doc)
	print(doc2.Find("body").Text())
}
