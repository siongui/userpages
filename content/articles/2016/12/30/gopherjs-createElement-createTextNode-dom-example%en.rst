[GopherJS] createElement and createTextNode DOM Example
#######################################################

:date: 2016-12-30T21:30+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: DOM_ Example of createElement_ and createTextNode_ via GopherJS_.
:adsu: yes


DOM_ Example of createElement_ and createTextNode_ via GopherJS_.

The following code creates DOM_ equivalent to
``<div><span>Hello</span> World!</div>``.

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  func createElementAndTextNode() *js.Object {
  	div := js.Global.Get("document").Call("createElement", "div")

  	span := js.Global.Get("document").Call("createElement", "span")
  	span.Set("textContent", "Hello")
  	div.Call("appendChild", span)

  	text := js.Global.Get("document").Call("createTextNode", " World!")
  	div.Call("appendChild", text)

  	return div
  }

.. adsu:: 2

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.4``
- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Document.createElement() - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement>`_

       `Document.createTextNode() - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Document/createTextNode>`_


.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _createElement: https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
.. _createTextNode: https://developer.mozilla.org/en-US/docs/Web/API/Document/createTextNode

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
