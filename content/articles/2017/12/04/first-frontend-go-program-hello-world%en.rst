First Frontend Go Program - Hello World in Browser
##################################################

:date: 2017-12-05T23:27+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: The first frontend program in Go - Show *Hello World* in your browser.
:og_image: https://pbs.twimg.com/profile_images/605816243870760960/4hP2sH_O.png
:adsu: yes


As usual, the first program of frontend programming in Go is to show 'Hello
World' in your browser. Here I assume you already have basic understanding of
JavaScript and DOM manipulation. If not, please read
`online JavaScript tutorials`_ first.


Install Go and GopherJS
+++++++++++++++++++++++

Visit `official website of Go`_ programming language. `Download and install`_ Go
in your working environment. The instructions on the Go official site are very
clear. Please follow them to install Go.

Next, visit GopherJS_ website. If you install Go properly, you can install
GopherJS via *go get* command as follows:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

Note that the current version of Go is 1.9.2, the current master branch of
GopherJS must be compiled with Go 1.9 series. If you use Go 1.8 series, you must
also use GopherJS 1.8 branch. Otherwise *go get* command will fail to compile
GopherJS.

OK, now we are ready for our first *Hello World* program.


Hello World via GopherJS
++++++++++++++++++++++++

Now we start to write our first frontend program in Go. We prepare the HTML file
for the running of JavaScript.

.. show_github_file:: siongui frontend-programming-in-go 001-hello-world/index.html
.. adsu:: 2

Next, we will write a Go program to call `alert() method of window object`_ and
show *Hello World* in pop-up window. We will compile the Go code into JavaScript
via GopherJS.

.. show_github_file:: siongui frontend-programming-in-go 001-hello-world/app.go

The `js.Global` here is actually the window_ object that we use in JavaScript
counterpart. You can try the code on GopherJS Playground:

.. rubric:: `Run Code on GopherJS Playground <https://gopherjs.github.io/playground/>`_
   :class: align-center

The following line in Go:

.. code-block:: go

  js.Global.Call("alert", "Hello World")

is actually the same as the following in JavaScript:

.. code-block:: javascript

  alert("Hello World");

Assume the name of previous Go file is `app.go`, we can compile the Go code to
JavaScript by the following command:

.. code-block:: bash

  $ gopherjs build app.go -o app.js

Put the compiled js file together with HTML and open the HTML with your browser,
you will see a pop-up windows with the message *Hello World*.

.. adsu:: 3


Hello World via GopherJS + godom
++++++++++++++++++++++++++++++++

In above example, we use the syntax directly provided by GopherJS. The syntax is
ugly and makes the code difficult to read. I write a package called godom_ which
makes the syntax similar to JavaScript.

Install godom_ package by the following command:

.. code-block:: bash

  $ go get -u github.com/siongui/godom

Now we use godom_ to re-write above example as follows:

.. show_github_file:: siongui frontend-programming-in-go 001-hello-world/appdom.go

The program now looks more like JavaScript counterpart and easy to read. Compile
and run this program as above. You will see the same result.

.. adsu:: 4

Summary
+++++++

This post shows the basic steps of frontend Go programming via GopherJS and
godom_. I do not recommend to use directly the syntax provided by GopherJS.
Instead, to use the syntax provided by godom_ will make the code more readable.

You can view and download the source code from `my GitHub project`_.

----

.. [1] `Go Hacker News - A HN client built with GopherJS : golang <https://old.reddit.com/r/golang/comments/9qr56e/go_hacker_news_a_hn_client_built_with_gopherjs/>`_
.. [2] `Animated QR data transfer with Gomobile and Gopherjs : golang <https://old.reddit.com/r/golang/comments/9yrk1g/animated_qr_data_transfer_with_gomobile_and/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _official website of Go: https://golang.org/
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _Go Playground: https://play.golang.org/
.. _online JavaScript tutorials: https://www.google.com/search?q=online+JavaScript+tutorials
.. _Download and install: https://golang.org/dl/
.. _alert() method of window object: https://www.w3schools.com/jsref/met_win_alert.asp
.. _window: https://www.w3schools.com/jsref/obj_window.asp
.. _godom: https://github.com/siongui/godom
.. _my GitHub project: https://github.com/siongui/frontend-programming-in-go/tree/master/001-hello-world
