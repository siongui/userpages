package getelementbyid

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

const indexHtml = `<!DOCTYPE html>
<html>
<head><title>[Go] HTML table to reStructuredText list-table</title></head>
<body>
  <table>
    <tr><td id="foo">R1, C1</td><td>R1, C2</td></tr>
    <tr><td>R2, C1</td><td>R2, C2</td></tr>
  </table>
</body>
</html>`

func TestTable2Rst(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(indexHtml))
	if err != nil {
		panic("Fail to parse!")
	}

	r1 := getElementById(doc, "foo")
	if r1.Data != "td" || r1.FirstChild.Data != "R1, C1" {
		t.Error("wrong element whose id is foo")
	}

	r2 := getElementById(doc, "foo2")
	if r2 != nil {
		t.Error("foo2 should not exist!")
	}
}
