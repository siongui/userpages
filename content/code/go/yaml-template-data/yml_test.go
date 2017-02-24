package yamldata

import (
	"os"
	"testing"
	"text/template"
)

const tmpl = `
sitename: {{.SiteName}}
open graph url: {{.Og.Url}}
open graph locale: {{.Og.Locale}}
`

func TestServeFromYAML(t *testing.T) {
	data, err := YamlToStruct("tmpldata.yaml")
	if err != nil {
		t.Error(err)
	}

	tl, err := template.New("").Parse(tmpl)
	if err != nil {
		t.Error(err)
	}

	err = tl.Execute(os.Stdout, &data)
	if err != nil {
		t.Error(err)
	}
}
