[Golang] Add Method to Existing Type in External Package
########################################################

:date: 2017-02-11T13:57+08:00
:tags: Go, Golang
:category: Go
:summary: Add method (function) to an existing type in external package
          in Go_ programming language.
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/Go-brown-side.sh-600x600.png
:adsu: yes


When I write frontend code in Go_ (compiled to JavaScript_ via GopherJS_), I
want to make my code look like JavaScript_ code as much as possible. For
example, I want to getElementById_ in Go_ as follows:

.. code-block:: go

  element := Document.GetElementById("elmId")

So I write a Go_/GopherJS_ library [1]_ to realize my idea:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  var Document = js.Global.Get("document")

  func (o *js.Object) GetElementById(id string) *js.Object {
  	return o.Call("getElementById", id)
  }

When I compile the code, I get the following error message:

.. code-block:: go

  panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident [recovered]
  	panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident

.. adsu:: 2

After some googling [2]_ [3]_ [4]_, I found that it is impossible to add method
to existing type in external package. But good news is that I can still do
something similar as follows:

.. code-block:: go

  type Object struct {
  	*js.Object
  }

  var Document = &Object{js.Global.Get("document")}

  func (o *Object) GetElementById(id string) *Object {
  	return &Object{o.Call("getElementById", id)}
  }

Here we wrap ``*js.Object`` in a new defined type *Object* in our package [1]_.
And we can getElementById_ in Go_ as follows:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  element := Document.GetElementById("elmId")

.. adsu:: 3

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.5``

----

References:

.. [1] `GitHub - siongui/godom: Make DOM manipultation in Go as similar to JavaScript as possible. (via GopherJS) <https://github.com/siongui/godom>`_

.. [2] | `panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident - Google search <https://www.google.com/search?q=panic:+interface+conversion:+ast.Expr+is+*ast.SelectorExpr,+not+*ast.Ident>`_
       | `panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident - DuckDuckGo search <https://duckduckgo.com/?q=panic:+interface+conversion:+ast.Expr+is+*ast.SelectorExpr,+not+*ast.Ident>`_
       | `panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident - Ecosia search <https://www.ecosia.org/search?q=panic:+interface+conversion:+ast.Expr+is+*ast.SelectorExpr,+not+*ast.Ident>`_
       | `panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident - Bing search <https://www.bing.com/search?q=panic:+interface+conversion:+ast.Expr+is+*ast.SelectorExpr,+not+*ast.Ident>`_
       | `panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident - Yahoo search <https://search.yahoo.com/search?p=panic:+interface+conversion:+ast.Expr+is+*ast.SelectorExpr,+not+*ast.Ident>`_
       | `panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident - Baidu search <https://www.baidu.com/s?wd=panic:+interface+conversion:+ast.Expr+is+*ast.SelectorExpr,+not+*ast.Ident>`_
       | `panic: interface conversion: ast.Expr is *ast.SelectorExpr, not *ast.Ident - Yandex search <https://www.yandex.com/search/?text=panic:+interface+conversion:+ast.Expr+is+*ast.SelectorExpr,+not+*ast.Ident>`_
.. adsu:: 4
.. [3] | `add a method to an external package - Google search <https://www.google.com/search?q=add+a+method+to+an+external+package>`_
       | `add a method to an external package - DuckDuckGo search <https://duckduckgo.com/?q=add+a+method+to+an+external+package>`_
       | `add a method to an external package - Ecosia search <https://www.ecosia.org/search?q=add+a+method+to+an+external+package>`_
       | `add a method to an external package - Bing search <https://www.bing.com/search?q=add+a+method+to+an+external+package>`_
       | `add a method to an external package - Yahoo search <https://search.yahoo.com/search?p=add+a+method+to+an+external+package>`_
       | `add a method to an external package - Baidu search <https://www.baidu.com/s?wd=add+a+method+to+an+external+package>`_
       | `add a method to an external package - Yandex search <https://www.yandex.com/search/?text=add+a+method+to+an+external+package>`_

.. [4] `How to add new methods to an existing type in go? - Stack Overflow <http://stackoverflow.com/a/28800807>`_

.. [5] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _Go playground: https://play.golang.org/
.. _getElementById: https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
