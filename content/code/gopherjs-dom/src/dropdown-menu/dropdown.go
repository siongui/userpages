package main

import "honnef.co/go/js/dom"

func checkParent(target, elm dom.Element) bool {
	for target.ParentElement() != nil {
		if target.IsEqualNode(elm) {
			return true
		}
		target = target.ParentElement()
	}
	return false
}

func main() {
	d := dom.GetWindow().Document()

	toggle := d.GetElementByID("menu-dropdown").(*dom.HTMLDivElement)
	menu := d.GetElementByID("menuDiv-dropdown").(*dom.HTMLDivElement)

	d.AddEventListener("click", false, func(event dom.Event) {
		if !checkParent(event.Target(), menu) {
			// click NOT on the menu
			if checkParent(event.Target(), toggle) {
				// click on the link
				menu.Class().Toggle("invisible")
			} else {
				// click both outside link and outside menu, hide menu
				menu.Class().Add("invisible")
			}
		}
	})
}
