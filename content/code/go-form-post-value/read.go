package main

import (
	"fmt"
	"net/http"
)

var indexHtml = `<!doctype html>
<html>
<head><title>Get Form POST Value</title></head>
<body>
<form action="/post" method="post">
  <input name="myValue">
  <button>Send</button>
</form>
</body>
</html>`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, indexHtml)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.PostFormValue("myValue"))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/post", postHandler)
	http.ListenAndServe(":8000", nil)
}
