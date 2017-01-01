[Golang] Test if Item Exist in Web Storage by GopherJS
######################################################

:date: 2016-02-16T02:23+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Web Storage
:category: GopherJS
:summary: Check if an item exists (is populated) in web storage (localStorage_
          or sessionStorage_) by GopherJS_.


Check if an item exists (is populated) in web storage (localStorage_ or
sessionStorage_) by GopherJS_. Use `js.Undefined`_ in your if test.

In the beginning, I use the following code:

.. code-block:: go

  if js.Global.Get("localStorage").Call("getItem", "demoSetting") == js.Undefined {
          // demoSetting not exist
  }

But it did not work as expected. After trial and error, the correct way is:

.. code-block:: go

  if js.Global.Get("localStorage").Get("demoSetting") == js.Undefined {
          // demoSetting not exist
  }

You can try it by yourself on
`GopherJS Playground <http://www.gopherjs.org/playground/#/6TbNwCz8ho>`__.

.. code-block:: go

  package main

  import (
          "fmt"

          "github.com/gopherjs/gopherjs/js"
  )

  func main() {
          if js.Global.Get("localStorage").Get("demoSetting") == js.Undefined {
                  fmt.Println("demoSetting not in localStorage API")
          }
          if js.Global.Get("localStorage").Call("getItem", "demoSetting") == js.Undefined {
                  fmt.Println("demoSetting not in localStorage API")
          }
  }

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Search undefined Results in GopherJS repo · GitHub <https://github.com/gopherjs/gopherjs/search?utf8=%E2%9C%93&q=undefined>`_

.. [3] `GitHub - go-humble/locstor: localstorage bindings for go and gopherjs <https://github.com/go-humble/locstor>`_

.. [4] `Google Search: localstorage <https://www.google.com/search?q=localstorage>`_

.. [5] `Storage - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Storage>`_

.. [6] `Using the Web Storage API - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Web_Storage_API/Using_the_Web_Storage_API>`_

.. _GopherJS: http://www.gopherjs.org/
.. _localStorage: https://developer.mozilla.org/en/docs/Web/API/Window/localStorage
.. _sessionStorage: https://developer.mozilla.org/en/docs/Web/API/Window/sessionStorage
.. _js.Undefined: https://godoc.org/github.com/gopherjs/gopherjs/js#Object

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
