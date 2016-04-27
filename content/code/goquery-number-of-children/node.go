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

func NumberOfChild(s *goquery.Selection) int {
	return s.Contents().Length()
	//return s.Contents().Size()
}

func NumberOfElementChild(s *goquery.Selection) int {
	return s.Children().Length()
	//return s.Children().Size()
}

func main() {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}

	println(NumberOfChild(doc.Find("div")))
	println(NumberOfElementChild(doc.Find("div")))
}
