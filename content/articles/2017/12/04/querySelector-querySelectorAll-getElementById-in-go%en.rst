querySelector, querySelectorAll, getElementById in Go
#####################################################

:date: 2017-12-09T22:05+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Show how to use *querySelector*, *querySelectorAll*, *getElementById*
          in Go/GopherJS.
:og_image: http://blog.guinatal.com/wp-content/uploads/2015/05/getElementById.png
:adsu: yes


This post shows you how to querySelector_, querySelectorAll_, and
getElementById_ to get the element(s) we need in the DOM tree for further
processing via Go/GopherJS. If you need offile HTML processing library, you can
take a look at goquery_ or soup_ on GitHub.

.. contents:: **Table of Content**

querySelector
=============

JavaScript
++++++++++

*querySelector* returns a DOM element. The following JavaScript code returns
the element with id *foo*:

.. code-block:: javascript

  var elm = document.querySelector("#foo");


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  d := js.Global.Get("document")
  elm := d.Call("querySelector", "#foo")


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  elm := Document.QuerySelector("#foo")


.. adsu:: 2

----


querySelectorAll
================

JavaScript
++++++++++

*querySelectorAll* returns a NodeList_. The following code selects all *div*
elements in the *document*.

.. code-block:: javascript

  var nodeList = document.querySelectorAll("div");
  for (var i = 0; i < nodeList.length; ++i) {
    var elm = nodeList[i];
    // do something with the element
  }


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  d := js.Global.Get("document")
  nodeList := d.Call("querySelectorAll", "div")
  length := nodeList.Get("length").Int()
  for i := 0; i < length; i++ {
  	elm := nodeList.Call("item", i)
  	// do something with the element
  }


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  nodeList := Document.QuerySelectorAll("div")
  for _, elm := range nodeList {
  	// do something with the element
  }

.. adsu:: 3


getElementById
==============

JavaScript
++++++++++

The following code returns the element with id *foo.*

.. code-block:: javascript

  var element = document.getElementById("foo");


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  element := js.Global.Get("document").Call("getElementById", "foo")


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  element := Document.GetElementById("foo")


.. adsu:: 4


----

References:

.. [1] `[Golang] GopherJS Synonyms with JavaScript <{filename}../../../2016/01/29/go-gopherjs-synonyms-with-javascript%en.rst>`_
.. [2] `[Golang] GopherJS DOM Example - getElementById and Set innerHTML <{filename}../../../2016/01/10/gopherjs-dom-example-getElementById-innerHTML%en.rst>`_
.. [3] `[Golang] querySelectorAll and querySelector Example by GopherJS <{filename}../../../2016/02/14/go-querySelectorAll-querySelector-by-gopherjs%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _querySelector: https://www.google.com/search?q=querySelector
.. _querySelectorAll: https://www.google.com/search?q=querySelectorAll
.. _getElementById: https://www.google.com/search?q=getElementById
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _soup: https://github.com/anaskhan96/soup
.. _NodeList: https://developer.mozilla.org/en-US/docs/Web/API/NodeList
