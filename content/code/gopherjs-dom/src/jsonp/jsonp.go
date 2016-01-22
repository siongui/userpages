package main

import "github.com/gopherjs/gopherjs/js"
import "honnef.co/go/js/dom"
import "net/url"

func main() {
	js.Global.Set("mycallback", func(jsonData map[string]string) {
		println(jsonData["input"])
	})

	targetUrl := "/jsonp?callback=mycallback" +
		"&input=" + url.QueryEscape("你好")

	d := dom.GetWindow().Document()
	ext := d.CreateElement("script")
	ext.SetAttribute("src", targetUrl)
	head := d.GetElementsByTagName("head")[0].(*dom.HTMLHeadElement)
	head.AppendChild(ext)
}
