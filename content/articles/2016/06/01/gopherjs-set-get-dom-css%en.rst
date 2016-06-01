[GopherJS] Set/Get DOM CSS
##########################

:date: 2016-06-01T22:22+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, CSS, DOM
:category: GopherJS
:summary: Example - Set or Get CSS_ property of DOM_ element via GopherJS_.


Example - Set or Get CSS_ property of DOM_ element via GopherJS_.

GopherJS_ Set DOM_ CSS
++++++++++++++++++++++

.. code-block:: go

  element := js.Global.Get("document").Call("getElementById", "foo")
  element.Get("style").Set("left", "50px")


GopherJS_ Get DOM_ CSS
++++++++++++++++++++++

.. code-block:: go

  element := js.Global.Get("document").Call("getElementById", "foo")
  // cast as int
  element.Get("style").Get("offsetWidth").Int()
  // cast as string
  element.Get("style").Get("offsetWidth").String()

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `javascript set css - Google search <https://www.google.com/search?q=javascript+set+css>`_

       `JavaScript DOM CSS <http://www.w3schools.com/js/js_htmldom_css.asp>`_


.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _CSS: https://www.google.com/search?q=CSS

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
