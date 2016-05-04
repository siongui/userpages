package link2rst

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"text/template"
)

const rstLink = "`{{.Text}} &lt;{{.Href}}&gt;`_"

type htmlLink struct {
	Text string
	Href string
}

func HtmlAnchorNodeToRstTextNode(doc *goquery.Document) *goquery.Document {
	tmpl := template.Must(template.New("link2rst").Parse(rstLink))

	doc.Find("a").Each(func(_ int, link *goquery.Selection) {
		text := strings.TrimSpace(link.Text())
		href, ok := link.Attr("href")
		if ok {
			var rstBuf bytes.Buffer
			err := tmpl.Execute(&rstBuf, &htmlLink{text, href})
			if err != nil {
				panic(err)
			}
			link.ReplaceWithHtml(rstBuf.String())
		}
	})

	return doc
}
