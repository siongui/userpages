package mylib

import "testing"

func TestGetSeed(t *testing.T) {
	v := GetSeed("https://news.ycombinator.com/rss")
	for _, item := range v.Channel.ItemList {
		t.Log(item)
	}
}
