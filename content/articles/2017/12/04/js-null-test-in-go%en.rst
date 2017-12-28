JavaScript null Check in Go
###########################

:date: 2017-12-28T23:37+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Test if a JavaScript variable is null in Go/GopherJS.
:og_image: https://i.stack.imgur.com/ThA4b.png
:adsu: yes


Sometimes we need to check if an value/object/variable in JavaScript is null.
For example, *querySelector* for elements will return null if nothing found, or
event state in the callback of history API is null in some cases. In this post
we will show you how to do null check in Go/GopherJS.

.. contents:: **Table of Content**

Test if Variable is null
========================

querySelector_ method will return first element which matches the selector, or
return *null* if nothing found. Here we have an empty HTML document without any
meaningful elements. We will query for an non-existing element with *id=foo* via
*querySelector* method. The query will return *null* and we will check for it.


JavaScript
++++++++++

.. code-block:: javascript

  var f = document.querySelector("#foo");

  if (f === null) {
      console.log("querySelector #foo returns null");
  } else {
      console.log("querySelector #foo returns element");
  }

Open your console and you will see *querySelector #foo returns null*.


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  f := js.Global.Get("document").Call("querySelector", "#foo")

  if f == nil {
  	println("querySelector #foo returns null")
  } else {
  	println("querySelector #foo returns element")
  }

.. rubric:: `Run Code on GopherJS Playground <https://gopherjs.github.io/playground/#/59HcuBcHOk>`__
   :class: align-center

.. adsu:: 2

The full code example of this post is `on my GitHub`_.

For real-world example, see [2]_.

----

References:

.. [1] `[GopherJS] null Test <{filename}../../../2017/01/05/gopherjs-null-test%en.rst>`_
.. [2] `[GopherJS] HTML Web History API Example <{filename}../../../2017/01/03/gopherjs-html-web-history-api-example%en.rst>`_
.. [3] `[Golang] undefined Test in GopherJS <{filename}../../../2016/02/06/go-undefined-test-in-gopherjs%en.rst>`_
.. [4] `[Golang] Test if Item Exist in Web Storage by GopherJS <{filename}../../../2016/02/16/go-test-if-item-exist-in-web-storage-by-gopherjs%en.rst>`_
.. [5] `[GopherJS] Setting Implementation via JSON and Web Storage (localStorage) <{filename}../../../2017/01/01/gopherjs-implement-setting-via-json-and-localStorage%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/014-javascript-null-test
.. _queryselector: https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
