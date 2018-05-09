[Golang] Get Instagram User Information From HTML Source
########################################################

:date: 2018-04-12T23:00+08:00
:tags: Go, Golang, Web Scrape, Go net/http, Go strings Package, JSON, Instagram,
       Regular Expression
:category: Go
:summary: Given user name, get Instagram user information, such as id, username,
          biography, etc., in Go.
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Instagram_logo_2016.svg/1024px-Instagram_logo_2016.svg.png
:adsu: yes


When I visited Instagram_ user profile page, looked at the HTML source code, I
accidentally found that the some user information is embedded in the
*window._sharedData* variable in *script* tag. The user information is stored in
JSON format, so I wrote a small program to retrieve the information from the
HTML source. The following is complete source code:

.. code-block:: go

  package main

  import (
  	"encoding/json"
  	"fmt"
  	"io/ioutil"
  	"net/http"
  	"regexp"
  	"strings"
  )

  type sharedData struct {
  	EntryData struct {
  		ProfilePage []struct {
  			GraphQL struct {
  				User IGUser `json:"user"`
  			} `json:"graphql"`
  		} `json:"ProfilePage"`
  	} `json:"entry_data"`
  }

  type IGUser struct {
  	Biography     string `json:"biography"`
  	Id            string `json:"id"`
  	Username      string `json:"username"`
  	ProfilePicUrl string `json:"profile_pic_url"`
  }

  func getSource(username string) (b []byte, err error) {
  	resp, err := http.Get("https://www.instagram.com/" + username + "/")
  	if err != nil {
  		return
  	}
  	defer resp.Body.Close()

  	return ioutil.ReadAll(resp.Body)
  }

  func getJsonStr(b []byte) string {
  	pattern := regexp.MustCompile(`<script type="text\/javascript">window\._sharedData = (.*?);<\/script>`)
  	m := string(pattern.Find(b))
  	m1 := strings.TrimPrefix(m, `<script type="text/javascript">window._sharedData = `)
  	return strings.TrimSuffix(m1, `;</script>`)
  }

  func decodeJsonString(s string) (user IGUser, err error) {
  	d := sharedData{}
  	err = json.Unmarshal([]byte(s), &d)
  	user = d.EntryData.ProfilePage[0].GraphQL.User
  	return
  }

  func main() {
  	b, err := getSource("instagram")
  	if err != nil {
  		panic(err)
  	}

  	jsonStr := getJsonStr(b)
  	user, err := decodeJsonString(jsonStr)
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(user)
  }

If you want, you can take a look at the JSON data and get more information of
the user.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``.

----

References:

.. [1] `Instagram API -Get the userId - Stack Overflow <https://stackoverflow.com/a/44773079>`_

.. _Instagram: https://www.instagram.com/
