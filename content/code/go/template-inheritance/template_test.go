package gossg

import (
	"os"
	"testing"
)

type TemplateData struct {
	SITENAME string
	SITEURL  string
}

func TestTemplateToHtml(t *testing.T) {
	data := TemplateData{
		SITENAME: "Theory and Practice",
		SITEURL:  "https://siongui.github.io/",
	}

	tmpl, err := ParseTemplateDir("tmpl")
	if err != nil {
		t.Error(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", &data)
	if err != nil {
		t.Error(err)
	}
}
