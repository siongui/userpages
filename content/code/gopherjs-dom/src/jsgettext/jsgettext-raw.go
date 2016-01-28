package main

import "github.com/gopherjs/gopherjs/js"
import "encoding/json"
import "strings"

const jsondata = `{"vi_VN":{"About":"Giới thiệu","Canon":"Kinh điển","Home":"Trang chính","Setting":"Thiết lập","Translation":"Dịch"},"zh_TW":{"About":"關於","Canon":"經典","Home":"首頁","Setting":"設定","Translation":"翻譯"}}`

type msgIdStrPairs map[string]string
type localesMsg map[string]msgIdStrPairs

var msg = localesMsg{}

func setupJSON() {
	dec := json.NewDecoder(strings.NewReader(jsondata))
	if err := dec.Decode(&msg); err != nil {
		panic(err)
	}
}

func gettext(locale, str string) string {
	if val, ok := msg[locale]; ok {
		if val2, ok2 := val[str]; ok2 {
			return val2
		} else {
			return str
		}
	} else {
		return str
	}
}

func translate(locale string) {
	d := js.Global.Get("document")
	nodeList := d.Call("querySelectorAll", "[data-default-string]")
	length := nodeList.Get("length").Int()
	for i := 0; i < length; i++ {
		element := nodeList.Call("item", i)
		str := element.Get("dataset").Get("defaultString").String()
		element.Set("textContent", gettext(locale, str))
	}
}

func main() {
	setupJSON()

	d := js.Global.Get("document")
	nodeList := d.Call("querySelectorAll", "button")
	length := nodeList.Get("length").Int()
	for i := 0; i < length; i++ {
		btn := nodeList.Call("item", i)
		btn.Call("addEventListener", "click", func() {
			translate(btn.Get("value").String())
		})
	}
}
