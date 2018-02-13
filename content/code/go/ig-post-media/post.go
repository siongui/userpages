package igmedia

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const urlPost = `https://www.instagram.com/p/{{CODE}}/?__a=1`
const userAgent = "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+"

type postInfo struct {
	GraphQL struct {
		ShortcodeMedia EdgeMedia `json:"shortcode_media"`
	} `json:"graphql"`
}

type EdgeMedia struct {
	Typename   string `json:"__typename"`
	Shortcode  string `json:"shortcode"`
	Dimensions struct {
		Height int64 `json:"height"`
		Width  int64 `json:"width"`
	} `json:"dimensions"`
	DisplayUrl       string `json:"display_url"`
	DisplayResources []struct {
		Src          string `json:"src"`
		ConfigWidth  int64  `json:"config_width"`
		ConfigHeight int64  `json:"config_height"`
	} `json:"display_resources"`
	VideoUrl         string `json:"video_url"`
	IsVideo          bool   `json:"is_video"`
	TakenAtTimestamp int64  `json:"taken_at_timestamp"`
	Location         struct {
		Id            string `json:"id"`
		HasPublicPage bool   `json:"has_public_page"`
		Name          string `json:"name"`
		Slug          string `json:"slug"`
	} `json:"location"`
	EdgeSidecarToChildren struct {
		Edges []struct {
			Node EdgeMedia `json:"node"`
		} `json:"edges"`
	} `json:"edge_sidecar_to_children"`
}

// return URL of image with best resolution
func (em *EdgeMedia) getImageUrl() string {
	res := em.DisplayResources
	return res[len(res)-1].Src
}

func (em *EdgeMedia) getVideoUrl() string {
	return em.VideoUrl
}

func (em *EdgeMedia) printEdgeMediaChildInfo() {
	indentation := "   "
	fmt.Println(indentation + em.Typename)

	switch em.Typename {
	case "GraphImage":
		fmt.Println(indentation + em.getImageUrl())
	case "GraphVideo":
		fmt.Println(indentation + em.getVideoUrl())
	default:
		panic(em.Typename)
	}
	fmt.Println("")
}

func (em *EdgeMedia) printEdgeMediaInfo() {
	fmt.Println(em.Typename)
	fmt.Println(stripQueryString(codeToUrl(em.Shortcode)))

	// print media (photos and videos) links
	switch em.Typename {
	case "GraphImage":
		fmt.Println(em.getImageUrl())
	case "GraphVideo":
		fmt.Println(em.getVideoUrl())
	case "GraphSidecar":
		fmt.Println("")
		for _, edge := range em.EdgeSidecarToChildren.Edges {
			edge.Node.printEdgeMediaChildInfo()
		}
	default:
		panic(em.Typename)
	}

	printTimestamp(em.TakenAtTimestamp)
	fmt.Println("")
}

// Given the code of the post, return url of the post.
func codeToUrl(code string) string {
	return strings.Replace(urlPost, "{{CODE}}", code, 1)
}

func printTimestamp(timestamp int64) {
	fmt.Println(formatTimestamp(timestamp))
}

func formatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

// Remove query string in the URL
func stripQueryString(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	return u.String()
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

// Given code of post, return information of the post.
func GetPostInfo(code, ds_user_id, sessionid, csrftoken string) (em EdgeMedia, err error) {
	url := codeToUrl(code)
	b, err := getHTTPResponse(url, ds_user_id, sessionid, csrftoken)
	if err != nil {
		return
	}

	pi := postInfo{}
	err = json.Unmarshal(b, &pi)
	if err != nil {
		return
	}
	em = pi.GraphQL.ShortcodeMedia
	return
}
