package toutf8

import (
	"io/ioutil"
	"testing"
)

func TestUrlToUtf8Encoding(t *testing.T) {
	r, name, _, err := UrlToUtf8Encoding("http://shenfang.com.tw/product-1.htm")
	if err != nil {
		t.Error(err)
		return
	}
	if name != "big5" {
		t.Error("bad guess!")
		return
	}
	b, err := ioutil.ReadAll(r)
	t.Log(string(b))

	r, name, _, err = UrlToUtf8Encoding("https://siongui.github.io/")
	if err != nil {
		t.Error(err)
		return
	}
	if name != "utf-8" {
		t.Error("bad guess!")
		return
	}
	b, err = ioutil.ReadAll(r)
	t.Log(string(b))
}
