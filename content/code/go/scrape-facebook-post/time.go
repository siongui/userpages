package parsefb

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"time"
)

func ParseTimeStamp(utime string) (string, error) {
	i, err := strconv.ParseInt(utime, 10, 64)
	if err != nil {
		return "", err
	}
	t := time.Unix(i, 0)
	return t.Format(time.RFC3339), nil
}

func GetTimeStamp(postHtml string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(postHtml))
	if err != nil {
		return "", err
	}

	s := doc.Find("abbr._5ptz").First()
	utime, ok := s.Attr("data-utime")
	if ok {
		return ParseTimeStamp(utime)
	}

	return "", errors.New("cannot find timestamp")
}
