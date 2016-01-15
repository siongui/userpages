package main

import "honnef.co/go/js/dom"

func main() {
	d := dom.GetWindow().Document()

	foo := d.GetElementByID("foo").(*dom.HTMLButtonElement)
	foo.AddEventListener("click", false, func(event dom.Event) {
		a := d.GetElementByID("audio").(*dom.HTMLAudioElement)
		// already specify source in audio tag, no need to set src
		//a.SetAttribute("src", "Wat_Metta_Buddha_Qualities.mp3")
		a.Play()
	})
}
