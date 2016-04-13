package table2rst

import (
	"testing"
)

func TestTable2Rst(t *testing.T) {
	htmlTableToRst("http://nanda.online-dhamma.net/Tipitaka/Sutta/Majjhima/mn.047.contrast-reading.html", "output.rst")
}
