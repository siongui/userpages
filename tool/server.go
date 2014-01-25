package main

import (
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"
	"io/ioutil"
)


type UrlData struct {
	Url	string
	Tag	string
}


const tmplFilePath = "view/index.html"
var tmpl, _ = template.ParseFiles(tmplFilePath)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL.Path)
	tmpl.Execute(w, nil)
}

func urlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL.Path)
	if r.Method != "POST" { return }

	var m UrlData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&m)
	if err != nil { panic(err) }
	fmt.Println("Tag: ", m.Tag, ", URL: ", m.Url)

	resp, err := http.Get(m.Url)
	if err != nil { panic(err) }
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }
	fmt.Fprint(w, string(body))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/url/", urlHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}
