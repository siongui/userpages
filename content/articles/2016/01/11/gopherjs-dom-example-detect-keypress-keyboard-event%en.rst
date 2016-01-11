[Golang] GopherJS DOM Example - Detect Keypress (Keyboard Event)
################################################################

:date: 2016-01-11T04:48+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, Keyboard Event
:category: Go
:summary: Run Golang_ program in your browser by GopherJS_. Show how to write a
          Go_ program to do DOM_ manipulation by example. Detect user keypress
          by registering `keyboard event`_ handler and print out the keycode_ of
          the keyboard event.

Introduction
++++++++++++

It is really cool to run Go_ code in the browser. GopherJS_ is a compiler from
Go_ to JavaScript_, which makes this possible. In this post, we will give an
example of DOM_ manipulation in Go_ program. This example shows how to detect
user keypress. addEventListener_ function is used to attach an onkeydown_
(`keyboard event`_) handler to window_ object. In the event handler, the
keycode_ of the keyboard event will be printed out.

Install GopherJS_ and DOM_ bindings
+++++++++++++++++++++++++++++++++++

Run the following command to install GopherJS_ and
`GopherJS bindings for the JavaScript DOM APIs`_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs
  $ go get -u honnef.co/go/js/dom

Source Code
+++++++++++

First we write a simple HTML for our demo:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/keyevent/index.html

We attach an onkeydown_ event handler to the window_ object. When users press any key,
the keycode_ of the `keyboard event`_ is printed out.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/keyevent/keycode.go

The above code is very similar to JavaScript counterpart. Note that *dom.Event*
is casted to *dom.KeyboardEvent* by `type assertions`_. Now compile the Go_ code
to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build keycode.go -o demo.js

Put *demo.js* together with the *index.html* in the same directory and open the
*index.html* with your browser. Press any key and you will see the keycode of
the keypress.

.. .. show_github_file:: siongui userpages content/code/gopherjs-dom/Makefile


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.2``.

----

GopherJS_ DOM_ Example series
+++++++++++++++++++++++++++++

- `[Golang] GopherJS DOM Example - getElementById and Set innerHTML <{filename}../10/gopherjs-dom-example-getElementById-innerHTML%en.rst>`_

- `[Golang] GopherJS DOM Example - Event Binding (addEventListener) <{filename}gopherjs-dom-example-event-binding-addEventListener%en.rst>`_

- `[Golang] GopherJS DOM Example - Access Input Element Value <{filename}gopherjs-dom-example-access-input-element-value%en.rst>`_

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)

.. [4] `Getting Started with GopherJS <https://www.hakkalabs.co/articles/getting-started-gopherjs>`_

.. [5] Google Search `golang struct casting <https://www.google.com/search?q=golang+struct+casting>`__

.. [6] `go - Casting back to more specialised interface - Stack Overflow <http://stackoverflow.com/questions/4799905/casting-back-to-more-specialised-interface>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _addEventListener: http://www.w3schools.com/jsref/met_element_addeventlistener.asp
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _keyboard event: http://www.w3schools.com/jsref/dom_obj_event.asp
.. _keycode: http://www.w3schools.com/jsref/event_key_keycode.asp
.. _window: http://www.w3schools.com/jsref/obj_window.asp
.. _onkeydown: http://www.w3schools.com/jsref/event_onkeydown.asp
.. _type assertions: https://golang.org/ref/spec#Type_assertions


.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
