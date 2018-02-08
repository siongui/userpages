package igstory

// Get highlight stories of a specific user

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const UrlUserHighlightStories = `https://i.instagram.com/api/v1/highlights/{{USERID}}/highlights_tray/`

func GetUserHighlightStories(id int64) (b []byte, err error) {
	url := strings.Replace(
		UrlUserHighlightStories,
		"{{USERID}}",
		strconv.FormatInt(id, 10),
		1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: config["ds_user_id"]})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: config["sessionid"]})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: config["csrftoken"]})

	req.Header.Set("User-Agent", config["User-Agent"])

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
