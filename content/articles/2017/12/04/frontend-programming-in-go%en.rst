Frontend Programming in Go
##########################

:date: 2017-12-04T22:06+08:00
:modified: 2020-07-15T08:43+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Go WebAssembly,
       Frontend Programming in Go
:category: Frontend Programming in Go
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
   code to web browsers without any modification if no IO involved. [4]_

.. adsu:: 2


Run Go Code on Browser
++++++++++++++++++++++

There are two ways to make your Go code run on browsers. One is via transpilers
which compiles the Go code to JavaScript. The other is via WebAssembly, which is
supported experimentally in the `Go 1.11`_ release.

1. Transpiler: The first way to run Go code on browsers is via GopherJS_, which
   transpiles Go code into JavaScript. As mentioned in previous section, it
   compiles almost everything in Go to JavaScript if no IO or OS operations
   involved. Besides, You may need a Go library to help you do DOM
   manipulation. You can take a look of my godom_ project on GitHub to make DOM
   manipulation easy and similar to JavaScript counterpart.

2. WebAssembly: WebAssembly is the future for web application development. Right
   now all major browsers support wasm, and the official Go compiler 1.11 adds
   an experimental port to WebAssembly (js/wasm). See [1]_ for more information
   and try it using `Go 1.11`_.


Table of Content
++++++++++++++++

This post is the first to introduce Go in the field of frontend programming. I
will write a series of posts to show you how to write and run Go code on the
browsers. If any advices, please leave your comments on GitHub!


Go WebAssembly
--------------

**Basics**:

- `Load and Run Go WebAssembly Module`_
- `[Go WebAssembly] First Wasm Program - Hello World`_
- `[Go WebAssembly] querySelector Example`_
- `[Go WebAssembly] querySelectorAll Example`_
- `[Go WebAssembly] Event Binding - addEventListener Example`_
- `[Go WebAssembly] XMLHttpRequest (XHR)`_


GopherJS
--------

**Basics**:

- `First Frontend Go Program - Hello World in Browser`_
- `JavaScript new Keyword in Go`_
- `querySelector, querySelectorAll, getElementById in Go`_
- `innerHTML and textContent in Go`_
- `Event Binding - addEventListener in Go`_
- `Keyboard Event (Arrow Keys) in Go`_
- `Create and Append Element or Text Node in Go`_
- `XMLHttpRequest (XHR) in Go`_
- `HTML data-* Attribute in Go`_
- `HTML Element style Property in Go`_
- `HTML Element classList Property in Go`_
- `setTimeout method in Go`_
- `Detect Browser Language Preference in Go`_
- `JavaScript null Check in Go`_
- `JavaScript undefined Check in Go`_
- `Element Position (Scroll Included) in Go`_
- `Check if Value of HTML Input Text Field is Integer in Go`_
- `Simulate Enter Key Pressed in Go`_

**Application**:

- `Show CSS Loader While Resource Loading in Go`_
- `Tooltip in Go`_
- `Tooltip with Close Delay in Go`_
- `Convert Text to HTML Link in Go`_
- `[GopherJS] WebSocket Client Example With Echo Server`_
- `Online Lemoine’s Conjecture Demo in Go`_

Bulma_:

- `Bulma Navbar with Go Toggle`_
- `Bulma Dropdown with Go Toggle`_
- `Bulma Modal with Go Toggle`_

Vue.js_:

- `Toggle (Show/Hide) HTML Element via Go and Vue.js`_
- `Show keyCode of Pressed Key via Go and Vue.js`_
- `Virtual Keyboard via Go and Vue.js`_
- `Online Sieve of Eratosthenes Demo via Go and Vue.js`_
- `Watch Data Change via Go and Vue.js`_
- `Watch Data Change With Options via Go and Vue.js`_

`Chrome Extension`_:

- `[Golang/GopherJS] Chrome Extension for Chinese Conversion`_
- `Get Current Tab URL From Chrome Extension in Go`_

**References**:

- `Synonyms - Go and JavaScript`_

.. adsu:: 3

----

References:

.. [1] `[Go WebAssembly] First Wasm Program - Hello World`_

.. [2] `GitHub - lpereira/gomoku: Gomoku compiles Go code to C++ <https://github.com/lpereira/gomoku>`_

.. [3] | `Introducing Joy: translate Go to JavaScript : golang <https://www.reddit.com/r/golang/comments/7jby77/introducing_joy_translate_go_to_javascript/>`_
       | `GitHub - matthewmueller/joy: A delightful Go to Javascript compiler <https://github.com/matthewmueller/joy>`_

.. [4] `GitHub - siongui/gojianfan: Traditional and Simplified Chinese Conversion in Go <https://github.com/siongui/gojianfan>`_

.. [5] | `Jsgo Playground: Edit and run Go code in the browser, supporting arbitrary import paths. : golang <https://www.reddit.com/r/golang/comments/8gxy9u/jsgo_playground_edit_and_run_go_code_in_the/>`_
       | `GitHub - dave/jsgo: GopherJS compiler, serving framework and CDN. <https://github.com/dave/jsgo>`_
       | `Jsgo Playground <https://play.jsgo.io/>`_
       | `GitHub - dave/play: jsgo playground: edit and run Go code in the browser, supporting arbitrary import paths <https://github.com/dave/play>`_

.. [6] `Gopherjs or webassembly ? : golang <https://old.reddit.com/r/golang/comments/9sb7fi/gopherjs_or_webassembly/>`_
.. [7] `Future of GopherJS and Go in the browser : golang <https://old.reddit.com/r/golang/comments/a68oop/future_of_gopherjs_and_go_in_the_browser/>`_
.. [8] `Pure Go web development? : golang <https://old.reddit.com/r/golang/comments/a8bbgq/pure_go_web_development/>`_

.. [9] | `Running Golang on the browser with WebAssembly and TinyGo : golang <https://old.reddit.com/r/golang/comments/fszeix/running_golang_on_the_browser_with_webassembly/>`_
       | `Running Golang on the browser with WebAssembly and TinyGo – Mariano Gappa's Blog <https://marianogappa.github.io/software/2020/04/01/webassembly-tinygo-cheesse/>`_

.. [10] | `Unify GopherJS's \`syscall/js\` and GopherWasm · Issue #899 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/899>`_
        | `GitHub - gopherjs/gopherwasm: This package is deprecated. Use syscall/js of GopherJS instead. <https://github.com/gopherjs/gopherwasm>`_
        | `Add syscall/js by hajimehoshi · Pull Request #908 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/pull/908>`_
        | `Build tag to indicate syscall/js support? · Issue #920 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/920>`_
        | `syscall/js: support for Go 1.14 changes · Issue #962 · gopherjs/gopherjs · GitHub <https://github.com/gopherjs/gopherjs/issues/962>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _official website of Go: https://golang.org/
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _Bulma: https://bulma.io/
.. _Vue.js: https://vuejs.org/
.. _Go Playground: https://play.golang.org/
.. _godom: https://github.com/siongui/godom
.. _First Frontend Go Program - Hello World in Browser: {filename}first-frontend-go-program-hello-world%en.rst
.. _JavaScript new Keyword in Go: {filename}js-new-keyword-in-go%en.rst
.. _Synonyms - Go and JavaScript: {filename}synonyms-go-and-javascript%en.rst
.. _querySelector, querySelectorAll, getElementById in Go: {filename}querySelector-querySelectorAll-getElementById-in-go%en.rst
.. _innerHTML and textContent in Go: {filename}innerHTML-textContent-in-go%en.rst
.. _Event Binding - addEventListener in Go: {filename}addEventListener-event-binding-in-go%en.rst
.. _Keyboard Event (Arrow Keys) in Go: {filename}keyboard-event-arrow-key-in-go%en.rst
.. _Create and Append Element or Text Node in Go: {filename}create-and-append-element-or-text-node-in-go%en.rst
.. _XMLHttpRequest (XHR) in Go: {filename}xmlhttprequest-xhr-in-go%en.rst
.. _HTML data-* Attribute in Go: {filename}html-data-attribute-in-go%en.rst
.. _HTML Element style Property in Go: {filename}html-element-style-property-in-go%en.rst
.. _HTML Element classList Property in Go: {filename}html-element-classlist-property-in-go%en.rst
.. _setTimeout method in Go: {filename}settimeout-in-go%en.rst
.. _Detect Browser Language Preference in Go: {filename}detect-browser-language-preference-in-go%en.rst
.. _JavaScript null Check in Go: {filename}js-null-test-in-go%en.rst
.. _JavaScript undefined Check in Go: {filename}js-undefined-test-in-go%en.rst
.. _Show CSS Loader While Resource Loading in Go: {filename}show-css-loader-while-resource-loading-in-go%en.rst
.. _Element Position (Scroll Included) in Go: {filename}element-position-scroll-included-in-go%en.rst
.. _Tooltip in Go: {filename}tooltip-in-go%en.rst
.. _Tooltip with Close Delay in Go: {filename}tooltip-with-close-delay-in-go%en.rst
.. _Convert Text to HTML Link in Go: {filename}convert-text-to-html-link-in-go%en.rst
.. _[GopherJS] WebSocket Client Example With Echo Server: {filename}../../../2017/05/18/go-websocket-client-example-with-echo-server%en.rst
.. _Bulma Navbar with Go Toggle: {filename}bulma-navbar-with-go-toggle%en.rst
.. _Toggle (Show/Hide) HTML Element via Go and Vue.js: {filename}toggle-dom-element-with-gopherjs-vue%en.rst
.. _Show keyCode of Pressed Key via Go and Vue.js: {filename}show-keyCode-of-pressed-key-via-gopherjs-vue%en.rst
.. _Bulma Dropdown with Go Toggle: {filename}bulma-dropdown-with-go-toggle%en.rst
.. _Chrome Extension: https://developer.chrome.com/extensions/getstarted
.. _[Golang/GopherJS] Chrome Extension for Chinese Conversion: {filename}/articles/2017/04/30/go-gopherjs-chrome-extension-for-chinese-translation%en.rst
.. _Get Current Tab URL From Chrome Extension in Go: {filename}tab-url-chrome-extension-in-go%en.rst
.. _Bulma Modal with Go Toggle: {filename}bulma-modal-with-go-toggle%en.rst
.. _Virtual Keyboard via Go and Vue.js: {filename}virtual-keypad-via-gopherjs-vue%en.rst
.. _Online Lemoine’s Conjecture Demo in Go: {filename}online-lemoine-conjecture-demo-in-go%en.rst
.. _Check if Value of HTML Input Text Field is Integer in Go: {filename}check-if-html-input-text-value-is-integer-in-go%en.rst
.. _Online Sieve of Eratosthenes Demo via Go and Vue.js: {filename}sieve-of-eratosthenes-via-gopherjs-vue%en.rst
.. _Watch Data Change via Go and Vue.js: {filename}watch-data-change-via-gopherjs-vue%en.rst
.. _Watch Data Change With Options via Go and Vue.js: {filename}watch-data-change-with-option-via-gopherjs-vue%en.rst
.. _Go 1.11: https://blog.golang.org/go1.11
.. _[Go WebAssembly] First Wasm Program - Hello World: {filename}/articles/2018/07/30/golang-wasm-hello-world%en.rst
.. _[Go WebAssembly] querySelector Example: {filename}/articles/2018/07/30/golang-wasm-querySelector%en.rst
.. _[Go WebAssembly] querySelectorAll Example: {filename}/articles/2018/07/30/golang-wasm-querySelectorAll%en.rst
.. _Load and Run Go WebAssembly Module: {filename}/articles/2018/07/30/load-and-run-golang-wasm-code%en.rst
.. _[Go WebAssembly] Event Binding - addEventListener Example: {filename}/articles/2018/07/30/golang-wasm-addEventListener-event-binding%en.rst
.. _[Go WebAssembly] XMLHttpRequest (XHR): {filename}/articles/2018/07/30/golang-wasm-xmlhttprequest-xhr%en.rst
.. _Simulate Enter Key Pressed in Go: {filename}simulate-enter-key-pressed-in-go%en.rst
