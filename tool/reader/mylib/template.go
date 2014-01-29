package mylib

import (
	"html/template"
)

func GetTemplate(filepath string) (* template.Template) {
	tmpl, err := template.ParseFiles(filepath)
	if err != nil { panic(err) }
	return tmpl
}
