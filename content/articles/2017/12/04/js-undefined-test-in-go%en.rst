JavaScript undefined Check in Go
################################

:date: 2017-12-30T23:33+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Check for *undefined* in Go/GopherJS.
:og_image: https://i.stack.imgur.com/UTBcN.png
:adsu: yes


Sometimes we need to check if something exists or not in JavaScript.
For example, old browsers may not support *localStorage* API, so we need to
check before we use. In this post, we will show you how to do undefined check in
Go/GopherJS.

.. contents:: **Table of Content**

Test if It's Undefined
======================

Use *localStorage* as example, we will check if it is undefined in the browser.


JavaScript
++++++++++

.. code-block:: javascript

  if (window.localStorage === undefined) {
    console.log("Your browser does not support localStorage API");
  } else {
    console.log("Your browser supports localStorage API");
  }

Open your console and you will know if your browser support it or not.


GopherJS
++++++++

GopherJS provides `js.Undefined`_ Object for undefined test. The above code in
Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  if js.Global.Get("localStorage") == js.Undefined {
  	println("Your browser does not support localStorage API")
  } else {
  	println("Your browser supports localStorage API")
  }

.. rubric:: `Run Code on GopherJS Playground <http://www.gopherjs.org/playground/#/Kxr4h5nxBQ>`__
   :class: align-center

.. adsu:: 2

The full code example of this post is `on my GitHub`_.

For real-world example, see [3]_.

----

References:

.. [1] `[Golang] undefined Test in GopherJS <{filename}../../../2016/02/06/go-undefined-test-in-gopherjs%en.rst>`_
.. [2] `[Golang] Test if Item Exist in Web Storage by GopherJS <{filename}../../../2016/02/16/go-test-if-item-exist-in-web-storage-by-gopherjs%en.rst>`_
.. [3] `[GopherJS] Setting Implementation via JSON and Web Storage (localStorage) <{filename}../../../2017/01/01/gopherjs-implement-setting-via-json-and-localStorage%en.rst>`_
.. [4] `[GopherJS] null Test <{filename}../../../2017/01/05/gopherjs-null-test%en.rst>`_
.. [5] `[GopherJS] HTML Web History API Example <{filename}../../../2017/01/03/gopherjs-html-web-history-api-example%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/015-javascript-undefined-test
.. _js.Undefined: https://godoc.org/github.com/gopherjs/gopherjs/js#Object
