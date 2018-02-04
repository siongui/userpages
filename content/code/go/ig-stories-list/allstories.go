package igstory

// Get all stories of users that a user follows

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const UrlAllStories = `https://i.instagram.com/api/v1/feed/reels_tray/`

func GetAllStories(cfg map[string]string) (err error) {
	req, err := http.NewRequest("GET", UrlAllStories, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: cfg["ds_user_id"]})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: cfg["sessionid"]})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: cfg["csrftoken"]})

	req.Header.Set("User-Agent", cfg["User-Agent"])

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

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	println(string(b))

	return
}
