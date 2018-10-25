package guess

import (
	"testing"
)

func TestUrlEncoding(t *testing.T) {
	name, _, err := UrlEncoding("http://shenfang.com.tw/")
	if err != nil {
		t.Error(err)
		return
	}
	if name != "big5" {
		t.Error("bad guess!")
		return
	}

	name, _, err = UrlEncoding("https://siongui.github.io/")
	if err != nil {
		t.Error(err)
		return
	}
	if name != "utf-8" {
		t.Error("bad guess!")
		return
	}
}
