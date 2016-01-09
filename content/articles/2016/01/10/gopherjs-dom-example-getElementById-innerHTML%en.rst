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

.. show_github_file:: siongui userpages content/code/gopherjs-dom/index.html

We will access the *div* element whose *id* is *foo*. Then add *Hello World*
text to the *div*. Now write a Go_ program to manipulate DOM_:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/dom.go

It is very easy and intuitive. Compile the Go_ code to JavaScript_:

.. code-block:: bash

  $ gopherjs build dom.go -o dom.js

Put *dom.js* together with the *index.html* in the same directory and open the
*index.html* with your browser. You will see *Hello World* in the browser
window.

.. .. show_github_file:: siongui userpages content/code/gopherjs-dom/Makefile


Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.2``.

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
.. _getElementById: http://www.w3schools.com/jsref/met_doc_getelementbyid.asp
.. _innerHTML: http://www.w3schools.com/jsref/prop_html_innerhtml.asp
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
