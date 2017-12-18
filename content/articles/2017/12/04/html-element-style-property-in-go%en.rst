HTML Element style Property in Go
#################################

:date: 2017-12-18T23:20+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, CSS
:category: Frontend Programming in Go
:summary: Set/Get inline style of a HTML element via *style* property in
          Go/GopherJS.
:adsu: yes


We can change how a HTML element is displayed by setting the `inline style`_ of
the element. This post will show you how to set as well as get the *style*
property of a HTML element.
The full code example of this post is `on my GitHub`_.

.. contents:: **Table of Content**

Inline Style Attribute
======================

Assume we have the following *div* element in our HTML:

.. code-block:: html

  <div id="foo">Hello World</div>

We will show you how to set and get the inline *style* attribute of the *div*
element. In the example, we will set the color of the element as red. For full
list of CSS properties accessible via *style*, see
`CSS Properties Reference on MDN`_.


JavaScript
++++++++++

.. code-block:: javascript

  var f = document.querySelector("#foo");

  // set the color of element
  f.style.color = "red";

  // get the color of element
  console.log(f.style.color);

.. adsu:: 2


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  func main() {
  	f := js.Global.Get("document").Call("querySelector", "#foo")

  	// set the color of element
  	f.Get("style").Set("color", "red")

  	// get the color of element
  	println(f.Get("style").Get("color").String())
  }

.. adsu:: 3


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  func main() {
  	f := Document.QuerySelector("#foo")

  	// set the color of element
  	f.Style().SetColor("red")

  	// get the color of element
  	println(f.Style().Color())
  }

.. adsu:: 4

----

References:

.. [1] `[Golang] GopherJS DOM Example - Hide Element by display:none <{filename}../../../2016/01/13/gopherjs-dom-example-hide-element-by-display-none%en.rst>`_
.. [2] `[GopherJS] Set/Get DOM CSS <{filename}../../../2016/06/01/gopherjs-set-get-dom-css%en.rst>`_
.. [3] `[GopherJS] Insert CSS Dynamically <{filename}../../../2016/06/04/gopherjs-add-css-dynamically%en.rst>`_
.. [4] `[GopherJS] Test if an Element Contains a Class <{filename}../../../2017/01/15/gopherjs-test-if-an-element-contains-a-class%en.rst>`_
.. [5] `[GopherJS] Animate.css Test Demo <{filename}../../../2017/01/24/gopherjs-animate.css-test-demo%en.rst>`_
.. [6] `[Golang] GopherJS Synonyms with JavaScript <{filename}../../../2016/01/29/go-gopherjs-synonyms-with-javascript%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/010-element-style
.. _inline style: https://www.google.com/search?q=element+style
.. _CSS Properties Reference on MDN: https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Properties_Reference
