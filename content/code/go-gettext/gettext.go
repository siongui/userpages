package mylib

import "github.com/chai2010/gettext-go/gettext"

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
