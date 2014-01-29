package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"./mylib"
)


var (
	sites = mylib.GetSites("sqlite3/sites.db")
	indexTmpl = mylib.GetTemplate("view/index.html")
	rssItemsTmpl = mylib.GetTemplate("view/rss_items.html")
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL.Path)
	err := indexTmpl.Execute(w, sites)
	if err != nil { panic(err) }
}

func siteHandler(w http.ResponseWriter, r *http.Request) {
	// @see https://groups.google.com/d/topic/golang-nuts/fv94aZWK3Go
	// @see http://stackoverflow.com/questions/15672556/handling-json-post-request-in-go
	if r.Method != "POST" { return }

	var s mylib.OpmlOutline
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&s)
	if err != nil { panic(err) }
	fmt.Println(s.Text)

	v := mylib.GetSeed(s.XmlUrl)
	rssItemsTmpl.Execute(w, v)
}

func main() {
	mylib.Poll(sites)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/site/", siteHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}
