GopherJS XMLHttpRequest (XHR) and MakeFunc Example
##################################################

:date: 2016-02-18T04:53+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, XMLHttpRequest
:category: GopherJS
:summary: Show how to use MakeFunc_ in GopherJS_ by XMLHttpRequest_ (XHR_)
          example.


Show how to use MakeFunc_ in GopherJS_ by XHR_ example. This example use
MakeFunc_ to wrap a callback function in XMLHttpRequest_ (XHR) requests [3]_.

.. code-block:: go

  var readyStateChange = js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
          if this.Get("readyState").Int() == 4 {
                  if this.Get("status").Int() == 200 {
                          handleGetWordOK(this.Get("responseText").String())
                  } else {
                          handleGetWordError()
                  }
          }
          return nil
  })

  func xhrGetWordJson(w string) {
          req := js.Global.Get("XMLHttpRequest").New()
          req.Call("addEventListener", "readystatechange", readyStateChange)
          req.Call("open", "GET", HttpWordJsonPath(w), true)
          req.Call("send")
  }

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `print out word explanation · siongui/pali@d140c64 · GitHub <https://github.com/siongui/pali/commit/d140c645d80afea99a344bb3ebf098f5fae0c63b>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _MakeFunc: https://godoc.org/github.com/gopherjs/gopherjs/js#MakeFunc
.. _XMLHttpRequest: https://www.google.com/search?q=XMLHttpRequest
.. _XHR: https://www.google.com/search?q=XMLHttpRequest

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
