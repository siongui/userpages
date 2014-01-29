package mylib

import (
	"testing"
)


func TestPoll(t *testing.T) {
	const filepath = "Feeder.opml"
	Poll(GetOutlineList(filepath))
}
