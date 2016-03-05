package concat

import (
	"io/ioutil"
	"testing"
)

func TestConcatenateJS(t *testing.T) {
	b := ConcatenateJS("/home/siongui/dev/pali/common/app/scripts/")
	t.Log(string(b))
	ioutil.WriteFile("all.js", b, 0644)
}
