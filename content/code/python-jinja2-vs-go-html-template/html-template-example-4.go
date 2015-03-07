package main

import (
	"html/template"
	"os"
)

const tmpl = `
{{range $name, $href := .}}
<a href="{{$href}}">{{$name}}</a>
{{end}}
`

func main() {
	// map
	var m = map[string]string{
	    "Google": "https://www.google.com/",
	    "Facebook": "https://www.facebook.com/",
	}
	t, _ := template.New("foo").Parse(tmpl)
	t.Execute(os.Stdout, m)
}
