package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
)

type PostData struct {
	PostUrl   string
	Title     string
	TimeStamp string
	Author    string
	Summary   string
	Content   string
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
	t := doc.Find("h3.post-title > a").First()
	return t.Text(), nil
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

	return &bs, nil
}

func main() {
	url := "https://oathbystyx.blogspot.tw/2018/01/descartes-rules-of-signs.html"
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
}
