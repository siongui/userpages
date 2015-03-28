package mylib

import "testing"

func TestSaveLinkAsJson(t *testing.T) {
	l1 := LinkInfo{"Google", "https://www.google.com/"}
	l2 := LinkInfo{"DuckDuckGo", "https://duckduckgo.com/"}
	d := "./links/"
	t.Log(l1, d)
	SaveLinkAsJson(l1, d)
	t.Log(l2, d)
	SaveLinkAsJson(l2, d)
}
