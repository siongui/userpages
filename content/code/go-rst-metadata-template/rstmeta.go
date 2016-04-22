package rstmeta

import (
	"golang.org/x/text/width"
	"io"
	"strings"
	"text/template"
)

var rstMetadataTmpl = `{{.Title}}
{{rstTitle .Title}}

:date: {{.Date}}
:tags: {{.Tags}}
:category: {{.Category}}
:summary: {{.Summary}}


`

type RstMetadata struct {
	Title    string
	Date     string
	Tags     string
	Category string
	Summary  string
}

func GetWidthUTF8String(s string) int {
	size := 0
	for _, runeValue := range s {
		p := width.LookupRune(runeValue)
		if p.Kind() == width.EastAsianWide {
			size += 2
			continue
		}
		if p.Kind() == width.EastAsianNarrow {
			size += 1
			continue
		}
		panic("cannot determine!")
	}
	return size
}

func rstTitle(title string) string {
	width := GetWidthUTF8String(title)
	return strings.Repeat("=", width)
}

func writeRstMetadata(wr io.Writer, title, date, tags, category, summary string) {
	meta := RstMetadata{
		Title:    title,
		Date:     date,
		Tags:     tags,
		Category: category,
		Summary:  summary,
	}

	var funcMap = template.FuncMap{
		"rstTitle": rstTitle,
	}
	tmpl := template.Must(template.New("rstmeta").Funcs(funcMap).Parse(rstMetadataTmpl))
	tmpl.Execute(wr, meta)
}
