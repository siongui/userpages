[Golang] Online Input Method (Pāli) by GopherJS
###############################################

:date: 2016-01-12T20:14+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, html, Pāli Input Method , IME, Keyboard Event, Web application, web
:category: GopherJS
:summary: Online Pali_ (Pāli, Pāḷi) Input Method using Go_ programming language,
          compiled to JavaScript_ by GopherJS_. (GopherJS DOM_ example)
:adsu: yes


.. rubric:: `Demo <https://siongui.github.io/go-online-input-method-pali/>`_ (`Fork me on Github <https://github.com/siongui/go-online-input-method-pali>`_)
    :class: align-center

Inspired by `palipad <https://code.google.com/p/palipad/>`_
(`demo of palipad <http://palipad.googlecode.com/git/palipad.html>`_),
I write a Golang_ webapp which help users input `Pāli`_ language online.
The implementation is simple. Create an `input:text`_ element first, and then
listen to the `onKeyup event`_ of the `input:text`_ element. The last two
characters of the `input:text`_ element will be replaced according to the
following rule. This application is an example of DOM manipulation by GopherJS.
See `GopherJS DOM Example series`_ for more details.

+------+-----+
| Type | For |
+======+=====+
|  AA  |  Ā  |
+------+-----+
|  aa  |  ā  |
+------+-----+
|  II  |  Ī  |
+------+-----+
|  ii  |  ī  |
+------+-----+
|  UU  |  Ū  |
+------+-----+
|  uu  |  ū  |
+------+-----+
|  "N  |  Ṅ  |
+------+-----+
|  "n  |  ṅ  |
+------+-----+
|  .M  |  Ṃ  |
+------+-----+
|  .m  |  ṃ  |
+------+-----+
|  ~N  |  Ñ  |
+------+-----+
|  ~n  |  ñ  |
+------+-----+
|  .T  |  Ṭ  |
+------+-----+
|  .t  |  ṭ  |
+------+-----+
|  .D  |  Ḍ  |
+------+-----+
|  .d  |  ḍ  |
+------+-----+
|  .N  |  Ṇ  |
+------+-----+
|  .n  |  ṇ  |
+------+-----+
|  .L  |  Ḷ  |
+------+-----+
|  .l  |  ḷ  |
+------+-----+

Source Code for Demo (*HTML*):

.. show_github_file:: siongui go-online-input-method-pali src/index.html

Source Code for Demo (*Go*):

.. show_github_file:: siongui go-online-input-method-pali src/pali.go

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

- `[Golang] Online Snake Game by GopherJS <{filename}../13/go-online-snake-game-by-gopherjs%en.rst>`_

- `[Golang] GopherJS DOM Example - Hide Element by display:none <{filename}../13/gopherjs-dom-example-hide-element-by-display-none%en.rst>`_

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

.. [5] `[Dart] Online Input Method - Pali (Pāli, Pāḷi) <{filename}../../../2015/02/23/dart-online-input-method-pali%en.rst>`_


.. _Pali: https://en.wikipedia.org/wiki/Pali
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _Pāli: http://en.wikipedia.org/wiki/Pali
.. _input\:text: http://www.w3schools.com/tags/tag_input.asp
.. _onKeyup event: http://www.w3schools.com/jsref/event_onkeyup.asp

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
