package main

import "honnef.co/go/js/dom"

func inputKeyUp(event dom.Event) {
	input := event.Target().(*dom.HTMLInputElement)

	span := dom.GetWindow().Document().GetElementByID("foo2")
	span.SetInnerHTML(input.Value)
}

func main() {
	d := dom.GetWindow().Document()
	p := d.GetElementByID("foo").(*dom.HTMLInputElement)

	p.Focus()
	p.AddEventListener("keyup", false, inputKeyUp)
}
