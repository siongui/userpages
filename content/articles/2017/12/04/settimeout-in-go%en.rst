setTimeout method in Go
#######################

:date: 2017-12-21T22:10+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Execute a function after waiting a specific length of time in
          Go/GopherJS.
:og_image: https://javascript.info/article/settimeout-setinterval/settimeout-interval.png
:adsu: yes


setTimeout_ method in JavaScript is used to run a function after waiting a
specific length of time. The equivalent in Go is time.AfterFunc_ method in Go
standard library. This post show you how to use time.AfterFunc_ in Go and
compile the Go code to run on browsers via GopherJS. Note that if you are not
doing frontend programming, you can still use time.AfterFunc_ on backend
servers or local machines.

.. contents:: **Table of Content**

JavaScript setTimeout = Go time.AfterFunc
=========================================


JavaScript
++++++++++

.. code-block:: javascript

  setTimeout(function() {
    console.log("3 seconds timeout");
  }, 3000);


Go or GopherJS
++++++++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"time"
  )

  func main() {
  	time.AfterFunc(3*time.Second, func() {
  		println("3 seconds timeout")
  	})
  }

.. rubric:: `Run Code on GopherJS Playground <https://gopherjs.github.io/playground/#/LjCARICREZ>`__
   :class: align-center

Note that you do not have to import GopherJS package in your Go code. GopherJS
will compile time.AfterFunc_ code to corresponding JavaScript code for you. You
feel like writing a local Go program and do not even have to know about
*setTimeout* method! The code here can be used both on front and backend!

.. adsu:: 2

The full code example of this post is `on my GitHub`_.
For more complicated example, see my another post [3]_.

----

References:

.. [1] `[JavaScript] Type Text Effect <{filename}../../../2017/03/07/javascript-type-text-effect%en.rst>`_
.. [2] `[JavaScript] Typing Text Effect <{filename}../../../2017/03/08/javascript-typing-text-effect%en.rst>`_
.. [3] `[Golang/GopherJS] setTimeout <{filename}../../../2016/06/03/go-gopherjs-setTimeout%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/012-setTimeout
.. _setTimeout: https://www.google.com/search?q=setTimeout
.. _time.AfterFunc: https://golang.org/pkg/time/#AfterFunc
