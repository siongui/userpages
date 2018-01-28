package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type PostData struct {
	PostUrl   string
	Title     string
	TimeStamp string
	Author    string
	Summary   string
	Content   string
	Tags      string
}

func GetBlogspotTimeStamp(doc *goquery.Document) (string, error) {
	abbr := doc.Find("a.timestamp-link > abbr").First()
	t, ok := abbr.Attr("title")
	if ok {
		return t, nil
	}

	return "", errors.New("cannot find timestamp")
}

func GetBlogspotTitle(doc *goquery.Document) (string, error) {
	t := doc.Find("h3.post-title").First()
	return strings.TrimSpace(t.Text()), nil
}

func GetBlogspotContent(doc *goquery.Document) (string, error) {
	c := doc.Find("div.post-body").First()
	return c.Html()
}

func GetBlogspotUrl(doc *goquery.Document) (string, error) {
	meta := doc.Find("meta[property='og:url']").First()
	u, ok := meta.Attr("content")
	if ok {
		return u, nil
	}

	return "", errors.New("cannot find url")
}

func GetBlogspotSummary(doc *goquery.Document) (string, error) {
	meta := doc.Find("meta[property='og:description']").First()
	d, ok := meta.Attr("content")
	if ok {
		return d, nil
	}

	return "", errors.New("cannot find summary")
}

func GetBlogspotAuthor(doc *goquery.Document) (string, error) {
	a := doc.Find("span.post-author > span.fn").First()
	return a.Text(), nil
}

func GetBlogspotTags(doc *goquery.Document) (string, error) {
	s := doc.Find("span.post-labels > a")
	labels := ""
	s.Each(func(_ int, l *goquery.Selection) {
		if labels != "" {
			labels += ", "
		}
		labels += l.Text()
	})
	return labels, nil
}

func ParseBlogspotPost(doc *goquery.Document) (*PostData, error) {
	bs := PostData{}
	var err error

	bs.TimeStamp, err = GetBlogspotTimeStamp(doc)
	if err != nil {
		return &bs, err
	}

	bs.Title, err = GetBlogspotTitle(doc)
	if err != nil {
		return &bs, err
	}

	bs.Content, err = GetBlogspotContent(doc)
	if err != nil {
		return &bs, err
	}

	bs.PostUrl, err = GetBlogspotUrl(doc)
	if err != nil {
		return &bs, err
	}

	bs.Summary, err = GetBlogspotSummary(doc)
	if err != nil {
		return &bs, err
	}

	bs.Author, err = GetBlogspotAuthor(doc)
	if err != nil {
		return &bs, err
	}

	bs.Tags, err = GetBlogspotTags(doc)
	if err != nil {
		return &bs, err
	}

	return &bs, nil
}

func main() {
	//url := "https://oathbystyx.blogspot.tw/2018/01/descartes-rules-of-signs.html"
	url := "https://timrau.blogspot.com/2017/11/avoid-vim-overwriting-indention-settings.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	post, err := ParseBlogspotPost(doc)
	if err != nil {
		panic(err)
	}

	println(post.TimeStamp)
	println(post.Title)
	println(post.Content)
	println(post.PostUrl)
	println(post.Summary)
	println(post.Author)
	println(post.Tags)
}
