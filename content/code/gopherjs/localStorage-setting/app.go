package main

import (
	"encoding/json"
	"github.com/gopherjs/gopherjs/js"
)

var localStorage = js.Global.Get("localStorage")
var mySettingName = "mySetting"

type Setting struct {
	en    bool   `json:"en"`
	ja    bool   `json:"ja"`
	order string `json:"order"`
}

func Setting2JSON(s Setting) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func JSON2Setting(j string) Setting {
	setting := Setting{}
	err := json.Unmarshal([]byte(j), &setting)
	if err != nil {
		panic(err)
	}
	return setting
}

func SaveSetting(s Setting) {
	localStorage.Set(mySettingName, Setting2JSON(s))
}

func LoadSetting() Setting {
	return JSON2Setting(localStorage.Get(mySettingName).String())
}

func SetupSetting() {
	d := js.Global.Get("document")
	enElm := d.Call("getElementById", "en")
	jaElm := d.Call("getElementById", "ja")
	orderElm := d.Call("getElementById", "order")

	s := Setting{}

	// check if there is saved setting in user browser
	if localStorage.Get(mySettingName) == js.Undefined {
		// no setting saved, use default setting
		s.en = enElm.Get("checked").Bool()
		s.ja = jaElm.Get("checked").Bool()
		s.order = orderElm.Get("options").Call("item", orderElm.Get("selectedIndex").Int()).Get("value").String()
		SaveSetting(s)
	} else {
		// use saved setting in user browser
		s = LoadSetting()
		enElm.Set("checked", s.en)
		jaElm.Set("checked", s.ja)
		orderElm.Set("value", s.order)
	}

	enElm.Call("addEventListener", "click", func(event *js.Object) {
		s.en = enElm.Get("checked").Bool()
		SaveSetting(s)
	})

	jaElm.Call("addEventListener", "click", func(event *js.Object) {
		s.ja = jaElm.Get("checked").Bool()
		SaveSetting(s)
	})

	orderElm.Call("addEventListener", "change", func(event *js.Object) {
		s.order = orderElm.Get("options").Call("item", orderElm.Get("selectedIndex").Int()).Get("value").String()
		SaveSetting(s)
	})
}

func main() {
	SetupSetting()
}
