package main

import (
	"github.com/PuerkitoBio/goquery"
	"os"
	"strings"
	"text/template"
)

const rstLink = "`{{.Text}} <{{.Href}}>`_\n"

type htmlLink struct {
	Text string
	Href string
}

func main() {
	url := "https://siongui.github.io/2016/04/09/js-copy-to-clipboard/"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.New("link2rst").Parse(rstLink))
	doc.Find("a").Each(func(_ int, link *goquery.Selection) {
		text := strings.TrimSpace(link.Text())
		href, ok := link.Attr("href")
		if ok {
			tmpl.Execute(os.Stdout, &htmlLink{text, href})
		}
	})
}
