package lines

import "testing"

func TestFileToLines(t *testing.T) {
	lines, err := FileToLines("file.go")
	if err != nil {
		t.Error(err)
	}

	for _, line := range lines {
		println(line)
	}
}
