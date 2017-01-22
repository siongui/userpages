[GopherJS] Access Child Nodes of DOM Element
############################################

:date: 2016-06-11T20:39+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: Access child nodes (childNodes_) of DOM_ element via GopherJS_.
:adsu: yes


Access child nodes (childNodes_) of DOM_ element via GopherJS_.

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  element := js.Global.Get("document").Call("getElementById", "foo")
  childNodesList := element.Get("childNodes")
  length := childNodesList.Get("length").Int()
  for i := 0; i < length; i++ {
  	// get i-th child
  	child := childNodesList.Call("item", i)
  	// do somthing to the child element here
  }

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `childNodes - Google search <https://www.google.com/search?q=childNodes>`_

       `Node.childNodes - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes>`_

.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _childNodes: https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
