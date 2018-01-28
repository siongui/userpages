package parsefb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ProfileLink struct {
	Name string
	Url  string
}

func GetProfileLink(postHtml string) (*ProfileLink, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(postHtml))
	if err != nil {
		return nil, err
	}

	s := doc.Find("a.profileLink").First()
	if s.Length() == 0 {
		s = doc.Find("span.fwb.fcg > a").First()
	}

	pl := ProfileLink{}

	pl.Name = s.Text()
	if pl.Name == "" {
		return nil, errors.New("cannot find name of profile link")
	}

	url, ok := s.Attr("href")
	if !ok {
		return nil, errors.New("cannot find url of profile link")
	}
	pl.Url = url

	return &pl, nil
}
