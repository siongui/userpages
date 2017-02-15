package taobaoitem2rst

import (
	"net/url"
)

func NormalizeURL(inputUrl string) string {
	u, err := url.Parse(inputUrl)

	if u.Host != "item.taobao.com" {
		return inputUrl
	}

	if err != nil {
		panic(err)
	}
	u.RawQuery = "id=" + u.Query().Get("id")
	return u.String()
}
