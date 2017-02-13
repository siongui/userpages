[GopherJS] Get head Element in HTML Document
############################################

:date: 2016-03-13T20:17+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, html
:category: GopherJS
:summary: Access head_ element in HTML document_ via GopherJS_.
:adsu: yes

Access head_ element in HTML document_ via getElementsByTagName_ and GopherJS_:

.. code-block:: go

  head := js.Global.Get("document").Call("getElementsByTagName", "head").Call("item", 0)

equivalent to the following JavaScript code:

.. code-block:: javascript

  var head = document.getElementsByTagName("head")[0];

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

.. adsu:: 2

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `javascript - document.head v. document.getElementsByTagName("head")[0] - Stack Overflow <http://stackoverflow.com/questions/16204756/document-head-v-document-getelementsbytagnamehead0>`_

.. [3] `Element.getElementsByTagName() - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName>`_

.. [4] `HTMLCollection - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _head: http://www.w3schools.com/html/html_head.asp
.. _document: http://www.w3schools.com/jsref/dom_obj_document.asp
.. _getElementsByTagName: https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
