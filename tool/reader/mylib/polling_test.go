package mylib

import "testing"


func TestPoll(t *testing.T) {
	const filepath = "Feeder.opml"
	sites := GetOutlineList(filepath)
	Poll(sites[:1])

	// block here
	for { select {} }
}
