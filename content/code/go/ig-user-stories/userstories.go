package igstory

// Get all stories of a specific user

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const UrlUserStories = `https://i.instagram.com/api/v1/feed/user/{{USERID}}/reel_media/`

// id: the id of user whose stories to be retrieved
// userid: your user id
// sessionid: your session id
// csrftoken: your csrftoken
//
// b: the JSON bytes of user stories
func GetUserStories(id int64, userid, sessionid, csrftoken string) (b []byte, err error) {
	url := strings.Replace(UrlUserStories, "{{USERID}}", strconv.FormatInt(id, 10), 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: userid})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: sessionid})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: csrftoken})

	req.Header.Set("User-Agent", "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("resp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	return ioutil.ReadAll(resp.Body)
}
