[GopherJS] Copy to Clipboard
############################

:date: 2016-04-01T22:44+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, Copy to Clipboard
:category: GopherJS
:summary: Copy content of textarea_ to clipboard via GopherJS_.
:adsu: yes

Copy content of textarea_ to clipboard via GopherJS_.

.. show_github_file:: siongui userpages content/code/gopherjs-copy-to-clipboard/index.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/gopherjs-copy-to-clipboard/copy.go
.. adsu:: 3

Compile ``copy.go`` to ``copy.js`` by GopherJS_.
Put ``copy.js`` and ``index.html`` together.
Open ``index.html`` with your browser to see demo.

This comes from the answer of [2]_.
See the answer of [2]_ for browser support.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)
.. adsu:: 4
.. [2] | `How do I copy to the clipboard in JavaScript? - Stack Overflow <http://stackoverflow.com/a/30810322>`_
       | `javascript copy to clipboard <https://www.google.com/search?q=javascript+copy+to+clipboard>`_


.. _GopherJS: http://www.gopherjs.org/
.. _textarea: http://www.w3schools.com/tags/tag_textarea.asp

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
