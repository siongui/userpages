[Golang] Get Instagram Following and Followers
##############################################

:date: 2018-02-12T21:39+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Instagram
:category: Go
:summary: Get Instagram following and followers of a specific user via Go
          programming language. Use only Go standard library and no third-party
          packages needed.
:og_image: https://i.ytimg.com/vi/GOuxesm0Ftk/maxresdefault.jpg
:adsu: yes

Interesting small program to get the JSON format data of following and followers
of a specific Instagram_ user.

The procedure of reading JSON data is the same as the Instagram story API [1]_.
Please read the post if you are not familiar with the API.

In this program only Go standard library is used, no third-party packages.

To access the Instagram API via local Go program, you need to login Instagram_
and get the following information from your browser:

- ``ds_user_id``
- ``sessionid``
- ``csrftoken``

Please see this `SO answer`_ to get above values on Chrome browser.

.. image:: https://i.stack.imgur.com/psJLZ.png
   :align: center
   :alt: ds_user_id sessionid csrftoken

Moreover, you need to know the user id of the specific user, please read my
previous post to get the id of the user. [3]_

After you get the values, you can get the JSON response from the following code:

.. code-block:: go

  package igfollow

  import (
  	"encoding/json"
  	"errors"
  	"net/http"
  	"strconv"
  	"strings"
  )

  const UrlFollowers = `https://i.instagram.com/api/v1/friendships/{{USERID}}/followers/`
  const UrlFollowing = `https://i.instagram.com/api/v1/friendships/{{USERID}}/following/`

  type RawFollow struct {
  	Users     []FollowUser `json:"users"`
  	BigList   bool         `json:"big_list"` // if false, no next_max_id in response
  	PageSize  int64        `json:"page_size"`
  	Status    string       `json:"status"`
  	NextMaxId string       `json:"next_max_id"` // used for pagination if list is too big
  }

  type FollowUser struct {
  	Pk       int64  `json:"pk"`
  	Username string `json:"username"`
  }

  // id: Id of the Instagram user whose followers to be retrieved.
  //
  // ds_user_id, sessionid, csrftoken: Login Instagram and get these three cookie
  // values from Chrome Developer Tools. See https://stackoverflow.com/a/44773079
  func GetFollowers(id, ds_user_id, sessionid, csrftoken string) (rf RawFollow, err error) {
  	url := strings.Replace(UrlFollowers, "{{USERID}}", id, 1)
  	rf, err = GetFollowResponse(url, ds_user_id, sessionid, csrftoken)
  	if err != nil {
  		return
  	}

  	// If the list is too big and next_max_id is not ""
  	for rf.NextMaxId != "" {
  		urln := url + "?max_id=" + rf.NextMaxId
  		rfn, err := GetFollowResponse(urln, ds_user_id, sessionid, csrftoken)
  		if err != nil {
  			return rf, err
  		}
  		rf.Users = append(rf.Users, rfn.Users...)
  		rf.NextMaxId = rfn.NextMaxId
  	}
  	return
  }

  // id: Id of the Instagram user whose following to be retrieved.
  //
  // ds_user_id, sessionid, csrftoken: Login Instagram and get these three cookie
  // values from Chrome Developer Tools. See https://stackoverflow.com/a/44773079
  func GetFollowing(id, ds_user_id, sessionid, csrftoken string) (rf RawFollow, err error) {
  	url := strings.Replace(UrlFollowing, "{{USERID}}", id, 1)
  	rf, err = GetFollowResponse(url, ds_user_id, sessionid, csrftoken)
  	if err != nil {
  		return
  	}

  	// If the list is too big and next_max_id is not ""
  	for rf.NextMaxId != "" {
  		urln := url + "?max_id=" + rf.NextMaxId
  		rfn, err := GetFollowResponse(urln, ds_user_id, sessionid, csrftoken)
  		if err != nil {
  			return rf, err
  		}
  		rf.Users = append(rf.Users, rfn.Users...)
  		rf.NextMaxId = rfn.NextMaxId
  	}
  	return
  }

  func GetFollowResponse(url, ds_user_id, sessionid, csrftoken string) (rf RawFollow, err error) {
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
  		err = errors.New("resp.StatusCode: " + strconv.Itoa(resp.StatusCode))
  		return
  	}

  	dec := json.NewDecoder(resp.Body)
  	err = dec.Decode(&rf)
  	return
  }

.. adsu:: 2

**Example**:

.. code-block:: go

  package igfollow

  import (
  	"os"
  	"testing"
  )

  func ExampleGetFollowing(t *testing.T) {
  	rf, err := GetFollowing(
  		os.Getenv("IG_TEST_USER_ID"),
  		os.Getenv("IG_DS_USER_ID"),
  		os.Getenv("IG_SESSIONID"),
  		os.Getenv("IG_CSRFTOKEN"))
  	if err != nil {
  		t.Error(err)
  		return
  	}
  	t.Log(rf)
  	t.Log(len(rf.Users))
  }

  func ExampleGetFollowers(t *testing.T) {
  	rf, err := GetFollowers(
  		os.Getenv("IG_TEST_USER_ID"),
  		os.Getenv("IG_DS_USER_ID"),
  		os.Getenv("IG_SESSIONID"),
  		os.Getenv("IG_CSRFTOKEN"))
  	if err != nil {
  		t.Error(err)
  		return
  	}
  	t.Log(rf)
  	t.Log(len(rf.Users))
  }

.. adsu:: 3

The full code is also available on my GitHub repo [2]_.

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.4``.

----

References:

.. [1] `Chrome IG Story — Bribing the Instagram Story API with cookies 🍪🍪🍪 <https://medium.com/@calialec/chrome-ig-story-bribing-the-instagram-story-api-with-cookies-c813e6dff911>`_
.. [2] `GitHub - siongui/goigfollow: Get Instagram following and followers in Go <https://github.com/siongui/goigfollow>`_
.. [3] `[Golang] Get Instagram User ID <{filename}/articles/2018/02/04/go-get-instagram-user-id%en.rst>`_

.. _Instagram: https://www.instagram.com/
.. _SO answer: https://stackoverflow.com/a/44773079
