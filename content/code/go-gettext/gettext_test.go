package mylib

import "testing"

func TestAll(t *testing.T) {
	setup("zh_TW", "messages", "locale")
	setup("vi_VN", "messages", "locale")

	changeLocale("zh_TW")
	t.Log(translate("Home"))
	t.Log(translate("Canon"))
	t.Log(translate("About"))
	t.Log(translate("Setting"))
	t.Log(translate("Translation"))

	changeLocale("vi_VN")
	t.Log(translate("Home"))
	t.Log(translate("Canon"))
	t.Log(translate("About"))
	t.Log(translate("Setting"))
	t.Log(translate("Translation"))
}
