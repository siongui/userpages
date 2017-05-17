[Golang] undefined Test in GopherJS
###################################

:date: 2016-02-06T02:24+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Web Storage
:category: GopherJS
:summary: Check if something (object, variable, API, ...) is undefined or not in
          GopherJS_.
:og_image: https://i.ytimg.com/vi/Ieue3ffpd9c/maxresdefault.jpg
:adsu: yes


Check if something (object, variable, API, ...) is undefined or not in
GopherJS_. Use `js.Undefined`_ in your if test.

For example, you can check if your browser supports localStorage_ API or not by:

.. code-block:: go

  package main

  import (
          "fmt"

          "github.com/gopherjs/gopherjs/js"
  )

  func main() {
          if js.Global.Get("localStorage") == js.Undefined {
                  fmt.Println("Your browser does not support localStorage API")
          } else {
                  fmt.Println("Your browser supports localStorage API")
          }
  }

.. adsu:: 2

`Run Code on GopherJS Playground <http://www.gopherjs.org/playground/#/Kxr4h5nxBQ>`_

For null_ test (if a variable is null), see [3]_.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Search undefined Results in GopherJS repo · GitHub <https://github.com/gopherjs/gopherjs/search?utf8=%E2%9C%93&q=undefined>`_

.. [3] `[GopherJS] null Test <{filename}../../../2017/01/05/gopherjs-null-test%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _localStorage: https://developer.mozilla.org/en/docs/Web/API/Window/localStorage
.. _js.Undefined: https://godoc.org/github.com/gopherjs/gopherjs/js#Object
.. _null: https://developer.mozilla.org/en/docs/Web/JavaScript/Reference/Global_Objects/null

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
