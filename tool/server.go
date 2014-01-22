package main

import (
	"fmt"
	"net/http"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL.Path)
	fmt.Fprint(w, "hello world")
}


func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
