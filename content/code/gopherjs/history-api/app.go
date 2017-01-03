// Google Search: html history api
// https://css-tricks.com/using-the-html5-history-api/
package main

import "github.com/gopherjs/gopherjs/js"

var history = js.Global.Get("history")

func main() {
	d := js.Global.Get("document")
	info := d.Call("getElementById", "info")
	nodeList := d.Call("querySelectorAll", ".person")
	length := nodeList.Get("length").Int()
	for i := 0; i < length; i++ {
		// get i-th element in nodelist
		elm := nodeList.Call("item", i)

		elm.Call("addEventListener", "click", func(event *js.Object) {
			event.Call("preventDefault")

			href := elm.Call("getAttribute", "href").String()
			content := elm.Get("dataset").Get("content").String()
			history.Call("pushState", content, nil, href)
			info.Set("innerHTML", content)
		})
	}

	js.Global.Call("addEventListener", "popstate", func(event *js.Object) {
		if event.Get("state") == nil {
			info.Set("innerHTML", "Entry Page")
		} else {
			info.Set("innerHTML", event.Get("state").String())
		}
	})
}
