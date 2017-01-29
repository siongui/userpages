package insert

import "testing"

func TestCountLeadingSpace(t *testing.T) {
	err := InsertStringToFile("test.txt", "hello world!\n", 2)
	if err != nil {
		t.Error(err)
	}
}
