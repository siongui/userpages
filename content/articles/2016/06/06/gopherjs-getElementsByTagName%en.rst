[GopherJS] getElementsByTagName
###############################

:date: 2016-06-06T03:06+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: Use getElementsByTagName_ to access head_ element via GopherJS_.
:adsu: yes


Use getElementsByTagName_ to access head_ element via GopherJS_.

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  head := js.Global.Get("document").Call("getElementsByTagName", "head").Call("item", 0)

You can also use querySelector_ [2]_ to access head_ element:

.. code-block:: go

  head := js.Global.Get("document").Call("querySelector", "head")

----

.. adsu:: 2

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `[Golang] querySelectorAll and querySelector Example by GopherJS <{filename}../../02/14/go-querySelectorAll-querySelector-by-gopherjs%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _CSS: https://www.google.com/search?q=CSS
.. _head: http://www.w3schools.com/html/html_head.asp
.. _getElementsByTagName: https://www.google.com/search?q=getElementsByTagName
.. _querySelector: https://www.google.com/search?q=querySelector

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
