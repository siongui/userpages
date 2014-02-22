package mylib

import (
	"net/url"
	"encoding/json"
	"os"
	"io/ioutil"
	"log"
)

type LinkInfo struct {
	// required
	Title		string
	Link		string
	Tag		string
	Language	string
	// optional
	HNComments	string
}

func SaveLinkAsJson(l LinkInfo) {
	path := url.QueryEscape(l.Link)
	os.Remove(path)

	b, err := json.Marshal(l)
	if err != nil { log.Println(err) }

	ioutil.WriteFile(path, b, 0644)
}
