package title

import (
	"net/http"
	"testing"
)

func TestHtmlToRst(t *testing.T) {
	resp, err := http.Get("http://nanda.online-dhamma.net/")
	//resp, err := http.Get("https://siongui.github.io/zh/2016/03/14/pillow-useful-items-for-me-notes/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if title, ok := GetHtmlTitle(resp.Body); ok {
		println(title)
	} else {
		println("Fail to get HTML title")
	}
}
