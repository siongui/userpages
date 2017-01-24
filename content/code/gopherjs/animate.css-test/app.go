package main

import "github.com/gopherjs/gopherjs/js"

func GetSelectValue(s *js.Object) string {
	return s.Get("options").Call("item", s.Get("selectedIndex").Int()).Get("value").String()
}

func RemoveAllClass(elm *js.Object) {
	elm.Set("className", "")
}

func main() {
	d := js.Global.Get("document")
	text := d.Call("getElementById", "animationSandbox")
	sel := d.Call("getElementById", "animate-select")
	btn := d.Call("getElementById", "animate-btn")

	btn.Call("addEventListener", "click", func(event *js.Object) {
		action := GetSelectValue(sel)
		RemoveAllClass(text)
		text.Set("className", "animated "+action)
	})

	text.Call("addEventListener", "animationend", func(event *js.Object) {
		RemoveAllClass(text)
	})
}
