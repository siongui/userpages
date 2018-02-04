package igstory

import (
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	iginfo, err := GetUserInfo("instagram")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(iginfo.Id)
	t.Log(iginfo.Biography)
}
