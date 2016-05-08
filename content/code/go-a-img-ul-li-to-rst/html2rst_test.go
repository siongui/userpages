package html2rst

import (
	"net/http"
	"testing"
)

func TestHtmlToRst(t *testing.T) {
	//resp, err := http.Get("http://nanda.online-dhamma.net/")
	resp, err := http.Get("https://siongui.github.io/zh/2016/03/14/pillow-useful-items-for-me-notes/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	print(HtmlToRst(resp.Body))
}
