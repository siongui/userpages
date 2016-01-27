package main

import "honnef.co/go/js/dom"
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
		} else { return str }
	} else { return str }
}

func translate(value string) {
	d := dom.GetWindow().Document()
	elements := d.QuerySelectorAll("[data-default-string]")
	for _, element := range elements {
		elm := element.(*dom.HTMLDivElement)
		elm.SetTextContent(gettext(value, elm.Dataset()["defaultString"]))
	}
}

func main() {
	setupJSON()

	d := dom.GetWindow().Document()
	buttons := d.QuerySelectorAll("button")

	for _, btn := range buttons {
		elm := btn.(*dom.HTMLButtonElement)
		elm.AddEventListener("click", false, func(event dom.Event) {
			translate(elm.Value)
		})
	}
}
