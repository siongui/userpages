[GopherJS] Test if an Element Contains a Class
##############################################

:date: 2017-01-15T06:10+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, CSS, DOM
:category: GopherJS
:summary: Test if an element contains a class via GopherJS_.


.. code-block:: go

  func IsContainClass(elm *js.Object, class string) bool {
  	return elm.Get("classList").Call("contains", class).Bool()
  }

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `javascript check class exists - Google search <https://www.google.com/search?q=javascript+check+class+exists>`_

       `javascript - Test if an element contains a class? - Stack Overflow <http://stackoverflow.com/a/5898748>`_

.. [3] `GitHub - siongui/gopherjs-utils: useful collections of utilites (functions) for front end (browser) development via GopherJS <https://github.com/siongui/gopherjs-utils>`_

.. _GopherJS: http://www.gopherjs.org/

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
