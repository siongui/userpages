package mylib

import "testing"

func TestGetSeed(t *testing.T) {
	v, ok := GetSeed("https://news.ycombinator.com/rss")
	if !ok { panic("GetSeed not ok") }
	for _, item := range v.Channel.ItemList {
		t.Log(item)
	}
}
