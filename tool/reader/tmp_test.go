package tmp

import "testing"

func TestGetSeed(t *testing.T) {
//	v, ok := GetSeed("https://news.ycombinator.com/rss")
//	v, ok := GetSeed("http://www.eettaiwan.com/rss/rss.xml")
//	v, ok := GetSeed("http://www.linuxeden.com/rss.php")
	v1, ok := GetSeed("http://zh.cn.nikkei.com/rss.html")
	if !ok { t.Fatal("GetSeed 日經中文網 not ok") }
	t.Log(v1)
	v2, ok2 := GetSeed("http://blog.go-china.org/feed.atom")
	if !ok2 { t.Fatal("GetSeed Golang 中國 not ok") }
	t.Log(v2)
}
