package fileqrcode

import (
	"testing"
)

func TestFileToQrcode(t *testing.T) {
	err := FileToQrcode("path_to_my/robots.txt")
	if err != nil {
		t.Error(err)
	}
}

func TestQrcodeToFile(t *testing.T) {
	err := QrcodeToFile("cm9ib3RzLnR4dA==.png")
	if err != nil {
		t.Error(err)
	}
}
