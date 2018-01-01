Show CSS Loader While Resource Loading in Go
############################################

:date: 2018-01-01T23:41+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, CSS
:category: Frontend Programming in Go
:summary: Show CSS loading spinner while resources of the website are loading
          in Go/GopherJS.
:adsu: yes

When we are writing web applications, the time for loading resources becomes
longer and longer as the code grows. It's a good idea to show a loading spinner
while the website resources are loading. This post will show you how to show a
CSS loading spinner while loading and after the whole website is loaded, we will
make the spinner invisible.

.. contents:: **Table of Content**

Show CSS Loading Spinner
========================

First we need to put loading spinner on our web page. after some googling, I
found `How To Make a CSS Loader`_ on w3schools. We put the following CSS code
for loading spinner in the *head* section of HTML:

.. code-block:: css

  /* From https://www.w3schools.com/howto/howto_css_loader.asp */
  .loader {
      border: 16px solid #f3f3f3; /* Light grey */
      border-top: 16px solid #3498db; /* Blue */
      border-radius: 50%;
      width: 120px;
      height: 120px;
      animation: spin 2s linear infinite;
  }

  @keyframes spin {
      0% { transform: rotate(0deg); }
      100% { transform: rotate(360deg); }
  }

We will hide the loader after finish. So also add the following CSS rule in the
*head* section of HTML.

.. code-block:: css

  .invisible { display: none; }

Then put the following loading spinner code in the appropriate place of HTML
body.

.. code-block:: html

  <div class="loader"></div>

After HTML and CSS code is ready, now we need to write some frontend code to
hide the spinner after the website is fully loaded. We will attach the event
handler to the `load event of window object`_. This event will be fired after
all resources are loaded. You can also attach the event handler to
`DOMContentLoaded event of document object`_, which fired when the initial HTML
document has been completely loaded and parsed, without waiting for stylesheets,
images, and subframes to finish loading.

In this example, we will attach the event handler to *load* event of window
object. When resources are fully loaded, hide the CSS loading spinner.

.. adsu:: 2


JavaScript
++++++++++

.. code-block:: javascript

  window.addEventListener("load", function(event) {
    var s = document.querySelector(".loader");
    s.classList.add("invisible");
    console.log("All resources finished loading!");
  });


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  func main() {
  	js.Global.Call("addEventListener", "load", func(event *js.Object) {
  		s := js.Global.Get("document").Call("querySelector", ".loader")
  		s.Get("classList").Call("add", "invisible")
  		println("All resources finished loading!")
  	})
  }


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )

  func main() {
  	Window.Call("addEventListener", "load", func(e Event) {
  		s := Document.QuerySelector(".loader")
  		s.ClassList().Add("invisible")
  		println("All resources finished loading!")
  	})
  }

The full code example of this post is `on my GitHub`_.

.. adsu:: 3

----

References:

.. [1] `Synonyms - Go and JavaScript <{filename}synonyms-go-and-javascript%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/016-show-css-loader
.. _How To Make a CSS Loader: https://www.w3schools.com/howto/howto_css_loader.asp
.. _load event of window object: https://developer.mozilla.org/en-US/docs/Web/Events/load
.. _DOMContentLoaded event of document object: https://developer.mozilla.org/en-US/docs/Web/Events/DOMContentLoaded
