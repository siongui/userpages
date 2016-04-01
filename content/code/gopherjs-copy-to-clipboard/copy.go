package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var document = js.Global.Get("document")

func copyTextarea(id string) {
	ta := document.Call("getElementById", id)
	ta.Call("select")

	isSuccessful := document.Call("execCommand", "copy").Bool()
	if isSuccessful {
		js.Global.Call("alert", "Copy Succeed")
	} else {
		js.Global.Call("alert", "Copy Fail")
	}
}

func main() {
	btn := document.Call("getElementById", "copybutton")
	btn.Call("addEventListener", "click", func(event *js.Object) {
		copyTextarea("demotextarea")
	}, false)
}
