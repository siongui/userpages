package igstory

import (
	"testing"
)

func TestGetAllStories(t *testing.T) {
	err := GetAllStories(config)
	if err != nil {
		t.Error(err)
		return
	}
}
