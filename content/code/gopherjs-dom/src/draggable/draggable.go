package main

import "github.com/gopherjs/gopherjs/js"
import "honnef.co/go/js/dom"
import "strconv"

func draggable(elm *dom.HTMLDivElement) {
	var startX, startY, initialMouseX, initialMouseY int
	var fmove, fup func(*js.Object)
	elm.Style().SetProperty("position", "absolute", "")
	d := dom.GetWindow().Document()

	mousemove := func(event dom.Event) {
		event.PreventDefault()
		e := event.(*dom.MouseEvent)
		dx := e.ClientX - initialMouseX
		dy := e.ClientY - initialMouseY
		elm.Style().SetProperty("top", strconv.Itoa(startY+dy)+"px", "")
		elm.Style().SetProperty("left", strconv.Itoa(startX+dx)+"px", "")
	}

	mouseup := func(event dom.Event) {
		d.RemoveEventListener("mousemove", false, fmove)
		d.RemoveEventListener("mouseup", false, fup)
	}

	elm.AddEventListener("mousedown", false, func(event dom.Event) {
		event.PreventDefault()
		e := event.(*dom.MouseEvent)
		startX = int(elm.OffsetLeft())
		startY = int(elm.OffsetTop())
		initialMouseX = e.ClientX
		initialMouseY = e.ClientY
		fmove = d.AddEventListener("mousemove", false, mousemove)
		fup = d.AddEventListener("mouseup", false, mouseup)
	})
}

func main() {
	draggable(dom.GetWindow().Document().GetElementByID("dragMe").(*dom.HTMLDivElement))
}
