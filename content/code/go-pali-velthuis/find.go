package velthuis

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var paliWordsInVelthuisScheme = regexp.MustCompile(`[abcdeghijklmnoprstuvyABCDEGHIJKLMNOPRSTUVY"~.]+`)

func replacePaliWordsInVelthuisScheme(b []byte) []byte {
	if len(b) == 1 {
		return b
	}
	if string(b) == ".." {
		return b
	}
	fmt.Println(string(b))
	return b
}

func processRst(path string) {
	fmt.Println("processing " + path + " ...")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	paliWordsInVelthuisScheme.ReplaceAllFunc(b, replacePaliWordsInVelthuisScheme)
}

func FindPaliWordsInVelthuisScheme(dirname string) {
	// walk all files in directory
	filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".rst") {
			processRst(path)
		}
		return nil
	})
}
