package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexHtml)
}

func jsonpHandler(w http.ResponseWriter, r *http.Request) {
	callbackName := r.URL.Query().Get("callback")
	if callbackName == "" {
		fmt.Fprintf(w, "Please give callback name in query string")
		return
	}

	b, err := json.Marshal(r.Header)
	if err != nil {
		fmt.Fprintf(w, "json encode error")
		return
	}

	w.Header().Set("Content-Type", "application/javascript")
	// return `callbackName(jsonString);`
	fmt.Fprintf(w, "%s(%s);", callbackName, b)
}

func main() {
	http.HandleFunc("/jsonp", jsonpHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
