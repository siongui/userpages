package mylib

import (
	"testing"
	"fmt"
)

const filepath = "Feeder_test.opml"

func TestGetSeed(t *testing.T) {
	siteList := GetOutlineList(filepath)
	for _, site := range siteList {
		v := GetSeed(site.XmlUrl)
		fmt.Println(v)
		break
	}
}
