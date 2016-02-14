[Golang] querySelectorAll and querySelector Example by GopherJS
###############################################################

:date: 2016-02-14T07:58+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript
:category: GopherJS
:summary: Go_ querySelectorAll_ and querySelector_ Example by GopherJS_.

Golang_ querySelectorAll_ and querySelector_ Example by GopherJS_.

querySelector_
++++++++++++++

The querySelector() method returns a DOM element_ object:

.. code-block:: go

  // document object
  d := js.Global.Get("document")
  // return element whose css class is setting-menu
  element := d.Call("querySelector", ".setting-menu")
  // do something with the element ...

querySelectorAll_
+++++++++++++++++

The querySelectorAll() method returns a NodeList_:

.. code-block:: go

  // document object
  d := js.Global.Get("document")
  // return a nodelist in which the elements whose css class are about-link
  nodeList := d.Call("querySelectorAll", ".about-link")
  // access individual element
  length := nodeList.Get("length").Int()
  for i := 0; i < length; i++ {
          // get i-th element in nodelist
          element := nodeList.Call("item", i)
          // do something with the element ...
  }

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _querySelectorAll: https://www.google.com/search?q=querySelectorAll
.. _querySelector: https://www.google.com/search?q=querySelector
.. _element: https://developer.mozilla.org/en-US/docs/Web/API/element
.. _NodeList: https://developer.mozilla.org/en-US/docs/Web/API/NodeList

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
