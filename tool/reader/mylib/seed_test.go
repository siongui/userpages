package mylib

import (
	"testing"
	"log"
)


func TestGetSeed(t *testing.T) {
	v := GetSeed("https://news.ycombinator.com/rss")
	for _, item := range v.Channel.ItemList {
		log.Println(item)
	}
}
