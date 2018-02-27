[Golang] HTTP Request With Custom User-Agent
############################################

:date: 2018-02-27T18:18+08:00
:tags: Go, Golang, Web Scrape, Go net/http
:category: Go
:summary: HTTP request with custom *User-Agent* header
          via Go standard *net/http* package.
:og_image: https://cdn-images-1.medium.com/max/1200/1*Gj7_v-DbYPEpMwyD1147cA.png
:adsu: yes


Example of HTTP request with custom `User-Agent`_ header via Go standard
`net/http`_ package.

.. code-block:: go

  package main

  import (
  	"errors"
  	"io/ioutil"
  	"net/http"
  	"strconv"
  )

  func HTTPRequestCustomUserAgent(url, userAgent string) (b []byte, err error) {
  	req, err := http.NewRequest("GET", url, nil)
  	if err != nil {
  		return
  	}

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

  func main() {
  	b, err := HTTPRequestCustomUserAgent("https://www.google.com/", "My User-Agent")
  	if err != nil {
  		panic(err)
  	}
  	println(string(b))
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``.

----

References:

.. [1] `GitHub - siongui/instago: Get photos, videos, stories, following and followers of Instagram <https://github.com/siongui/instago>`_

.. _User-Agent: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent
.. _net/http: https://golang.org/pkg/net/http/
