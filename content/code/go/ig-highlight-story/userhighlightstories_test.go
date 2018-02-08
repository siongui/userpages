package igstory

import (
	"os"
	"testing"
)

func ExampleGetUserHighlightStories(t *testing.T) {
	SetUserId(os.Getenv("IG_DS_USER_ID"))
	SetSessionId(os.Getenv("IG_SESSIONID"))
	SetCsrfToken(os.Getenv("IG_CSRFTOKEN"))
	// id of user *instagram* is 25025320
	b, err := GetUserHighlightStories(25025320)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(b))
}
