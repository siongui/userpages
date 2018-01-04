Event Binding - addEventListener in Go
######################################

:date: 2017-12-11T22:41+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Event binding in Go/GopherJS - Register an event handler to the
          specified DOM element via *addEventListener* method.
:og_image: https://cdn.tutsplus.com/active/uploads/legacy/flashtuts/074_EventListenersBasics/1.jpg
:adsu: yes


This post shows how to do event binding in Go/GopherJS on your browser, i.e.,
tell the browser to do some actions (i.e., handle the event) when some event
occurs. We will use addEventListener_ method to register the event handler to
the event of the specified DOM element.
The full code example of this post is `on my GitHub`_.

.. contents:: **Table of Content**

addEventListener() method
=========================

Assume we have the following *div* element in our HTML

.. code-block:: html

  <div id="foo">Click Here</div>

We will attach a *click* event handler to the *div* element, and when users
click on the div, the text content of the *div* element will be changed to
*I am clicked*.


JavaScript
++++++++++

.. code-block:: javascript

  var f = document.querySelector("#foo");

  f.addEventListener("click", function(e) {
    f.textContent = "I am clicked";
  });


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  f := js.Global.Get("document").Call("querySelector", "#foo")

  f.Call("addEventListener", "click", func(event *js.Object) {
  	f.Set("textContent", "I am clicked")
  })


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  f := Document.QuerySelector("#foo")

  f.AddEventListener("click", func(e Event) {
  	f.SetTextContent("I am clicked")
  })

.. adsu:: 2

You can add third argument (*useCapture*, a boolean value) to the
*addEventListener()* method. The default is *false* is not specified.

----

References:

.. [1] `[Golang] GopherJS Synonyms with JavaScript <{filename}../../../2016/01/29/go-gopherjs-synonyms-with-javascript%en.rst>`_
.. [2] `[GopherJS] Register Event Handler (Event Binding) <{filename}../../../2016/06/07/gopherjs-register-event-handler-event-binding%en.rst>`_
.. [3] `[Golang] GopherJS DOM Example - Event Binding (addEventListener) <{filename}../../../2016/01/11/gopherjs-dom-example-event-binding-addEventListener%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _addEventListener: https://www.google.com/search?q=addEventListener
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/005-addEventListener
