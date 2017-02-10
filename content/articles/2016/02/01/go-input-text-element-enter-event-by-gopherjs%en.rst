[Golang] Input Text Element Enter Keypress Event by GopherJS
############################################################

:date: 2016-02-01T12:53+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, Keyboard Event
:category: GopherJS
:summary: Add the enter keypress event handler for the `input text element`_ by
          GopherJS
:adsu: yes

Introduction
++++++++++++

In this post, we will give an example of DOM_ manipulation in Go_ program.
This example shows how to detect enter keypress event of the
`input text element`_ and register an event handler to it.

Install GopherJS_
+++++++++++++++++

Run the following command to install GopherJS_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

Source Code
+++++++++++

HTML_ code:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/input-enter/index.html
.. adsu:: 2

Go_ code:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/input-enter/enter.go
.. adsu:: 3

Compile the Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build enter.go -o demo.js

Put *demo.js* together with the *index.html* in the same directory. Open the
*index.html* with your browser and also the browser console. Press enter in the
input and you will see the message.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``,
``Chromium Version 48.0.2564.82 Ubuntu 15.10 (64-bit)``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_
.. adsu:: 4
.. [3] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)

.. [4] `input text enter event <https://www.google.com/search?q=input+text+enter+event>`_

.. [5] `Trigger a button click with JavaScript on the Enter key in a text box - Stack Overflow <http://stackoverflow.com/questions/155188/trigger-a-button-click-with-javascript-on-the-enter-key-in-a-text-box>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _HTML: http://www.w3schools.com/html/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _input text element: http://www.w3schools.com/tags/tag_input.asp

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
