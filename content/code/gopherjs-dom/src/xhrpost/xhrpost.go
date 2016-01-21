package main

import "net/http"
import "strings"

func main() {
	jsonStr := `{"title":"test","data":"none"}`
	resp, err := http.Post("/post", "application/json", strings.NewReader(jsonStr))
	if err != nil {
		// handle error
		println(err)
	}
	defer resp.Body.Close()
}
