package minhtml

import (
	"io/ioutil"
	"testing"
)

var html = []byte(`
<!--- hello --->
<div id="foo"
name="foo2"></div>
<!--
world
-->
<span></span><!--ddd-->
`)

var htmlNoComments = []byte(`

<div id="foo"
name="foo2"></div>

<span></span>
`)

var htmlMinified = `<div id="foo" name="foo2"></div> <span></span> `

func TestRemoveHTMLComments(t *testing.T) {
	if string(RemoveHTMLComments(html)) != string(htmlNoComments) {
		t.Error("Remove HTML comments failure!")
	}
	//println(string(RemoveHTMLComments(html)))
}

func TestMinifyHTML(t *testing.T) {
	if MinifyHTML(html) != htmlMinified {
		t.Error("Minify HTML failure!")
	}
	//println(MinifyHTML(html))
}

func TestHTMLFile(t *testing.T) {
	htmlPath := "../../../output/index.html"
	b, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		panic(err)
	}
	minifiedHTML := MinifyHTML(b)
	println(minifiedHTML)
	if err = ioutil.WriteFile(htmlPath, []byte(minifiedHTML), 0644); err != nil {
		panic(err)
	}
}
