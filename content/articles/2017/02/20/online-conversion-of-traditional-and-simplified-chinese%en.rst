Online Conversion of Traditional and Simplified Chinese in Go/GopherJS
######################################################################

:date: 2017-02-20T17:01+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Web application, web, DOM, html,
       Conversion of Traditional and Simplified Chinese, String Manipulation
:category: GopherJS
:summary: Go_ front-end programming - online service for conversion of
          Traditional and Simplified Chinese via GopherJS_/godom_/gojianfan_.
:og_image: https://pbs.twimg.com/profile_images/605816243870760960/4hP2sH_O.png
:adsu: yes


Go_ front-end programming - online service for conversion of Traditional and
Simplified Chinese via GopherJS_/godom_/gojianfan_.

.. rubric:: `Demo <https://siongui.github.io/gojianfan/>`_
   :class: align-center

This shows you the great power of GopherJS_ compiler. As long as the Go packages
are written in Go and not involved with OS operation, you can compile almost
any Go code to JavaScript and run on browser.

This online service use the following packages to provide Chinese conversion:

- godom_: `DOM manipulation`_ in Go.
- gojianfan_: converter for Traditional-Simplified Chinese

.. adsu:: 2

**HTML**:

.. show_github_file:: siongui gojianfan gopherjs/index.html
.. adsu:: 3

**Go**: (compiled to JavaScript by GopherJS)

.. show_github_file:: siongui gojianfan gopherjs/app.go
.. adsu:: 4

The demo consists of only a few lines of code. View
`gopherjs directory in gojianfan repository`_ for more details.

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8``, ``GopherJS 1.8-1``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)
.. adsu:: 5
.. [2] `GitHub - siongui/godom: Make DOM manipultation in Go as similar to JavaScript as possible. (via GopherJS) <https://github.com/siongui/godom>`_
.. [3] `GitHub - siongui/gojianfan: Traditional and Simplified Chinese Conversion in Go <https://github.com/siongui/gojianfan>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _gojianfan: https://github.com/siongui/gojianfan
.. _godom: https://github.com/siongui/godom
.. _GopherJS: http://www.gopherjs.org/
.. _DOM manipulation: https://www.google.com/search?q=DOM+manipulation
.. _gopherjs directory in gojianfan repository: https://github.com/siongui/gojianfan/tree/master/gopherjs

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
