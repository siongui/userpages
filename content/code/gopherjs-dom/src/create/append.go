package main

import "honnef.co/go/js/dom"

func main() {
	d := dom.GetWindow().Document()

	foo := d.GetElementByID("foo").(*dom.HTMLDivElement)
	foo.AddEventListener("click", false, func(event dom.Event) {
		div := d.CreateElement("div").(*dom.HTMLDivElement)
		div.Style().SetProperty("color", "red", "")
		div.SetTextContent("I am new div")
		foo.AppendChild(div)
	})
}
