package main

import "github.com/gopherjs/gopherjs/js"

func handleInputKeyUp(event *js.Object) {
	if keycode := event.Get("keyCode").Int(); keycode == 13 {
		// user press enter key
		print("enter key")
	}
}

func main() {
	foo := js.Global.Get("document").Call("getElementById", "foo")
	foo.Call("addEventListener", "keyup", handleInputKeyUp, false)
}
