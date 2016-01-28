[Golang] GopherJS Synonyms with JavaScript
##########################################

:date: 2016-01-29T03:25+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM, JavaScript
:category: GopherJS
:summary: Synonyms - Go_/GopherJS_ idioms and snippets translated to JavaScript_


Inspired by [4]_, Golang_/GopherJS_ (with / without
`GopherJS bindings for the JavaScript DOM APIs`_) synonyms with JavaScript_.


.. list-table:: JavaScript_ vs GopherJS_ vs GopherJS with `DOM binding`_
   :header-rows: 1
   :class: table-syntax-diff

   * - JavaScript_
     - GopherJS_
     - GopherJS with `DOM binding`_

   * - .. code-block:: javascript

         // no need to import anything

     - .. code-block:: go

         import "github.com/gopherjs/gopherjs/js"

     - .. code-block:: go

         import "honnef.co/go/js/dom"

   * - The window_ object

       .. code-block:: javascript

         window

     - type Object_

       .. code-block:: go

         js.Global

     - `GetWindow()`_ function

       .. code-block:: go

         dom.GetWindow()


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``,
``Chromium Version 48.0.2564.82 Ubuntu 15.10 (64-bit)``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)

.. [4] `Synonyms - Dart, JavaScript, C#, Python | Dart <https://www.dartlang.org/docs/synonyms/>`_

.. _GopherJS: http://www.gopherjs.org/
.. _DOM binding: https://godoc.org/honnef.co/go/js/dom
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _window: http://www.w3schools.com/jsref/obj_window.asp
.. _Object: https://godoc.org/github.com/gopherjs/gopherjs/js#Object
.. _GetWindow(): https://godoc.org/honnef.co/go/js/dom#GetWindow

.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
