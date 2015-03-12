// http://stackoverflow.com/questions/17843311/template-and-custom-function-panic-function-not-defined
package main

import (
	"html/template"
	"os"
)

const tmpl = `
<span>hello {{gettext .}}</span>
<span>hello {{. | gettext}}</span>
`

var funcMap = template.FuncMap{
	"gettext": gettext,
}

func gettext(s string) string {
	if s == "world" {
		return "世界"
	}
	return s
}

func main() {
	t, _ := template.New("foo").Funcs(funcMap).Parse(tmpl)
	s := "world"
	t.Execute(os.Stdout, s)
}
