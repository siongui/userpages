package main

import "honnef.co/go/js/dom"
import "strconv"

func main() {
	w := dom.GetWindow()
	s := w.Document().GetElementByID("foo")

	w.AddEventListener("keydown", false, func(event dom.Event) {
		ke := event.(*dom.KeyboardEvent)
		s.SetInnerHTML(strconv.Itoa(ke.KeyCode))
	})
}
