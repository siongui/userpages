package mylib

import "testing"

func TestGetSeed(t *testing.T) {
//	v, ok := GetSeed("https://news.ycombinator.com/rss")
//	v, ok := GetSeed("http://www.eettaiwan.com/rss/rss.xml")
//	v, ok := GetSeed("http://www.linuxeden.com/rss.php")
	a, ok := GetAtom("http://blog.go-china.org/feed.atom")
	if !ok { panic("GetSeed not ok") }
	t.Log(a)
//	t.Log(v)
	/*
	for _, item := range v.Channel.ItemList {
		t.Log(item)
	}
	*/
}
