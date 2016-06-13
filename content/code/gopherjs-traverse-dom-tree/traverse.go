package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func traverse(elm *js.Object) {
	nodeType := elm.Get("nodeType").Int()

	if nodeType == 1 || nodeType == 9 {
		// element node or document node
		childNodesList := elm.Get("childNodes")
		length := childNodesList.Get("length").Int()
		for i := 0; i < length; i++ {
			// recursively call to traverse
			traverse(childNodesList.Call("item", i))
		}
		return
	}

	if nodeType == 3 {
		// text node
		s := elm.Get("nodeValue").String()
		println(s)
		return
	}
}

func main() {
	doc := js.Global.Get("document")
	traverse(doc)
}
