[Golang] GopherJS DOM Example - Create and Append Element
#########################################################

:date: 2016-01-14T19:21+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript
:category: Go
:summary: Run Golang_ program in your browser by GopherJS_. Show how to write a
          Go_ program to do DOM_ manipulation by example. This example show how
          to create a new HTML_ element and append it to the DOM tree.

Introduction
++++++++++++

It is really cool to run Go_ code in the browser. GopherJS_ is a compiler from
Go_ to JavaScript_, which makes this possible.
In this post, we will give a simple example of DOM_ manipulation in Go_ program.
This example shows how to create a new HTML_ element by createElement_ method
and append it to the DOM tree by appendChild_ method.
If you are not familiar with basic DOM_ manipulation in Go_, read the
posts in `GopherJS DOM Example series`_ first.

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

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/create/index.html

We will bind a onclick_ event handler to the *div* element whose *id* is *foo*.
When users click the *div* element, A new **div** element will be created and
appended to the *foo div* element.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/create/append.go

Compile the Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build append.go -o demo.js

Put *demo.js* together with the *index.html* in the same directory and open the
*index.html* with your browser. Click the text and a new red text will show up!

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

GopherJS_ DOM_ Example series
+++++++++++++++++++++++++++++

- `[Golang] GopherJS DOM Example - getElementById and Set innerHTML <{filename}../10/gopherjs-dom-example-getElementById-innerHTML%en.rst>`_

- `[Golang] GopherJS DOM Example - Event Binding (addEventListener) <{filename}../11/gopherjs-dom-example-event-binding-addEventListener%en.rst>`_

- `[Golang] GopherJS DOM Example - Detect Keypress (Keyboard Event) <{filename}../11/gopherjs-dom-example-detect-keypress-keyboard-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Access Input Element Value <{filename}../11/gopherjs-dom-example-access-input-element-value%en.rst>`_

- `[Golang] GopherJS DOM Example - Access HTML Data Attribute <{filename}../12/gopherjs-dom-example-access-html-data-attribute%en.rst>`_

- `[Golang] Online Input Method (Pāli) by GopherJS <{filename}../12/go-online-input-method-pali-by-gopherjs%en.rst>`_

- `[Golang] Online Snake Game by GopherJS <{filename}../13/go-online-snake-game-by-gopherjs%en.rst>`_

- `[Golang] GopherJS DOM Example - Hide Element by display:none <{filename}../13/gopherjs-dom-example-hide-element-by-display-none%en.rst>`_

- `[Golang] GopherJS DOM Example - Play Sound on Click Event <{filename}../15/gopherjs-dom-example-play-sound-onclick-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Toggle (Play/Pause) Sound on Click Event <{filename}../15/gopherjs-dom-example-toggle-sound-onclick-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Dropdown Menu <{filename}../16/gopherjs-dom-example-dropdown-menu%en.rst>`_

- `[Golang] Draggable (Movable) Element by GopherJS <{filename}../17/go-draggable-movable-element-by-gopherjs%en.rst>`_

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


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _HTML: http://www.w3schools.com/html/
.. _createElement: http://www.w3schools.com/jsref/met_document_createelement.asp
.. _appendChild: http://www.w3schools.com/jsref/met_node_appendchild.asp
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _onclick: http://www.w3schools.com/jsref/event_onclick.asp
.. _assert the type: https://golang.org/ref/spec#Type_assertions

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
