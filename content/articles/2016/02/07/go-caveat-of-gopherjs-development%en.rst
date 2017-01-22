[Golang] Caveats of GopherJS Development
########################################

:date: 2016-02-07T11:32+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, XMLHttpRequest, html, Goroutine
:category: GopherJS
:summary: Fix runtime error via goroutine_ in Go_ code compiled to JavaScript_
          by GopherJS_.
:adsu: yes

Fix runtime error via goroutine_ in Golang_ code compiled to JavaScript_ by
GopherJS_. I got runtime error when writing code to make XHR_ requests and also
code to generate HTML_ via `html/template`_

The following code is an XHR_ request to retrieve JSON data from server: [7]_

.. code-block:: go

    func XhrGetWordInfo(word string) (w WordInfo, err error) {
        resp, err := http.Get(HttpWordJsonPath(word))
        if err != nil {
                return
        }
        defer resp.Body.Close()

        dec := json.NewDecoder(resp.Body)
        err = dec.Decode(&w)
        return
    }

*XhrGetWordInfo* is called by:

.. code-block:: go

  XhrGetWordInfo("sacca")

The following error showed up in browser console: [4]_

.. code-block:: txt

  Uncaught Error: runtime error: cannot block in JavaScript callback, fix by wrapping code in goroutine
  Uncaught TypeError: r is not a function

After some trial and error, the correct way to make XHR_ requests is: [5]_

.. code-block:: go

  go XhrGetWordInfo("sacca")

The code must be wrapped in goroutine_ to prevent blocking the execution.

Similar situation [6]_ happened again and the same solution was applied to fix
the runtime error when generating HTML_ via `html/template`_.


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``,
``Chromium Version 48.0.2564.82 Ubuntu 15.10 (64-bit)``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)

.. [4] `net/http not working · Issue #389 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/389>`_

.. [5] `use net/http by goroutine · siongui/pali@2445656 · GitHub <https://github.com/siongui/pali/commit/244565656c019a41625fd4337594b757cbfb606e>`_

.. [6] `fix bug: use html/template by goroutine · siongui/pali@39ffdf9 · GitHub <https://github.com/siongui/pali/commit/39ffdf9589c98fa8ed85a09a609cfe3e936897d8>`_

.. [7] `handle input text enter event. Bug: net/http not working · siongui/pali@f8f1475 · GitHub <https://github.com/siongui/pali/commit/f8f1475af0935419b29e7f79963e1d7e4a0b5944>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model
.. _HTML: http://www.w3schools.com/html/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _XHR: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
.. _goroutine: https://tour.golang.org/concurrency/1
.. _html/template: https://golang.org/pkg/html/template/

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
