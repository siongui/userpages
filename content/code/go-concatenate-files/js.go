package concat

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ConcatenateJS(dirPath string) []byte {
	var js_code []byte
	// walk all files in directory
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".js") {
			println("concatenating " + path + " ...")
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			js_code = append(js_code, b...)
		}
		return nil
	})
	return js_code
}
