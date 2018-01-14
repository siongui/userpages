Convert Text to HTML Link in Go
###############################

:date: 2018-01-14T23:11+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       String Manipulation, Regular Expression
:category: Frontend Programming in Go
:summary: Convert specific texts to clickable links (<a> tag) in HTML document
          via Go/GopherJS.
:og_image: http://talkerscode.com/webtricks/images/convert_url.jpg
:adsu: yes


Convert some specific plain texts to clickable links in HTML documents
via Go/GopherJS.
This is actually the Go/GopherJS translation of my previous post
`[JavaScript] Convert Text to Link in HTML`_ [1]_.
Please read my previous post and see the demo first.
Here I only show the Go code of the implementation. The HTML code is the same as
that in my previous post.
The full code example of this post is `on my GitHub`_.

.. code-block:: go

  package main
  
  import (
  	. "github.com/siongui/godom"
  	"regexp"
  )
  
  var textUrl = map[string]string{
  	"世尊譯詞的探討": "http://agama.buddhason.org/book/bb/bb21.htm",
  }
  
  var text = regexp.MustCompile(`〈.+〉`)
  
  func TextToLink(elm *Object) {
  	h := text.ReplaceAllStringFunc(elm.InnerHTML(), func(match string) string {
  		key := match[3 : len(match)-3]
  		url, ok := textUrl[key]
  		if ok {
  			return `〈<a href="` + url +
  				`" target="_blank">` + key + `</a>〉`
  		}
  		return match
  	})
  	elm.SetInnerHTML(h)
  }
  
  func main() {
  	b := Document.QuerySelector(".line-block")
  	TextToLink(b)
  }

The above code use godom_ package to make the code more readable.

.. adsu:: 2

----

References:

.. [1] `[JavaScript] Convert Text to Link in HTML <{filename}../../../2018/01/12/javascript-convert-text-to-link-in-html%en.rst>`_

.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/020-convert-text-to-link
.. _[JavaScript] Convert Text to Link in HTML: {filename}../../../2018/01/12/javascript-convert-text-to-link-in-html%en.rst
