package main

import (
	"html/template"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index-go.html")
	name := "World"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
