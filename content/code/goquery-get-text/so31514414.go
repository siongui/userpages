package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const html = `
<div class="container">
    <div class="row">
      <div class="col-lg-8">
        <p align="justify"><b>Name</b>Priyaka</p>
        <p align="justify"><b>Surname</b>Patil</p>
        <p align="justify"><b>Adress</b><br>India,Kolhapur</p>
        <p align="justify"><b>Hobbies&nbsp;</b><br>Playing</p>
        <p align="justify"><b>Eduction</b><br>12th</p>
        <p align="justify"><b>School</b><br>New Highschool</p>
       </div>
    </div>
</div>
`

func main() {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}

	doc.Find(".container").Find("[align=\"justify\"]").Each(func(_ int, s *goquery.Selection) {
		prefix := s.Find("b").Text()
		result := strings.TrimPrefix(s.Text(), prefix)
		println(result)
	})
}
