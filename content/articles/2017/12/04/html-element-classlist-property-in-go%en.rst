HTML Element classList Property in Go
#####################################

:date: 2017-12-20T23:09+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, CSS
:category: Frontend Programming in Go
:summary: Manipulate the *class* attribute of a HTML element via *classList*
          property in Go/GopherJS.
:og_image: https://www.bennadel.com/resources/uploads/2017/dom-class-list-property.png
:adsu: yes


Manipulate CSS classes of an HTML element via classList_ property of the
element. We can *add*, *remove*, or *toggle* the class of an element, or test if
some class exists in the class attribute of the element via *contains* method.
For the full list of methods provided in *classList*, see
`Element.classList - Web APIs | MDN`_.

.. contents:: **Table of Content**

Manipulate classList Property
=============================

Assume we have the following *div* element in our HTML:

.. code-block:: html

  <div id="foo">Hello World</div>

And also the following CSS class:

.. code-block:: css

  .invisible { display: none; }

We will show you how to add the *invisible* class to the *div* element via the
*classList* property. For other manipulations like *remove*, *toggle*, or
*contains*, the usage is similar and we will not cover here.


JavaScript
++++++++++

.. code-block:: javascript

  var f = document.querySelector("#foo");

  f.classList.add("invisible");

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

  	f.Get("classList").Call("add", "invisible")
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

  	f.ClassList().Add("invisible")
  }


The full code example of this post is `on my GitHub`_.

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
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/011-element-classlist
.. _classList: https://www.google.com/search?q=classList
.. _Element.classList - Web APIs | MDN: https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
