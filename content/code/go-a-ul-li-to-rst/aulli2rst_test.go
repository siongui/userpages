package aulli2rst

import (
	"strings"
	"testing"
)

var testHtml = `<!DOCTYPE html><html>
<head><title>a ul li to rst</title></head>
<body>
  <ul>
    <li>item 1</li>
    <li>item 2 <a href="/">link 1</a></li>
    <li>item 3
      <ul>
        <li>item 3-1</li>
        <li>item 3-2</li>
      </ul>
    </li>
  </ul>
  <a href="/">link 2</a>
</body></html>`

func TestHtmlAUlLiToRst(t *testing.T) {
	print(HtmlAUlLiToRst(strings.NewReader(testHtml)))
}
