package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var onWordMouseOver = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Get("style").Set("color", "red")
	return nil
})

var onWordMouseOut = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
	this.Get("style").Set("color", "")
	return nil
})

func main() {
	spans := js.Global.Get("document").Call("getElementById", "container").Call("querySelectorAll", "span")
	// access individual span
	length := spans.Get("length").Int()
	for i := 0; i < length; i++ {
		span := spans.Call("item", i)
		span.Set("onmouseover", onWordMouseOver)
		span.Set("onmouseout", onWordMouseOut)
	}
}
