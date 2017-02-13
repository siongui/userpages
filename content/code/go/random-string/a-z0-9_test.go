package mylib

import "testing"

func TestRandomString(t *testing.T) {
	t.Log(RandomString(10))
	t.Log(RandomString(20))
}

func TestRandomString2(t *testing.T) {
	t.Log(RandomString2(10))
	t.Log(RandomString2(20))
}
