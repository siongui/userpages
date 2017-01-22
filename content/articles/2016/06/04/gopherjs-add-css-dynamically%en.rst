[GopherJS] Insert CSS Dynamically
#################################

:date: 2016-06-04T19:56+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, CSS, DOM
:category: GopherJS
:summary: Insert (add, append) CSS_ to head_ element via GopherJS_.
:adsu: yes


Insert (add, append) CSS_ to head_ element via GopherJS_.

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  func appendCSSToHeadElement() {
  	css := `.tooltip {
  		position: absolute;
  		left: -9999px;
  		background-color: #CCFFFF;
  		border-radius: 10px;
  		font-family: Tahoma, Arial, serif;
  		word-wrap: break-word;
  	}`
  	s := js.Global.Get("document").Call("createElement", "style")
  	s.Set("innerHTML", css)
  	// insert style of tooltip at the end of head element
  	js.Global.Get("document").Call("getElementsByTagName", "head").Call("item", 0).Call("appendChild", s)
  }

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `javascript insert css - Google search <https://www.google.com/search?q=javascript+insert+css>`_

.. [3] `[JavaScript] Load CSS Dynamically <{filename}../../../2012/10/10/javascript-load-css-dynamically%en.rst>`_

.. [4] `Load External JavaScript or CSS file Dynamically <{filename}../../../2012/06/18/load-external-javascript-or-css-file-dynamically%en.rst>`_

.. [5] `[GopherJS] Set/Get DOM CSS <{filename}../01/gopherjs-set-get-dom-css%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _CSS: https://www.google.com/search?q=CSS
.. _head: http://www.w3schools.com/html/html_head.asp

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
