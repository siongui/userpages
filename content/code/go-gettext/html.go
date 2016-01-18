package main

import (
	"github.com/chai2010/gettext-go/gettext"
	"html/template"
	"os"
)

func setup(locale string, domain string, dir string) {
	gettext.SetLocale(locale)
	gettext.Textdomain(domain)

	gettext.BindTextdomain(domain, dir, nil)
}

func changeLocale(locale string) {
	gettext.SetLocale(locale)
}

func translate(input string) string {
	return gettext.PGettext("", input)
}

const tmpl = `
<span>{{gettext "Home"}}</span>
<span>{{gettext "Canon"}}</span>
<span>{{gettext "About"}}</span>
<span>{{gettext "Setting"}}</span>
<span>{{gettext "Translation"}}</span>
`

func main() {
	setup("zh_TW", "messages", "locale")
	setup("vi_VN", "messages", "locale")
	funcMap := template.FuncMap{
		"gettext": translate,
	}

	t, _ := template.New("foo").Funcs(funcMap).Parse(tmpl)
	changeLocale("zh_TW")
	t.Execute(os.Stdout, nil)
	changeLocale("vi_VN")
	t.Execute(os.Stdout, nil)
}
