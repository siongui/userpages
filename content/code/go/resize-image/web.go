package main

import (
	"github.com/nfnt/resize"
	"image/png"
	"net/http"
	"os"
)

func main() {
	// open image on web
	resp, err := http.Get("https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_120x44dp.png")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// decode png into image.Image
	img, err := png.Decode(resp.Body)
	if err != nil {
		panic(err)
	}

	// resize to width 60 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(60, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.png")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
