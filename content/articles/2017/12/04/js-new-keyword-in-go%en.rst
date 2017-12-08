JavaScript new Keyword in Go
############################

:date: 2017-12-08T23:04+08:00
:modified: 2017-12-09T00:41+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Show how to use JavaScript *new* keyword in Go/GopherJS.
:adsu: yes


JavaScript *new* keyword are used very often. For example, `Date()`_,
`XMLHttpRequest()`_, or `WebSocket()`_ need to be *newed* before being used.
We will use *Date()* as example and show you how to write equivalent code in
Go/GopherJS.


JavaScript
++++++++++

To print the date info in JavaScript, you can do the following with *new*
keyword:

.. code-block:: javascript

  var d = new Date();
  console.log(d);

You can try the `code online in w3schools`_. [1]_


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  d := js.Global.Get("Date").New()
  println(d)

You can try the code online:

.. rubric:: `Run Code on GopherJS Playground <https://gopherjs.github.io/playground/#/svZwXAls_H>`_
   :class: align-center

.. adsu:: 2


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  d := Window.Get("Date").New()
  println(d)


For more exaples for JavaScript *new* keyword in Go, see [2]_, [3]_, [4]_.


Chain Dots
++++++++++

From `the comment`_ of `@pbrown12303`_, if you have to *new* the following chain
dots in JavaScript:

.. code-block:: javascript

  var x = new joint.dia.Graph;

The following code in Go/GopherJS will not work:

.. code-block:: go

  x := js.Global.Get("joint.dia.Graph").New()

The correct syntax in Go/GopherJS is:

.. code-block:: go

  x := js.Global.Get("joint").Get("dia").Get("Graph").New()


.. adsu:: 3

----

References:

.. [1] `JavaScript Dates <https://www.w3schools.com/js/js_dates.asp>`_
.. [2] `GopherJS XMLHttpRequest (XHR) and MakeFunc Example <{filename}../../../2016/02/18/gopherjs-XMLHttpRequest-XHR-and-MakeFunc-example%en.rst>`_
.. [3] `[GopherJS] WebSocket Client Example With Echo Server <{filename}../../05/18/go-websocket-client-example-with-echo-server%en.rst>`_
.. [4] `[Golang] Access HTTP Request Header by XHR getAllResponseHeaders() <{filename}../../../2016/01/25/go-http-request-header-by-xhr-getAllResponseHeaders%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _window: http://www.w3schools.com/jsref/obj_window.asp
.. _Object: https://godoc.org/github.com/gopherjs/gopherjs/js#Object
.. _document: http://www.w3schools.com/jsref/dom_obj_document.asp
.. _godom: https://github.com/siongui/godom
.. _Date(): https://www.google.com/search?q=date+javascript
.. _WebSocket(): https://www.google.com/search?q=WebSocket+javascript
.. _XMLHttpRequest(): https://www.google.com/search?q=XMLHttpRequest+javascript
.. _code online in w3schools: https://www.w3schools.com/js/tryit.asp?filename=tryjs_date_current
.. _the comment: https://github.com/siongui/userpages/issues/4#issuecomment-350292186
.. _@pbrown12303: https://github.com/pbrown12303
