// Google search keyword: golang http template
// https://golang.org/doc/articles/wiki/
// http://golang.org/pkg/html/template/

// Google search keyword: go serve static file
// http://stackoverflow.com/questions/17690230/go-web-page-static-file-serving
// http://stackoverflow.com/questions/25945538/go-golang-to-serve-a-specific-html-file

// Google search keyword: golang http minetype set
// https://golang.org/src/net/http/pprof/pprof.go#L73
// http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang

// Handling JSON Post Request in Go
// http://stackoverflow.com/questions/15672556/handling-json-post-request-in-go
package main

import (
	"html/template"
	"net/http"
	"encoding/json"
	"fmt"
)

type dataFromClient struct {
	Title string // cannot use title (lower case will case json decode failure)
	Url string // cannot use url (lower case will case json decode failure)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/app.dart" {
		w.Header().Set("Content-Type", "application/dart; charset=utf-8")
		http.ServeFile(w, r, "./app.dart")
		return
	}
	if r.URL.Path == "/app.js" {
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		http.ServeFile(w, r, "./app.js")
		return
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, r.URL.Path)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { return }
	if r.URL.Path != "/post" { return }

	decoder := json.NewDecoder(r.Body)
	var d dataFromClient

	// should handle error here
	decoder.Decode(&d)
	fmt.Fprintf(w, "<a href=\"%s\">%s</a>", d.Url, d.Title)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":8000", nil)
}
