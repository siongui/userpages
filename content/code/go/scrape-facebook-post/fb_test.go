package parsefb

import (
	"testing"
)

func TestParse(t *testing.T) {
	url := "https://www.facebook.com/jayasaro.panyaprateep.org/posts/1095007907274561:0"
	postHtml, err := Parse(url)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(postHtml)

	timestamp, err := GetTimeStamp(postHtml)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(timestamp)

	pl, err := GetProfileLink(postHtml)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(pl.Name)
	t.Log(pl.Url)

	imgurl, err := GetImageUrl(postHtml)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(imgurl)

	content, err := GetContent(postHtml)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(content)
}
