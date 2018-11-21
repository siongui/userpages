[Golang] HTTP Request With Cookies
##################################

:date: 2018-03-03T20:21+08:00
:tags: Go, Golang, Web Scrape, Go net/http, HTTP GET, HTTP Header
:category: Go
:summary: Send HTTP request with cookies via Go standard *net/http* package.
:og_image: https://www.georanker.com/app/uploads/2014/09/cookie.jpg
:adsu: yes


Example of sending HTTP request with cookies_ via Go standard `net/http`_
package.

.. code-block:: go

  package main

  import (
  	"errors"
  	"io/ioutil"
  	"net/http"
  	"strconv"
  )

  func HTTPwithCookies(url, ds_user_id, sessionid, csrftoken string) (b []byte, err error) {
  	req, err := http.NewRequest("GET", url, nil)
  	if err != nil {
  		return
  	}

  	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: ds_user_id})
  	req.AddCookie(&http.Cookie{Name: "sessionid", Value: sessionid})
  	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: csrftoken})

  	client := &http.Client{}
  	resp, err := client.Do(req)
  	if err != nil {
  		return
  	}
  	defer resp.Body.Close()

  	if resp.StatusCode != 200 {
  		err = errors.New(url +
  			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
  		return
  	}

  	return ioutil.ReadAll(resp.Body)
  }

  func main() {
  	b, err := HTTPwithCookies("https://www.instagram.com/", "my_id", "my_sessionid", "my_csrftoken")
  	if err != nil {
  		panic(err)
  	}
  	println(string(b))
  }

If you want to know how to send HTTP request with headers, for example,
*User-Agent* header, please see [1]_.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``.

----

References:

.. [1] `[Golang] HTTP Request With Custom User-Agent <{filename}/articles/2018/02/27/go-http-request-with-custom-user-agent%en.rst>`_
.. [2] `soup - Web Scraper for Go with a similar interface to BeautifulSoup : golang <https://old.reddit.com/r/golang/comments/9x2btf/soup_web_scraper_for_go_with_a_similar_interface/>`_

.. _cookies: https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies
.. _net/http: https://golang.org/pkg/net/http/
