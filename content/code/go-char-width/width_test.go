package mylib

import "testing"

func TestGetWidthUTF8String(t *testing.T) {
	if GetWidthUTF8String("s: 你好") != 7 {
		t.Error("s: 你好")
	}
	if GetWidthUTF8String("s --asji勿勿") != 12 {
		t.Error("s --asji勿勿")
	}
}
