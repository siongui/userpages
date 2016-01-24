package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	req := js.Global.Get("XMLHttpRequest").New()
	req.Call("addEventListener", "load", func() {
		headers := req.Call("getAllResponseHeaders").String()
		println(headers)
	})
	req.Call("open", "GET", "/", true)
	req.Call("send")
}
