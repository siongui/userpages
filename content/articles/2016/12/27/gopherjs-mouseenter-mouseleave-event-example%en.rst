[GopherJS] mouseenter and mouseleave Event Example
##################################################

:date: 2016-12-27T21:34+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM, mouseenter, mouseleave
:category: GopherJS
:summary: Example of onmouseenter_ and onmouseleave_ event via GopherJS_.
:adsu: yes


Example of onmouseenter_ and onmouseleave_ event via GopherJS_.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/mouseenter-mouseleave/index.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/mouseenter-mouseleave/app.go
.. adsu:: 3

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
.. adsu:: 4
.. [2] | `mouseenter - Event reference | MDN <https://developer.mozilla.org/en/docs/Web/Events/mouseenter>`_
       | `mouseleave - Event reference | MDN <https://developer.mozilla.org/en-US/docs/Web/Events/mouseleave>`_


.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _onmouseenter: https://developer.mozilla.org/en/docs/Web/Events/mouseenter
.. _onmouseleave: https://developer.mozilla.org/en/docs/Web/Events/mouseleave

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
