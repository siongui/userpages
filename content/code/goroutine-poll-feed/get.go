// http://tour.golang.org/concurrency/1
// http://golang.org/pkg/net/http/
package main

import (
	"time"
	"net/http"
	"io/ioutil"
	"fmt"
)

var pollingInterval = 1 * time.Minute

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(body[:10])
}

func getFeed(url string) {
	for {
		fetch(url)
		time.Sleep(pollingInterval)
	}
}

func poll(feedUrls []string) {
	for _, url := range feedUrls {
		go getFeed(url)
	}
}

func main() {
	feedUrls := []string{
	"https://news.ycombinator.com/rss",
	"http://www.csdn.net/article/rss_lastnews",
	"http://www.solidot.org/index.rss",
	"http://blog.jobbole.com/feed/"}

	poll(feedUrls)

	// block here
	for { select {} }
}
