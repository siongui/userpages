package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	demo := js.Global.Get("document").Call("getElementById", "demo")
	info := js.Global.Get("document").Call("getElementById", "info")

	demo.Call("addEventListener", "mouseenter", func(event *js.Object) {
		info.Set("innerHTML", "mouse entered")
	}, false)

	demo.Call("addEventListener", "mouseleave", func(event *js.Object) {
		info.Set("innerHTML", "mouse leaved")
	}, false)
}
