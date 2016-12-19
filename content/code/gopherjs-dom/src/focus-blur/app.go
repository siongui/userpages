package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	i1 := js.Global.Get("document").Call("getElementById", "i1")
	i2 := js.Global.Get("document").Call("getElementById", "i2")
	btn := js.Global.Get("document").Call("getElementById", "btn")

	btn.Set("onclick", func(e *js.Object) {
		i1.Call("blur")
		i2.Call("focus")
	})
}
