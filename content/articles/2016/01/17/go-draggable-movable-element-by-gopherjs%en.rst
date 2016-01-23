[Golang] Draggable (Movable) Element by GopherJS
################################################

:date: 2016-01-17T23:35+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, draggable
:category: GopherJS
:summary: Run Golang_ program in your browser by GopherJS_. Show how to write a
          Go_ program to do DOM_ manipulation by example. This example show how
          to write a draggable/movable element by GopherJS_ and its
          `DOM binding`_.

Introduction
++++++++++++

It is really cool to run Go_ code in the browser. GopherJS_ is a compiler from
Go_ to JavaScript_, which makes this possible.
In this post, we will give an example of DOM_ manipulation in Go_ program.
This example shows how to write a draggable/movable HTML element. This example
is ported from Dart_ draggable/movable element [5]_.
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

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/draggable/index.html

The basic idea is that we bind onmousedown_ event handler to the draggable
element. In the onmousedown_ event handler, we bind onmousemove_ and onmouseup_
event handler to document_ object. In the onmousemove_ event handler, we
calculate the movement and change the *CSS* property of the draggable element to
make it move. In the onmouseup_ event handler, we remove the onmousemove_ and
onmouseup_ event handler from document_ object.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/draggable/draggable.go

Compile the Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build draggable.go -o demo.js

Put *demo.js* together with the *index.html* in the same directory. Open the
*index.html* with your browser, and start to drag the element in the browser!

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

- `[Golang] GopherJS DOM Example - Create and Append Element <{filename}../14/gopherjs-dom-example-create-and-append-element%en.rst>`_

- `[Golang] GopherJS DOM Example - Play Sound on Click Event <{filename}../15/gopherjs-dom-example-play-sound-onclick-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Toggle (Play/Pause) Sound on Click Event <{filename}../15/gopherjs-dom-example-toggle-sound-onclick-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Dropdown Menu <{filename}../16/gopherjs-dom-example-dropdown-menu%en.rst>`_

- `[Golang] Toggle (Show/Hide) HTML Element by GopherJS <{filename}../18/go-toggle-show-hide-element-by-gopherjs%en.rst>`_

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

.. [5] `[Dart] Draggable (Movable) Element <{filename}../../../2015/02/17/dart-draggable-movable-element%en.rst>`_

.. [6] `[AngularJS] Draggable (Movable) Element <{filename}../../../2013/04/04/angularjs-draggable-movable-element%en.rst>`_

.. [7] `JavaScript Drag and Drop (Draggable, Movable) Element without External Library <{filename}../../../2012/07/13/javascript-drag-and-drop-draggable-movable-element%en.rst>`_

.. [8] `golang function inside function <https://www.google.com/search?q=golang+function+inside+function>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _HTML: http://www.w3schools.com/html/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _DOM binding: https://godoc.org/honnef.co/go/js/dom
.. _document: http://www.w3schools.com/jsref/dom_obj_document.asp
.. _Dart: https://www.dartlang.org/
.. _onmousedown: http://www.w3schools.com/jsref/event_onmousedown.asp
.. _onmouseup: http://www.w3schools.com/jsref/event_onmouseup.asp
.. _onmousemove: http://www.w3schools.com/jsref/event_onmousemove.asp

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
