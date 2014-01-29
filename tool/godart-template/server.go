package main

import (
	"fmt"
	"net/http"
	"html/template"
)

const tmplFilePath = "view/index.html"
var tmpl, _ = template.ParseFiles(tmplFilePath)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL.Path)
	tmpl.Execute(w, nil)
}


func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}
