package main

import "honnef.co/go/js/dom"

func main() {
	d := dom.GetWindow().Document()

	foo := d.GetElementByID("foo").(*dom.HTMLButtonElement)
	foo.AddEventListener("click", false, func(event dom.Event) {
		a := d.GetElementByID("audio").(*dom.HTMLAudioElement)
		if a.Paused {
			a.Play()
			foo.SetTextContent("Click Me to Pause Sound")
		} else {
			a.Pause()
			foo.SetTextContent("Click Me to Play Sound")
		}
	})
}
