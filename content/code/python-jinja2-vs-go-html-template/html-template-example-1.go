package main

import (
	"html/template"
	"os"
)

func main() {
	t, _ := template.New("foo").Parse(`Hello {{.}}!`)
	t.Execute(os.Stdout, "World")
}
