[Golang] Remove All Child Nodes of a DOM Element by GopherJS
############################################################

:date: 2016-01-31T21:21+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM, JavaScript
:category: GopherJS
:summary: Go_ programming language - Remove all children of a DOM_ element by
          GopherJS_.
:adsu: yes


In JavaScript_, we can remove all child nodes of a DOM element [4]_ by:

.. code-block:: javascript

  while (elm.hasChildNodes()) {
    elm.removeChild(elm.lastChild);
  }

In Golang_, we can do the same by GopherJS_:

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  for elm.Call("hasChildNodes").Bool() {
          elm.Call("removeChild", elm.Get("lastChild"))
  }

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``,
``Chromium Version 48.0.2564.82 Ubuntu 15.10 (64-bit)``.

----

.. adsu:: 2

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)

.. [4] `JavaScript Remove All Children of a DOM Element <{filename}../../../2012/09/26/javascript-remove-all-children-of-dom-element%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
