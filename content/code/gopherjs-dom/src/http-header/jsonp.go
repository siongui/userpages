package main

import "github.com/gopherjs/gopherjs/js"
import "honnef.co/go/js/dom"

func main() {
	js.Global.Set("mycallback", func(headers map[string]string) {
		println(headers)
		println(headers["Accept-Language"])
		println(headers["User-Agent"])
	})

	targetUrl := "http://ajaxhttpheaders.appspot.com/?callback=mycallback"

	d := dom.GetWindow().Document()
	ext := d.CreateElement("script")
	ext.SetAttribute("src", targetUrl)
	head := d.GetElementsByTagName("head")[0].(*dom.HTMLHeadElement)
	head.AppendChild(ext)
}
