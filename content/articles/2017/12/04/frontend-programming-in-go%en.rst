Frontend Programming in Go
##########################

:date: 2017-12-04T22:06+08:00
:modified: 2017-12-05T23:36+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go
:category: Go
:summary: Discuss why I choose Go to be my frontend programming language. Wirte
          Go code and run your application on web browsers.
:og_image: https://pbs.twimg.com/profile_images/605816243870760960/4hP2sH_O.png
:adsu: yes


When it comes to frontend programming, most people will mention JavaScript,
Vue.js, React, etc. But when the code grow bigger and bigger, it becomes very
difficult to maintain the code, especially when you come back to the code after
half a year without touching the code.

The JavaScript is full of "surprise", which means a lot of tricks and
unintuitive meanings of some concepts are inherent in the JavaScript. It makes
other people difficult to understand the code after code becomes bigger. Even
myself cannot understand the JavaScript code I written half an year ago.


Advantages of Go
++++++++++++++++

1. Go is simple: The `official website of Go`_ programming language provides
   very good online tutorials. You can master the main language concepts and run
   the code online very easily. There are no fancy syntax sugars in Go, which
   makes you get less pain when you read the Go code written by other people.

2. Go standard library is rich: When you want to write some feature in your
   application, usually you can easily implement the feature via the methods
   provided in standard library, or search GitHub or do googling for the feature
   you need, the chances are that someone already implement it.

3. Modularity: When your program grows bigger, inevitably you need to organize
   your code into pieces which make the code easy to understand and maintain.
   In Go, you can move some of the code into *packages*, and combine the
   packages into the your application.

4. Easy for testing: Go officially provides automated testing for packages.
   Follow the Go conventions to write you testing code and run `go test`, then
   it's done!

5. Reusable for both frontend and backend: One reason for node.js to be so
   popular is that you can write JavaScript code and run it on both frontend
   (browsers) and backend (servers). Go is already very popular on backend
   server development, why not write Go code also on the frontend?

6. Maintenance: Go's *less is more* philosophy, packages, and automated testing
   make it easy to maintain the Go code. When you leave your code for several
   months and come back to the code, or read the code by others, you will easily
   pick it up and do whatever you need to do.

7. Easy to transpile to JavaScript: GopherJS is a great transpiler
   (source-to-source compiler) which compiles ALL your Go code into JavaScript
   under almost all situations (if your code does not involve with IO
   operation). For example, I have an offline application which does Chinese
   conversion. Without any modification, I successfully make it run on browsers
   with GopherJS. This is really impressing, and very easy to port existing Go
   code to web browsers without any modification if no IO involved.

.. adsu:: 2


Run Go Code on Browser
++++++++++++++++++++++

There are two ways to make your Go code run on browsers. One is via transpilers
which compiles the Go code to JavaScript. The other is via WebAssembly, which is
still not usable and under development now.

1. Transpiler: The first way to run Go code on browsers is via GopherJS_, which
   transpiles Go code into JavaScript. As mentioned in previous section, it
   compiles almost everything in Go to JavaScript if no IO or OS operations
   involved. Besides, You may need a Go library to help you do DOM
   manipulation. You can take a look of my godom_ project on GitHub to make DOM
   manipulation easy and similar to JavaScript counterpart.

2. WebAssembly: WebAssembly is the future for web application development. Right
   now all major browsers support wasm, but there is no working compiler so far
   to compile Go to wasm. There is an issue that tracks the current status of
   Go wasm compiler, see [1]_ for more information.


Table of Content
++++++++++++++++

This post is the first to introduce Go in the field of frontend programming. I
will write a series of posts to show you how to write and run Go code on the
browsers. If any advices, please leave your comments on GitHub!

- `First Frontend Go Program - Hello World in Browser`_

.. adsu:: 3

----

References:

.. [1] | `meta: WebAssembly ("wasm") support · Issue #18892 · golang/go · GitHub <https://github.com/golang/go/issues/18892>`_
       | `Is there a webassembly compiler for Go in the works? : golang <https://www.reddit.com/r/golang/comments/5yl984/is_there_a_webassembly_compiler_for_go_in_the/>`_

.. [2] `GitHub - lpereira/gomoku: Gomoku compiles Go code to C++ <https://github.com/lpereira/gomoku>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _official website of Go: https://golang.org/
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _Go Playground: https://play.golang.org/
.. _godom: https://github.com/siongui/godom
.. _First Frontend Go Program - Hello World in Browser: {filename}first-frontend-go-program-hello-world%en.rst
