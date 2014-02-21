package mylib

import (
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
	log.Println(l)
}
