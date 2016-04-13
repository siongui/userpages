package table2rst

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
    <tr><td>R1, C1</td><td>R1, C2</td></tr>
    <tr><td>R2, C1</td><td>R2, C2</td></tr>
  </table>
</body>
</html>`

const tableRst = `.. list-table::

  * - R1, C1
    - R1, C2
  * - R2, C1
    - R2, C2
`

func TestTable2Rst(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(indexHtml))
	if err != nil {
		panic("Fail to parse!")
	}

	if HtmlTableToRstListTable(doc) != tableRst {
		t.Error("Fail to convert html table to rst")
	}
}
