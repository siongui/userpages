package main

import "honnef.co/go/js/dom"

func main() {
	d := dom.GetWindow().Document()

	foo := d.GetElementByID("foo").(*dom.HTMLDivElement)
	foo.AddEventListener("click", false, func(event dom.Event) {
		foo.Style().SetProperty("display", "none", "")
	})
}
