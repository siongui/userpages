package mincss

import "testing"

func TestMinifyCSS(t *testing.T) {
	cssPathes := []string{
		"/home/siongui/dev/pali/tipitaka/app/css/tipitaka-latn.css",
		"/home/siongui/dev/pali/tipitaka/app/css/app.css",
	}
	minifiedCss := MinifyCSS(cssPathes)
	println(minifiedCss)
}
