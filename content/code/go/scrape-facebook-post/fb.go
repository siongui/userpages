package parsefb

import (
	"github.com/PuerkitoBio/goquery"
)

func Parse(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	s := doc.Find("div.hidden_elem > code").First()
	cmt, err := s.Html()
	if err != nil {
		return "", err
	}
	postHtml := cmt[5 : len(cmt)-4]
	return postHtml, nil
}
