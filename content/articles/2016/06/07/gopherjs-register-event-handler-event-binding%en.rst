[GopherJS] Register Event Handler (Event Binding)
#################################################

:date: 2016-06-07T21:08+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: Two ways to register event handler via GopherJS_:
          Set element on-Event property or addEventListener_.


.. contents:: Two ways to register event handler, use `onclick event`_ as example:

Set onclick_ property
+++++++++++++++++++++

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  element := js.Global.Get("document").Call("getElementById", "foo")
  element.Set("onclick", func(event *js.Object) {
  	/* do something to handle click event */
  })


addEventListener_
+++++++++++++++++

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  element := js.Global.Get("document").Call("getElementById", "foo")
  element.Call("addEventListener", "click", func(event *js.Object) {
  	/* do something to handle click event */
  })

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `[Golang] GopherJS DOM Example - Event Binding (addEventListener) <{filename}../../01/11/gopherjs-dom-example-event-binding-addEventListener%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _CSS: https://www.google.com/search?q=CSS
.. _head: http://www.w3schools.com/html/html_head.asp
.. _getElementsByTagName: https://www.google.com/search?q=getElementsByTagName
.. _querySelector: https://www.google.com/search?q=querySelector
.. _addEventListener: http://www.w3schools.com/jsref/met_element_addeventlistener.asp
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _onclick event: http://www.w3schools.com/jsref/event_onclick.asp
.. _onclick: https://developer.mozilla.org/en-US/docs/Web/API/GlobalEventHandlers/onclick

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
