package resize

import (
	"github.com/nfnt/resize"
	"image/png"
	"io"
	"os"
)

func ResizePng(r io.Reader, filepath string) {
	// decode png into image.Image
	img, err := png.Decode(r)
	if err != nil {
		panic(err)
	}

	// resize to width 60 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(60, 0, img, resize.Lanczos3)

	out, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
