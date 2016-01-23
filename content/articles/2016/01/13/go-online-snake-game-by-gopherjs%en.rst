[Golang] Online Snake Game by GopherJS
######################################

:date: 2016-01-13T19:22+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, html, Web application, web, Goroutine, Keyboard Event
:category: GopherJS
:summary: Online `snake game`_ using Go_ programming language, compiled to
          JavaScript_ by GopherJS_. (GopherJS DOM_ example)


.. rubric:: `Demo <https://siongui.github.io/go-online-snake-game/>`__ (`Fork me on Github <https://github.com/siongui/go-online-snake-game>`__)
    :class: align-center

In `GopherJS DOM Example series`_, we show a lot of examples of DOM manipulation
by GopherJS_ and its `DOM binding`_. Now we will use what we learn to write
a simple `snake game`_. I wrote an online snake game in Dart_ (see [5]_) before.
This Golang_ snake game is ported from the Dart version. The DOM APIs provided
by GopherJS and its DOM bindings are basically the same as JavaScript DOM APIs,
except it is re-written in idiomatic Go. This is a good exercise to make you
familiar with `DOM manipulation`_ in Go_ programming language.


Source Code for Demo (*HTML*):

.. show_github_file:: siongui go-online-snake-game src/index.html

Source Code for Demo (*CSS*):

.. show_github_file:: siongui go-online-snake-game src/style.css

Source Code for Demo (*Go*):

.. show_github_file:: siongui go-online-snake-game src/snake.go


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.2``.

----

GopherJS_ DOM_ Example series
+++++++++++++++++++++++++++++

- `[Golang] GopherJS DOM Example - getElementById and Set innerHTML <{filename}../10/gopherjs-dom-example-getElementById-innerHTML%en.rst>`_

- `[Golang] GopherJS DOM Example - Event Binding (addEventListener) <{filename}../11/gopherjs-dom-example-event-binding-addEventListener%en.rst>`_

- `[Golang] GopherJS DOM Example - Detect Keypress (Keyboard Event) <{filename}../11/gopherjs-dom-example-detect-keypress-keyboard-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Access Input Element Value <{filename}../11/gopherjs-dom-example-access-input-element-value%en.rst>`_

- `[Golang] GopherJS DOM Example - Access HTML Data Attribute <{filename}../12/gopherjs-dom-example-access-html-data-attribute%en.rst>`_

- `[Golang] Online Input Method (Pāli) by GopherJS <{filename}../12/go-online-input-method-pali-by-gopherjs%en.rst>`_

- `[Golang] GopherJS DOM Example - Hide Element by display:none <{filename}gopherjs-dom-example-hide-element-by-display-none%en.rst>`_

- `[Golang] GopherJS DOM Example - Create and Append Element <{filename}../14/gopherjs-dom-example-create-and-append-element%en.rst>`_

- `[Golang] GopherJS DOM Example - Play Sound on Click Event <{filename}../15/gopherjs-dom-example-play-sound-onclick-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Toggle (Play/Pause) Sound on Click Event <{filename}../15/gopherjs-dom-example-toggle-sound-onclick-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Dropdown Menu <{filename}../16/gopherjs-dom-example-dropdown-menu%en.rst>`_

- `[Golang] Draggable (Movable) Element by GopherJS <{filename}../17/go-draggable-movable-element-by-gopherjs%en.rst>`_

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

.. [5] `siongui/dart-snake-game · GitHub <https://github.com/siongui/dart-snake-game>`_
       (`demo <https://siongui.github.io/dart-snake-game/>`__)


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _snake game: https://www.google.com/search?q=snake+game
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _DOM binding: https://godoc.org/honnef.co/go/js/dom
.. _Dart: https://www.dartlang.org/
.. _DOM manipulation: https://www.google.com/search?q=DOM+manipulation

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
