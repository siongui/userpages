package instago

// This file implements Instagram discover top live API.

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const urlToplive = `https://i.instagram.com/api/v1/discover/top_live/`

// Decode JSON data returned from Instagram top live API
type IGTopliveResp struct {
	Broadcasts []struct {
		Id                 int64   `json:"id"`
		RtmpPlaybackUrl    string  `json:"rtmp_playback_url"`
		DashPlaybackUrl    string  `json:"dash_playback_url"`
		DashAbrPlaybackUrl string  `json:"dash_abr_playback_url"`
		BroadcastStatus    string  `json:"broadcast_status"`
		ViewerCount        float64 `json:"viewer_count"`
		InternalOnly       bool    `json:"internal_only"`
		CoverFrameUrl      string  `json:"cover_frame_url"`
		BroadcastOwner     struct {
			Pk               int64  `json:"pk"`
			Username         string `json:"username"`
			FullName         string `json:"full_name"`
			IsPrivate        bool   `json:"is_private"`
			ProfilePicUrl    string `json:"profile_pic_url"`
			ProfilePicId     string `json:"profile_pic_id"`
			FriendshipStatus struct {
				Following       bool `json:"following"`
				FollowedBy      bool `json:"followed_by"`
				Blocking        bool `json:"blocking"`
				IsPrivate       bool `json:"is_private"`
				IncomingRequest bool `json:"incoming_request"`
				OutgoingRequest bool `json:"outgoing_request"`
				IsBestie        bool `json:"is_bestie"`
			} `json:"friendship_status"`
			IsVerified bool `json:"is_verified"`
		} `json:"broadcast_owner"`
		PublishedTime        int64  `json:"published_time"`
		MediaId              string `json:"media_id"`
		BroadcastMessage     string `json:"broadcast_message"`
		OrganicTrackingToken string `json:"organic_tracking_token"`
	} `json:"broadcasts"`
	MoreAvailable       bool   `json:"more_available"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	NextMaxId           int64  `json:"next_max_id"`
	Status              string `json:"status"`
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
func Toplive(ds_user_id, sessionid, csrftoken string) (tr IGTopliveResp, err error) {
	b, err := getHTTPResponse(urlToplive, ds_user_id, sessionid, csrftoken)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &tr)
	return
}
