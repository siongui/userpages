package main

import (
	"github.com/gopherjs/gopherjs/js"
	"time"
)

func main() {
	info := js.Global.Get("document").Call("getElementById", "info")
	button := js.Global.Get("document").Call("getElementById", "btn")

	button.Set("onclick", func(e *js.Object) {
		// time.AfterFunc is Go's setTimeout
		time.AfterFunc(3*time.Second, func() {
			info.Set("innerHTML", "time is up")
		})
	})
}
