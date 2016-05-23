package velthuis

import (
	"io/ioutil"
	"testing"
)

func TestProcessBytes(t *testing.T) {
	path := "/home/siongui/dev/nanda/content/articles/tipitaka/tipitaka%zh.rst"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	ProcessBytes(b)
	//b2 := ProcessBytes(b)
	//ioutil.WriteFile(path, b2, 0644)
}
