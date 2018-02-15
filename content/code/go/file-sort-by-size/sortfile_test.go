package sortfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func ExampleSortFileBySize(t *testing.T) {
	files, err := ioutil.ReadDir(os.Getenv("MY_DIR"))
	if err != nil {
		panic(err)
	}
	SortFileBySize(files)
	for _, file := range files {
		fmt.Println(file.Name()+": ", file.Size())
	}
}
