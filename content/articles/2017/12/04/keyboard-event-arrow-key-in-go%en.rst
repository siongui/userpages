Keyboard Event (Arrow Keys) in Go
#################################

:date: 2017-12-12T23:41+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Detect user keypress on browsers in Go/GopherJS.
:adsu: yes


This post show how to detect user keypress on browsers in Go. We will use arrow
keys as example and show you how to detect arrow key presses on browsers in
Go/GopherJS.
The full code example of this post is `on my GitHub`_.

.. contents:: **Table of Content**

Detect Arrow Keys
=================

We will attach an *keyup* event handler to the *window* object of current
browser window/tab. When users press any key on the current browser window/tab,
the event handler will run and we will check the *keyCode* property of the
event. If the *keyCode* happens to be arrow keys, we will show which arrow key
user press via *div* element. The *div* element is as follows:

.. code-block:: html

  <div id="info">Press any arrow key</div>


JavaScript
++++++++++

.. show_github_file:: siongui frontend-programming-in-go 006-keyboard-event/app.js
.. adsu:: 2


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. show_github_file:: siongui frontend-programming-in-go 006-keyboard-event/app.go
.. adsu:: 3


GopherJS + godom
++++++++++++++++

To make your code more readable, we can prettify the above code with godom_:

.. show_github_file:: siongui frontend-programming-in-go 006-keyboard-event/appdom.go
.. adsu:: 4

----

References:

.. [1] `[GopherJS] Keyboard Event - Arrow Keys Example <{filename}../../../2016/12/31/gopherjs-keyboard-event-arrow-keys-example%en.rst>`_
.. [2] `[Golang] GopherJS DOM Example - Detect Keypress (Keyboard Event) <{filename}../../../2016/01/11/gopherjs-dom-example-detect-keypress-keyboard-event%en.rst>`_
.. [3] `JavaScript Keyboard Event (Arrow Key Example) <{filename}../../../2012/06/25/javascript-keyboard-event-arrow-key-example%en.rst>`_
.. [4] `JavaScript Arrow Key Example via event.key in Keyboard Event <{filename}../../02/14/javascript-arrow-key-example-via-event-key%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _addEventListener: https://www.google.com/search?q=addEventListener
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/006-keyboard-event
