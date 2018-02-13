package igpost

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// no need to login or cookie to access this URL. But if login to Instagram,
// private account will return private data if you are allowed to view the
// private account.
const urlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`

const userAgent = "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+"

// used to decode the JSON data
type RawUserResp struct {
	User UserInfo `json:"user"`
}

// You can add more fields in the struct to get more information
// See response/types.go in github.com/ahmdrz/goinsta
type UserInfo struct {
	Biography       string `json:"biography"`
	ExternalUrl     string `json:"external_url"`
	FullName        string `json:"full_name"`
	Id              string `json:"id"`
	IsPrivate       bool   `json:"is_private"`
	ProfilePicUrlHd string `json:"profile_pic_url_hd"`
	Username        string `json:"username"`
	Media           struct {
		Nodes    []MediaNode `json:"nodes"`
		Count    int64       `json:"count"`
		PageInfo struct {
			HasNextPage bool   `json:"has_next_page"`
			EndCursor   string `json:"end_cursor"`
		} `json:"page_info"`
	} `json:"media"`
}

type MediaNode struct {
	Code    string `json:"code"` // url of the post
	Date    int64  `json:"date"`
	Caption string `json:"caption"`
}

// Send HTTP request and get http response on behalf of a specific Instagram
// user. After login to Instagram, you can get the cookies of *ds_user_id*,
// *sessionid*, *csrftoken* in Chrome Developer Tools.
// See https://stackoverflow.com/a/44773079
// or
// https://github.com/hoschiCZ/instastories-backup#obtain-cookies
func getHTTPResponse(url, ds_user_id, sessionid, csrftoken string) (b []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: ds_user_id})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: sessionid})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: csrftoken})

	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(
			"resp.StatusCode: " +
				strconv.Itoa(resp.StatusCode))
		return
	}

	return ioutil.ReadAll(resp.Body)
}

// Given user name, return codes of all posts of the user.
func GetAllPostCode(username, ds_user_id, sessionid, csrftoken string) (codes []string, err error) {
	r := RawUserResp{}
	r.User.Media.PageInfo.HasNextPage = true
	for r.User.Media.PageInfo.HasNextPage == true {
		url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
		if len(codes) != 0 {
			url = url + "&max_id=" + r.User.Media.PageInfo.EndCursor
		}

		b, err := getHTTPResponse(url, ds_user_id, sessionid, csrftoken)
		if err != nil {
			return codes, err
		}

		if err = json.Unmarshal(b, &r); err != nil {
			return codes, err
		}

		for _, node := range r.User.Media.Nodes {
			codes = append(codes, node.Code)
		}
		fmt.Printf("Getting %d from %s ...\n", len(codes), url)
	}
	return
}
