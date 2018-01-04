Create and Append Element or Text Node in Go
############################################

:date: 2017-12-13T23:17+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Create and append DOM element or text node in Go/GopherJS.
          Show how to use *createElement*, *createTextNode*, and *appendChild*
          methods.
:og_image: https://image.slidesharecdn.com/introductiontojscnt-170123115538/95/introduction-to-js-cnt-14-638.jpg
:adsu: yes


This post shows how to create an DOM element or text node, and then append the
element or text node to another element. We will show you how to use
createElement_, createTextNode_, and appendChild_ methods in Go/GopherJS.
The full code example of this post is `on my GitHub`_.

.. contents:: **Table of Content**

Create/Append Element or Text Node
==================================

Assume we have the following *div* element in our HTML:

.. code-block:: html

  <div id="foo"></div>

We will create a *span* element and append it to the *div* element. Then we
create a text node whose content is *Hello World*, and append it to the *span*
element.


JavaScript
++++++++++

.. code-block:: javascript

  var f = document.querySelector("#foo");

  var s = document.createElement("span");
  f.appendChild(s);

  var t = document.createTextNode("Hello World");
  s.appendChild(t);

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

  	// create element
  	s := js.Global.Get("document").Call("createElement", "span")
  	// append element
  	f.Call("appendChild", s)

  	// create text node
  	t := js.Global.Get("document").Call("createTextNode", "Hello World")
  	// append text node
  	s.Call("appendChild", t)
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

  	// create element
  	s := Document.CreateElement("span")
  	// append element
  	f.AppendChild(s)

  	// create text node
  	t := Document.CreateTextNode("Hello World")
  	// append text node
  	s.AppendChild(t)
  }

.. adsu:: 4


----

References:

.. [1] `[GopherJS] createElement and createTextNode DOM Example <{filename}../../../2016/12/30/gopherjs-createElement-createTextNode-dom-example%en.rst>`_
.. [2] `[Golang] GopherJS DOM Example - Create and Append Element <{filename}../../../2016/01/14/gopherjs-dom-example-create-and-append-element%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/007-create-append-element
.. _createElement: https://www.google.com/search?q=createElement
.. _createTextNode: https://www.google.com/search?q=createTextNode
.. _appendChild: https://www.google.com/search?q=appendChild
