package main

import "honnef.co/go/js/dom"

func main() {
	d := dom.GetWindow().Document()

	div1 := d.GetElementByID("foo").(*dom.HTMLDivElement)
	print(div1.Dataset()["myDemoValue"])

	div2 := d.GetElementByID("foo2").(*dom.HTMLDivElement)
	print(div2.Dataset()["hello"])
}
