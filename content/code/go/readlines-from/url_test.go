package lines

import "testing"

func TestUrlToLines(t *testing.T) {
	lines, err := UrlToLines("https://raw.githubusercontent.com/siongui/userpages/master/README.rst")
	if err != nil {
		t.Error(err)
	}

	for _, line := range lines {
		println(line)
	}
}
