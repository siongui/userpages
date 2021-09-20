Install GopherJS in Go Modules Mode
###################################

:date: 2021-09-20T22:57+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript
:category: GopherJS
:summary: Build and run GopherJS in Go modules mode.
:og_image: https://opengraph.githubassets.com/3a3dec5f3d903360d5085f3ad77dcd8cc0c629abcd331d6c86bb3f2961e9a5d8/goplusjs/gopherjs
:adsu: yes


Starting from Go_ 1.16, the module mode is on by default. You may heard that
GOPATH_ is not needed in module mode, but if you want to install GopherJS_,
you still need to set it to build the GopherJS binary.

To install GopherJS, you need to set GOPATH_ environment variable first.

.. code-block:: bash

  $ export GOPATH=/my/path/to/gopath

Then install GopherJS:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

The executable GopherJS binary will be located in ``$(GOPATH)/bin``, and
remember to add this to your ``PATH`` variable.

Tested on: ``Ubuntu Linux 20.04``, ``Go 1.17.1``.

----

References:

.. [1] `gopherjs: command not found  · Issue #906 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/906>`_

.. _Go: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _godom: https://github.com/siongui/godom
.. _GOPATH: https://golang.org/doc/gopath_code
