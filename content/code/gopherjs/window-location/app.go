package main

import "github.com/gopherjs/gopherjs/js"

var location = js.Global.Get("location")

func main() {
	str := location.Get("href").String()
	str += "<br>"
	str += location.Get("host").String()
	str += "<br>"
	str += location.Get("hostname").String()
	str += "<br>"
	str += location.Get("pathname").String()
	str += "<br>"
	str += location.Get("protocol").String()
	str += "<br>"
	str += location.Get("origin").String()
	str += "<br>"
	str += location.Get("port").String()
	str += "<br>"
	str += location.Get("search").String()
	str += "<br>"

	info := js.Global.Get("document").Call("getElementById", "info")
	info.Set("innerHTML", str)
}
