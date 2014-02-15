package mylib

import "testing"


func TestPoll(t *testing.T) {
	const filepath = "sqlite3/Feeder.opml"
	sites := GetOutlineList(filepath)
	setTestPollingInterval()
	Poll(sites[:1])

	// block here
	for { select {} }
}
