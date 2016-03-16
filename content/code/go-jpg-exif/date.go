package jpgdate

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"io/ioutil"
	"os"
	"path"
)

func ReadJpegDate(dirname string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fpath := path.Join(dirname, file.Name())
		println(fpath)
		f, err := os.Open(fpath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		x, err := exif.Decode(f)
		if err != nil {
			panic(err)
		}
		tm, _ := x.DateTime()
		fmt.Println(tm.Date())
	}
}
