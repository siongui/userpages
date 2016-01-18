package main

import "honnef.co/go/js/dom"


func main() {
	d := dom.GetWindow().Document()
	foo := d.GetElementByID("foo").(*dom.HTMLDivElement)
	foo1 := d.GetElementByID("foo1").(*dom.HTMLSpanElement)
	foo2 := d.GetElementByID("foo2").(*dom.HTMLSpanElement)
	foo3 := d.GetElementByID("foo3").(*dom.HTMLDivElement)
	foo.AddEventListener("click", false, func(event dom.Event) {
		foo1.Class().Toggle("invisible")
		foo2.Class().Toggle("invisible")
		foo3.Class().Toggle("invisible")
	})
}
