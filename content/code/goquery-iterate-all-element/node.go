package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const html = `<html>
  <head>
    <title>traverse</title>
  </head>
  <body>
    <div>
      Hello
      <span>World</span>
      <!-- Goquery -->
    </div>
  </body>
</html>`

func main() {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}

	doc.Find("*").Each(func(_ int, node *goquery.Selection) {
		println(node.Text())
	})
}
