package mylib

import (
	"testing"
)


func TestDbAll(t *testing.T) {
	l := LinkInfo{
		Title: "Google",
		Link: "http://www.google.com/",
		Tag: "test",
		Language: "en",
	}

	SaveLinkAsJson(l)
}
