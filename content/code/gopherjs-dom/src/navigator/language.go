package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	println(js.Global.Get("navigator").Get("language"))
	println(js.Global.Get("navigator").Get("languages"))
	println(js.Global.Get("navigator").Get("onLine"))
}
