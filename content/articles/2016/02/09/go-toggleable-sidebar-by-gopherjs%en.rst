[Golang] Toggle-able Sidebar by GopherJS
########################################

:date: 2016-02-09T21:19+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, toggle, toggleable
:category: GopherJS
:summary: Go_ toggle-able sidebar by GopherJS_, inspired by Octopress_.

Golang_ toggle-able sidebar by GopherJS_, inspired by Octopress_.
This is a re-write of toggle-able sidebar of Dart_ version [3]_.

.. rubric:: Demo_ (GitHub_)
   :class: align-center

Install GopherJS_
+++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

Source Code
+++++++++++

.. show_github_file:: siongui go-toggle-sidebar-example src/toggle.html

.. show_github_file:: siongui go-toggle-sidebar-example src/toggle.css

.. show_github_file:: siongui go-toggle-sidebar-example src/toggle.go

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `GitHub - siongui/dart-toggle-sidebar-example: Inspired by Octopress, create toggle-able sidebar with Dart. <https://github.com/siongui/dart-toggle-sidebar-example>`_

.. [4] `GitHub - siongui/go-toggle-sidebar-example: Inspired by Octopress, create toggle-able sidebar with Golang. <https://github.com/siongui/go-toggle-sidebar-example>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _Octopress: http://octopress.org/
.. _Dart: https://www.dartlang.org/
.. _Demo: https://siongui.github.io/go-toggle-sidebar-example/
.. _GitHub: https://github.com/siongui/go-toggle-sidebar-example

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
