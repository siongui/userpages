package main

import (
	"fmt"
	"flag"
	"net/http"
	"encoding/json"
	"./mylib"
)


var (
	sites = mylib.GetSites("sqlite3/sites.db")
	indexTmpl = mylib.GetTemplate("view/index.html")
	rssItemsTmpl = mylib.GetTemplate("view/rss_items.html")
	isProductonServer = flag.Bool("production", false, "localhost:8080")
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
	// FIXME: do not use panic here
	if err != nil { panic(err) }
	fmt.Println(s.Text)

	items := mylib.GetItemsFromDB(s.XmlUrl)
	rssItemsTmpl.Execute(w, items)
}

func saveLinkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { return }

	var l mylib.LinkInfo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&l)
	// FIXME: do not use panic here
	if err != nil { panic(err) }

	mylib.SaveLinkAsJson(l)
}

func main() {
	mylib.Poll(sites)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/site/", siteHandler)
	http.HandleFunc("/savelink/", saveLinkHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	flag.Parse()
	if *isProductonServer {
		http.ListenAndServe(":8080", nil)
	} else {
		http.ListenAndServe(":8000", nil)
	}
}
