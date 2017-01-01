package main

import (
	"flag"
	"net/http"
)

func main() {
	dir := flag.String("dir", "default value", "Directory of GopherJS Demo")
	flag.Parse()
	http.ListenAndServe(":8000", http.FileServer(http.Dir(*dir)))
}
