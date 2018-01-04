Element Position (Scroll Included) in Go
########################################

:date: 2018-01-05T05:05+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, DOM,
       element offset, element position
:category: Frontend Programming in Go
:summary: Get HTML DOM element position (including scroll position of the
          document) in Go/GopherJS.
:og_image: https://www.kirupa.com/html5/images/boxmodel.png
:adsu: yes

This post will show you how to get HTML DOM element position, including scroll
position of the document, in Go/GopherJS. An example of getting element position
is to show the tooltip when the mouse hovers over texts [2]_. We need to get the
position of the text and show the tooltip next to the text.

.. contents:: **Table of Content**

Get HTML DOM Element Position
=============================

- To get the position of HTML DOM element, use getBoundingClientRect_ method.
- To get scroll position of the docuemnt, use pageXOffset_ and pageYOffset_
  properties of window object.
- Combine the position from getBoundingClientRect_ and
  pageXOffset_/pageYOffset_, we can obtain the position of the element.

JavaScript
++++++++++

.. code-block:: javascript

  function GetPosition(elm) {
    var x = elm.getBoundingClientRect().left + window.pageXOffset;
    var y = elm.getBoundingClientRect().top + window.pageYOffset;
    return [x, y];
  }


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  func GetPosition(elm *js.Object) (x, y float64) {
  	x = elm.Call("getBoundingClientRect").Get("left").Float() +
  		js.Global.Get("pageXOffset").Float()
  	y = elm.Call("getBoundingClientRect").Get("top").Float() +
  		js.Global.Get("pageYOffset").Float()
  	return
  }


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  func GetPosition(elm *Object) (x, y float64) {
  	x = elm.GetBoundingClientRect().Left() + Window.PageXOffset()
  	y = elm.GetBoundingClientRect().Top() + Window.PageYOffset()
  	return
  }

The full code example of this post is `on my GitHub`_.

.. adsu:: 2

----

References:

.. [1] `JavaScript DOM Element Position (Scroll Position Included) <{filename}../../../2012/07/01/javascript-dom-element-position-scroll-included%en.rst>`_
.. [2] `[JavaScript] Show Note on Mouse Hovering Over Text <{filename}../../../2018/01/04/javascript-show-annotation-on-mouse-hover-over-text%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/017-element-position
.. _getBoundingClientRect: https://www.google.com/search?q=getBoundingClientRect
.. _pageXOffset: https://www.google.com/search?q=pageXOffset
.. _pageYOffset: https://www.google.com/search?q=pageYOffset
