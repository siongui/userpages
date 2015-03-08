package main

import (
	"html/template"
	"os"
)

func main() {
	// cannot put base-go.html before index-go.html
	// the following will give empty output
	//t, _ := template.ParseFiles("base-go.html", "index-go.html")
	t, _ := template.ParseFiles("index-go.html", "base-go.html")
	name := "world"
	t.Execute(os.Stdout, name)
}
