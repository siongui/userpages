package main

import (
	"encoding/json"
	"net/http"
)

type TestData struct {
	Title string
	Data  string
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	if r.URL.Path != "/post" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	d := TestData{}

	// should handle error here
	decoder.Decode(&d)
	print(d.Title)
	print(d.Data)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":8000", nil)
}
