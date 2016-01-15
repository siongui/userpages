[Golang] GopherJS DOM Example - Access HTML Data Attribute
##########################################################

:date: 2016-01-12T00:05+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, html
:category: Go
:summary: Run Golang_ program in your browser by GopherJS_. Show how to write a
          Go_ program to do DOM_ manipulation by example. This example shows how
          to access HTML_ `data-* attribute`_.

Introduction
++++++++++++

It is really cool to run Go_ code in the browser. GopherJS_ is a compiler from
Go_ to JavaScript_, which makes this possible. In this post, we will give a
simple example of DOM_ manipulation in Go_ program. This example shows how to
access HTML_ `data-* attribute`_.

Install GopherJS_ and DOM_ bindings
+++++++++++++++++++++++++++++++++++

Run the following command to install GopherJS_ and
`GopherJS bindings for the JavaScript DOM APIs`_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs
  $ go get -u honnef.co/go/js/dom

Source Code
+++++++++++

If you have following element in HTML:

.. code-block:: html

  <div id="foo" data-my-demo-value="Hey"></div>

Access the value by *Dataset()* function, which returns *map[string]string*:

.. code-block:: go

  import "honnef.co/go/js/dom"

  ...

  d := dom.GetWindow().Document()
  div1 := d.GetElementByID("foo").(*dom.HTMLDivElement)
  v := div1.Dataset()["myDemoValue"]

Note that **data-my-demo-value** becomes **myDemoValue** when accessed.
Full example is as follows:

*HTML*:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/data/index.html

*Go*:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/data/data.go

Compile the Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build data.go -o demo.js

Put *demo.js* together with the *index.html* in the same directory and open the
*index.html* with your browser. Open the browser console and you will see the
printed value.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.2``.

----

GopherJS_ DOM_ Example series
+++++++++++++++++++++++++++++

- `[Golang] GopherJS DOM Example - getElementById and Set innerHTML <{filename}../10/gopherjs-dom-example-getElementById-innerHTML%en.rst>`_

- `[Golang] GopherJS DOM Example - Event Binding (addEventListener) <{filename}../11/gopherjs-dom-example-event-binding-addEventListener%en.rst>`_

- `[Golang] GopherJS DOM Example - Detect Keypress (Keyboard Event) <{filename}../11/gopherjs-dom-example-detect-keypress-keyboard-event%en.rst>`_

- `[Golang] GopherJS DOM Example - Access Input Element Value <{filename}../11/gopherjs-dom-example-access-input-element-value%en.rst>`_

- `[Golang] Online Input Method (Pāli) by GopherJS <{filename}go-online-input-method-pali-by-gopherjs%en.rst>`_

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

.. [5] `[Dart] Access HTML Data Attribute <{filename}../../../2015/03/01/dart-access-html-data-attribute%en.rst>`_

.. [6] `HTML Global data-* Attributes - W3Schools <http://www.w3schools.com/tags/att_global_data.asp>`_

.. [7] Google search : `HTML Data Attribute <https://www.google.com/search?q=HTML+Data+Attribute>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _HTML: http://www.w3schools.com/html/
.. _data-* attribute: http://www.w3schools.com/tags/att_global_data.asp
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
