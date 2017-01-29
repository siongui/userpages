[GopherJS] window.location - Access URL
#######################################

:date: 2017-01-02T22:39+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: Access URL, path, query string, etc. of current URL via
          window.location_ and GopherJS_.
:adsu: yes


Access URL, path, query string, etc. of current URL via window.location_ and
GopherJS_.

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  var location = js.Global.Get("location")

  func accessURL() {
  	// current URL: http://localhost:8000/code/gopherjs/window-location/index.html?a=1

  	// return - http://localhost:8000/code/gopherjs/window-location/index.html?a=1
  	location.Get("href").String()
  	// return - localhost:8000
  	location.Get("host").String()
  	// return - localhost
  	location.Get("hostname").String()
  	// return - /code/gopherjs/window-location/index.html
  	location.Get("pathname").String()
  	// return - http:
  	location.Get("protocol").String()
  	// return - http://localhost:8000
  	location.Get("origin").String()
  	// return - 8000
  	location.Get("port").String()
  	// return - ?a=1
  	location.Get("search").String()
  }

.. adsu:: 2

For JavaScript_ equivalent, see [2]_.

..
  .. rubric:: `Demo <{filename}/code/gopherjs/window-location/index.html?a=1>`_
     :class: align-center

  .. show_github_file:: siongui userpages content/code/gopherjs/window-location/index.html

  .. show_github_file:: siongui userpages content/code/gopherjs/window-location/app.go

  To see demo: use GopherJS_ to compile ``app.go`` to ``app.js``. Put
  ``index.html`` and ``app.js`` in the same directory. Open ``index.html`` with
  your browser.

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.4``
- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `[javascript] window.location example - access browser url <{filename}../09/javascript-window-location-example-access-url%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _window.location: http://www.w3schools.com/jsref/obj_location.asp
.. _JavaScript: https://www.google.com/search?q=JavaScript

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
