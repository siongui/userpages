package write

import (
	"testing"
)

func TestBadWriteStringToFile(t *testing.T) {
	if err := BadWriteStringToFile("bad.txt", "% in string\n"); err != nil {
		t.Error(err)
	}

	if err := WriteStringToFile("good.txt", "% in string\n"); err != nil {
		t.Error(err)
	}
}
