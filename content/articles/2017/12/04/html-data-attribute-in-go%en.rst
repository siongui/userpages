HTML data-* Attribute in Go
###########################

:date: 2017-12-15T23:05+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, html
:category: Frontend Programming in Go
:summary: Access HTML data-* attributes of elements in Go/GopherJS.
:adsu: yes


Access `HTML data-* attributes`_ of elements in Go/GopherJS. Data attribute of
a element allows you to store and retrieve data from the specific element. For
example, you may store the information, such as geolocation or time, of an image
just in the *img* element. Data attributes can help you achieve this goal.
The full code example of this post is `on my GitHub`_.

.. contents:: **Table of Content**

Using data-* Attribute
======================

Assume we have the following *div* element in our HTML:

.. code-block:: html

  <div id="foo"></div>

We want to store custom data in the element. We can do the following:

.. code-block:: html

  <div id="foo" data-demo-value="hello world"></div>

We add an attribute starting with *data-* to the element. We can access the
*data-* attribute as follows:


JavaScript
++++++++++

.. code-block:: javascript

  var f = document.querySelector("#foo");

  // get value
  console.log(f.dataset.demoValue);

  // set value
  f.dataset.demoValue = "world hello";

The name *demo-value* in HTML will become *demoValue* in JavaScript.

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

  	// get value
  	println(f.Get("dataset").Get("demoValue").String())

  	// set value
  	f.Get("dataset").Set("demoValue", "world hello")
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

  	// get value
  	println(f.Dataset().Get("demoValue").String())

  	// set value
  	f.Dataset().Set("demoValue", "world hello")
  }

.. adsu:: 4

You can read `more detailed tutorial on MDN`_.

----

References:

.. [1] `[Golang] GopherJS DOM Example - Access HTML Data Attribute <{filename}../../../2016/01/12/gopherjs-dom-example-access-html-data-attribute%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/009-html-data-attribute
.. _HTML data-* attributes: https://www.google.com/search?q=html+data+attribute
.. _more detailed tutorial on MDN: https://developer.mozilla.org/en-US/docs/Learn/HTML/Howto/Use_data_attributes
