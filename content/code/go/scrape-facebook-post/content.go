package parsefb

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetContent(postHtml string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(postHtml))
	if err != nil {
		return "", err
	}

	s := doc.Find("div._5pbx.userContent").First()
	if s.Length() == 0 {
		return "no content", nil
	}

	content, err := s.Html()
	if err != nil {
		return "", err
	}

	return content, nil
}
