package main

import (
	"flag"
	"net/http"
)

func main() {
	dir := flag.String("dir", "src/focus-blur", "Directory of Demo")
	flag.Parse()
	http.ListenAndServe(":8000", http.FileServer(http.Dir(*dir)))
}
