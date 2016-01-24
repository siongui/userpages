package main

import "honnef.co/go/js/xhr"

func main() {
	req := xhr.NewRequest("GET", "/")
	err := req.Send(nil)
	if err != nil { println(err) }
	headers := req.ResponseHeaders()
	println(headers)
}
