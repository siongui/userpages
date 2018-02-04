package igstory

// read user information, such as id,  via username
// see ``Instagram API -Get the userId - Stack Overflow``
// https://stackoverflow.com/a/44773079

import (
	"encoding/json"
	"net/http"
	"strings"
)

// no need to login or cookie to access this URL
const UrlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`

type RawUserResp struct {
	User UserInfo
}

type UserInfo struct {
	Id        string `json:"id"`
	Biography string `json:"biography"`
}

func GetUserInfo(username string) (userinfo *UserInfo, err error) {
	url := strings.Replace(UrlUserInfo, "{{USERNAME}}", username, 1)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	r := RawUserResp{}
	if err = dec.Decode(&r); err != nil {
		return
	}
	userinfo = &(r.User)

	return
}
