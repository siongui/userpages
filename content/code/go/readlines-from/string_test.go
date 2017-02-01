package lines

import "testing"

func TestStringToLines(t *testing.T) {
	str := "hello\nworld\n"
	lines, err := StringToLines(str)
	if err != nil {
		t.Error(err)
	}

	for _, line := range lines {
		println(line)
	}
}
