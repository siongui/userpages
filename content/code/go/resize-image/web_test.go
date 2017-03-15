package resize

import (
	"net/http"
	"testing"
)

func TestResizePng(t *testing.T) {
	// open image on web
	resp, err := http.Get("https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_120x44dp.png")
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	ResizePng(resp.Body, "test_resized.png")
}
