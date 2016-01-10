package main

import "honnef.co/go/js/dom"

func main() {
	d := dom.GetWindow().Document()
	h := d.GetElementByID("foo")

	h.AddEventListener("click", false, func(event dom.Event) {
		event.PreventDefault()
		h.SetInnerHTML("I am Clicked")
		println("This message is printed in browser console")
	})
}
