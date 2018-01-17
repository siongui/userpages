[Golang] Remove HTML Inline Style
#################################

:date: 2018-01-16T06:28+08:00
:tags: Go, Golang, DOM, Go net/html, Web Scrape, html, CSS
:category: Go
:summary: Remove HTML inline style, i.e., remove *style* attribute from HTML
          node via Go *net/html* package.
:adsu: yes
:og_image: https://i.stack.imgur.com/k2PBF.png


Remove HTML inline style, i.e., remove *style* attribute from HTML node via Go
`net/html`_ package.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ZYivPISjA3J>`__
   :class: align-center

.. code-block:: go

  package main
  
  import (
  	"golang.org/x/net/html"
  	"os"
  	"strings"
  )
  
  var testhtml = `
  <table style="s1">
    <tr><td id="foo" style="s2">R1, C1</td><td>R1, C2</td></tr>
    <tr><td alt="hi">R2, C1</td><td style="s3">R2, C2</td></tr>
  </table>`
  
  func RemoveStyleAttr(n *html.Node) {
  	i := -1
  	for index, attr := range n.Attr {
  		if attr.Key == "style" {
  			i = index
  			break
  		}
  	}
  	if i != -1 {
  		n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
  	}
  
  	for c := n.FirstChild; c != nil; c = c.NextSibling {
  		RemoveStyleAttr(c)
  	}
  }
  
  func main() {
  	doc, err := html.Parse(strings.NewReader(testhtml))
  	if err != nil {
  		panic(err)
  	}
  
  	RemoveStyleAttr(doc)
  
  	html.Render(os.Stdout, doc)
  }

.. adsu:: 2

----

Tested on: `Go Playground`_

----

References:

.. [1] `[Golang] getElementById via net/html Package <{filename}../../../2016/04/15/go-getElementById-via-net-html-package%en.rst>`_
.. [2] | `html remove all styles online - Google search <https://www.google.com/search?q=html+remove+all+styles+online>`_
       | `html remove all styles online - DuckDuckGo search <https://duckduckgo.com/?q=html+remove+all+styles+online>`_
       | `html remove all styles online - Ecosia search <https://www.ecosia.org/search?q=html+remove+all+styles+online>`_
       | `html remove all styles online - Qwant search <https://www.qwant.com/?q=html+remove+all+styles+online>`_
       | `html remove all styles online - Bing search <https://www.bing.com/search?q=html+remove+all+styles+online>`_
       | `html remove all styles online - Yahoo search <https://search.yahoo.com/search?p=html+remove+all+styles+online>`_
       | `html remove all styles online - Baidu search <https://www.baidu.com/s?wd=html+remove+all+styles+online>`_
       | `html remove all styles online - Yandex search <https://www.yandex.com/search/?text=html+remove+all+styles+online>`_
.. [3] `go - How to delete an element from array in golang - Stack Overflow <https://stackoverflow.com/a/37335777>`_

.. _Go Playground: https://play.golang.org/
.. _net/html: https://godoc.org/golang.org/x/net/html
