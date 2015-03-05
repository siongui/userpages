package main

import (
	"html/template"
	"os"
)

const tmpl = `
{{range $link := .}}
<a href="{{$link.Href}}">{{$link.Name}}</a>
{{end}}
`

type Link struct {
	Name	string
	Href	string
}

func main() {
	// arrays
	var la [2]Link
	la[0] = Link{"Google", "https://www.google.com/"}
	la[1] = Link{"Facebook", "https://www.facebook.com/"}
	t, _ := template.New("foo").Parse(tmpl)
	t.Execute(os.Stdout, la)

	// slices
	var ls []Link
	ls = append(ls, Link{"Google", "https://www.google.com/"})
	ls = append(ls, Link{"Facebook", "https://www.facebook.com/"})
	t.Execute(os.Stdout, ls)
}
