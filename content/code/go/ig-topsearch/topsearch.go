package instago

// This file implements topsearch of Instagram web.

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const urlTopsearch = `https://www.instagram.com/web/search/topsearch/?query=`

// Decode JSON data returned from Instagram topsearch API
type IGTopsearchResp struct {
	Users []struct {
		Position int64 `json:"position"`
		User     struct {
			Pk               string `json:"pk"`
			Username         string `json:"username"`
			FullName         string `json:"full_name"`
			IsPrivate        bool   `json:"is_private"`
			ProfilePicUrl    string `json:"profile_pic_url"`
			ProfilePicId     string `json:"profile_pic_id"`
			FriendshipStatus struct {
				Following       bool `json:"following"`
				IsPrivate       bool `json:"is_private"`
				IncomingRequest bool `json:"incoming_request"`
				OutgoingRequest bool `json:"outgoing_request"`
				IsBestie        bool `json:"is_bestie"`
			} `json:"friendship_status"`
			IsVerified                 bool    `json:"is_verified"`
			HasAnonymousProfilePicture bool    `json:"has_anonymous_profile_picture"`
			FollowerCount              int64   `json:"follower_count"`
			Byline                     string  `json:"byline"`
			MutualFollowersCount       float64 `json:"mutual_followers_count"`
			LatestReelMedia            int64   `json:"latest_reel_media"`
		} `json:"user"`
	} `json:"users"`

	//TODO: Places ... `json:"places"`

	Hashtags []struct {
		Position int64 `json:"position"`
		Hashtag  struct {
			Name       string `json:"name"`
			Id         int64  `json:"id"`
			MediaCount int64  `json:"media_count"`
		} `json:"hashtag"`
	} `json:"hashtags"`
	HasMore          bool   `json:"has_more"`
	RankToken        string `json:"rank_token"`
	ClearClientCache bool   `json:"clear_client_cache"`
	Status           string `json:"status"`
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

	req.Header.Set("User-Agent", "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+")

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

// Given a string, return the users that best matches the string. This is
// actually *topsearch* on Instagram web.
func Topsearch(str, ds_user_id, sessionid, csrftoken string) (tr IGTopsearchResp, err error) {
	url := urlTopsearch + str
	b, err := getHTTPResponse(url, ds_user_id, sessionid, csrftoken)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &tr)
	return
}
