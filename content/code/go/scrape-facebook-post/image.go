package parsefb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetImageUrl(postHtml string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(postHtml))
	if err != nil {
		return "", err
	}

	s := doc.Find("img.scaledImageFitHeight").First()
	if s.Length() == 0 {
		s = doc.Find("img.scaledImageFitWidth").First()
	}

	url, ok := s.Attr("src")
	if !ok {
		return "", errors.New("cannot find image url")
	}

	return url, nil
}
