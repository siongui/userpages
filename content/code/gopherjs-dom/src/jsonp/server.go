package main

import "net/http"
import "fmt"
import "encoding/json"

func jsonpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	if r.URL.Path != "/jsonp" {
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	b, _ := json.Marshal(map[string]string{
		"input": r.URL.Query().Get("input"),
	})
	fmt.Fprintf(w, "%s(%s);", r.URL.Query().Get("callback"), b)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/jsonp", jsonpHandler)
	http.ListenAndServe(":8000", nil)
}
