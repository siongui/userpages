package mylib

import "testing"

func TestRandomString(t *testing.T) {
	t.Log(RandomString(10))
	t.Log(RandomString(20))
}
