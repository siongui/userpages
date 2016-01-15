[Golang] GopherJS DOM Example - getElementById and Set innerHTML
################################################################

:date: 2016-01-10T01:23+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript
:category: Go
:summary: Run Golang_ program in your browser by GopherJS_. Show how to write a
          Go_ program to do DOM_ manipulation by example. Use getElementById_ to
          access DOM_ element and set innerHTML_ of the element.

Introduction
++++++++++++

It is really cool to run Go_ code in the browser. GopherJS_ is a compiler from
Go_ to JavaScript_, which makes this possible. In this post, we will give a very
simple example of DOM_ manipulation in Go_ program. This example uses
getElementById_ to access DOM_ element and set innerHTML_ of the element.

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

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/demo/index.html

We will access the *div* element whose *id* is *foo*. Then add *Hello World*
text to the *div*. Now write a Go_ program to manipulate DOM_:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/demo/dom.go

It is very easy and intuitive. Compile the Go_ code to JavaScript_:

.. code-block:: bash

  $ gopherjs build dom.go -o dom.js

Put *dom.js* together with the *index.html* in the same directory and open the
*index.html* with your browser. You will see *Hello World* in the browser
window.

.. .. show_github_file:: siongui userpages content/code/gopherjs-dom/Makefile


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.2``.

----

GopherJS_ DOM_ Example series
+++++++++++++++++++++++++++++

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

.. [5] `GopherJSの紹介 - GolangRdyJp <http://golang.rdy.jp/2015/10/15/gopherjs/>`_

.. [6] `albrow/gopherjs-live · GitHub <https://github.com/albrow/gopherjs-live>`_
       (Automatic watching and recompiling for gopherjs)

.. [7] `ajhager/srvi · GitHub <https://github.com/ajhager/srvi>`_
       (Quickly build, serve, run, and refresh your GopherJS programs)

.. [8] `cmd/gopherjs_serve_html at master · shurcooL/cmd · GitHub <https://github.com/shurcooL/cmd/tree/master/gopherjs_serve_html>`_

.. [9] `Add "gopherjs serve" command · Issue #121 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/121>`_

.. [10] `It's easy to get an infinite loop with the watch flag · Issue #212 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/212>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _getElementById: http://www.w3schools.com/jsref/met_doc_getelementbyid.asp
.. _innerHTML: http://www.w3schools.com/jsref/prop_html_innerhtml.asp
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
