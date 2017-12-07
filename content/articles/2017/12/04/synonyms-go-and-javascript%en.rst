Synonyms - Go and JavaScript
############################

:date: 2017-12-07T21:47+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM, JavaScript,
       Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Synonyms - Go_/GopherJS_ idioms and snippets translated to JavaScript_
:adsu: yes


Inspired by [4]_, This post provides synonyms of Golang_ and JavaScript_ in
frontend programming, and served as references for developers who already have
basic understanding of JavaScript and GopherJS_/godom_.

.. contents:: **Table of Content**


----


import packages
+++++++++++++++

**JavaScript**: no need to import any packages

**GopherJS**

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

**GopherJS + godom**

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )


----


window_ object
++++++++++++++

**JavaScript**

.. code-block:: javascript

  window

**GopherJS**

.. code-block:: go

  js.Global

**GopherJS + godom**

.. code-block:: go

  Window


----


`alert()`_ method
+++++++++++++++++

**JavaScript**

.. code-block:: javascript

  alert("Hello World");

**GopherJS**

.. code-block:: go

  js.Global.Call("alert", "Hello World")

**GopherJS + godom**

.. code-block:: go

  Window.Alert("Hello World")


----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)
.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_
.. [3] `GitHub - siongui/godom: Make DOM manipulation in Go as similar to JavaScript as possible. (via GopherJS) <https://github.com/siongui/godom>`_
.. [4] `Synonyms - Dart, JavaScript, C#, Python | Dart <https://www.dartlang.org/resources/synonyms>`_
.. [5] `[Golang] GopherJS Synonyms with JavaScript <{filename}../../../2016/01/29/go-gopherjs-synonyms-with-javascript%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _DOM binding: https://godoc.org/honnef.co/go/js/dom
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _window: http://www.w3schools.com/jsref/obj_window.asp
.. _Object: https://godoc.org/github.com/gopherjs/gopherjs/js#Object
.. _GetWindow(): https://godoc.org/honnef.co/go/js/dom#GetWindow
.. _document: http://www.w3schools.com/jsref/dom_obj_document.asp
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _alert(): http://www.w3schools.com/jsref/met_win_alert.asp
.. _navigator: https://developer.mozilla.org/en-US/docs/Web/API/Navigator
.. _NavigatorLanguage: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage
.. _getElementById(): https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
.. _innerHTML: http://www.w3schools.com/jsref/prop_html_innerhtml.asp
.. _textContent: http://www.w3schools.com/jsref/prop_node_textcontent.asp
.. _addEventListener(): https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
.. _Remove all child nodes: https://www.google.com/search?q=javascript+remove+all+child+nodes
.. _createElement: https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
.. _createTextNode: https://developer.mozilla.org/en-US/docs/Web/API/Document/createTextNode
.. _location: http://www.w3schools.com/jsref/obj_location.asp
.. _querySelector: https://www.google.com/search?q=querySelector
.. _querySelectorAll: https://www.google.com/search?q=querySelectorAll
.. _NodeList: https://developer.mozilla.org/en-US/docs/Web/API/NodeList
.. _godom: https://github.com/siongui/godom

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
