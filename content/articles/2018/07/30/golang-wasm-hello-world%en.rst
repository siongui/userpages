[Go WebAssembly] First Wasm Program - Hello World
#################################################

:date: 2018-07-30T00:17+08:00
:modified: 2018-10-04T23:13+08:00
:tags: Go, Golang, Go WebAssembly, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: First try of Go WebAssembly - Say *Hello World* via *alert* method.
:og_image: https://golang.org/doc/gopher/gophercolor.png
:adsu: yes


It is really exciting moment to be able to use Go to do frontend programming.
Start from Go 1.11, WebAssembly_ is `supported officially in Go`_. At the time
of this writing, Go 1.11 is not released yet. You can try with go1.11beta2_.

As usual, we will say *Hello World* in our first Go wasm program. Go 1.11
provides `syscall/js`_ package to call into JavaScript. Since the author of Go
wasm and GopherJS_ is the same person, it is no surprise that the way to call
into JavaScript in Go wasm is similar to that in GopherJS. If you have
experiences about GopherJS, you can turn your Go/GopherJS code into Go wasm code
just by a few replacement of words in the code. If you want to learn more about
GopherJS, search *GopherJS* in this site.

Because of readability, I am not going to use `syscall/js`_ package directly in
my code. I update my godom_ package, a wrapper package for both GopherJS/Go
WebAssembly, which make code of DOM manipulation look more like JavaScript. You
can install the wrapper for Go wasm by:

.. code-block:: go

  $ GOARCH=wasm GOOS=js go get -u github.com/siongui/godom/wasm

Next, we use alert_ method to display an alert box which says *Hello World*. It
is easy -

.. code-block:: go

  package main

  import (
  	. "github.com/siongui/godom/wasm"
  )

  func main() {
  	Window.Alert("hello world!")
  }

We will compile the above code into wasm module. Save above code as *alert.go*
and then run the following command:

.. code-block:: go

  $ GOARCH=wasm GOOS=js go build -o alert.wasm alert.go

So far so good, we have our wasm module ready, but how to load the module via
JavaScript? This is the most tricky part in my opinion. You need to download the
official `wasm_exec.js`_, which loads Go wasm module for you. The following
HTML/JavaScript show you how to use the loader:

.. code-block:: html

  <!doctype html>
  <html>
  <head>
    <meta charset="utf-8">
    <title>Go wasm - Hello World</title>
  </head>
  <body>
    <script src="wasm_exec.js"></script>
    <script>
  	const go = new Go();
  	let mod, inst;
  	WebAssembly.instantiateStreaming(
  		fetch("alert.wasm", {cache: 'no-cache'}), go.importObject).then((result) => {
  		mod = result.module;
  		inst = result.instance;
  		run();
  	});
  	async function run() {
  		await go.run(inst);
  	};
    </script>
  </body>
  </html>

Put *alert.wasm*, *wasm_exec.js* and above HTML together, open the HTML file
with your browser. Visit the following demo to see the final result -

.. rubric:: `Go Wasm Hello World Demo <https://siongui.github.io/frontend-programming-in-go/wasm/001-hello-world/demo/>`__
   :class: align-center

The full source code is also available `in my GitHub repo`_.

.. adsu:: 2

----

Tested on:

- ``Ubuntu Linux 18.04``
- ``Go 1.11 beta2``
- ``Chromium Version 67.0.3396.99 on Ubuntu 18.04 (64-bit)``/
  ``Firefox Quantum 61.0.1 (64-bit) on Ubuntu 18.04``

----

References:

.. [1] | `meta: WebAssembly ("wasm") support · Issue #18892 · golang/go · GitHub <https://github.com/golang/go/issues/18892>`_
       | `Is there a webassembly compiler for Go in the works? : golang <https://www.reddit.com/r/golang/comments/5yl984/is_there_a_webassembly_compiler_for_go_in_the/>`_
       |
       | `Go and wasm: Advent Day 17 : golang <https://www.reddit.com/r/golang/comments/7ke47z/go_and_wasm_advent_day_17/>`_
       | `Go and wasm: generating and executing wasm with Go GopherAcademy <https://blog.gopheracademy.com/advent-2017/go-wasm/>`_
       |
       | `WebAssembly architecture for Go <https://docs.google.com/document/d/131vjr4DH6JFnb-blm_uRdaC0_Nv3OUwjEY5qVCxCup4/edit>`_
       | `WebAssembly architecture for Go : golang <https://www.reddit.com/r/golang/comments/81dt49/webassembly_architecture_for_go/>`_
       |
       | `WebAssembly support lands in Go language: golang wasm/js <https://react-etc.net/entry/webassembly-support-lands-in-go-language-golang-wasm-js>`_
       | `WebAssembly support lands in Go language: golang wasm/js : golang <https://www.reddit.com/r/golang/comments/8c64ix/webassembly_support_lands_in_go_language_golang/>`_
       |
       | `Notes on WASM in Go 1.11 : golang <https://www.reddit.com/r/golang/comments/8q04c8/notes_on_wasm_in_go_111/>`_
       | `Some notes about the upcoming WebAssembly support in Go - Unladen swallow <https://blog.owulveryck.info/2018/06/08/some-notes-about-the-upcoming-webassembly-support-in-go.html>`_
       |
       | `Early experimentation with Go's WebAssembly : golang <https://www.reddit.com/r/golang/comments/8q71ed/early_experimentation_with_gos_webassembly/>`_
       | `Lazy Hacker Babble: Go with WebAssembly Early Examples <https://blog.lazyhacker.com/2018/06/go-with-webassembly-early-examples.html>`_
       |
       | `GopherJS vs WebAssembly for Go : golang <https://www.reddit.com/r/golang/comments/8rdtxi/gopherjs_vs_webassembly_for_go/>`_
       | `GopherJS vs WebAssembly for Go <https://dev.to/hajimehoshi/gopherjs-vs-webassembly-for-go-148m>`_
       |
       | `Experiments with image manipulation in WASM using Go : golang <https://www.reddit.com/r/golang/comments/8s07p5/experiments_with_image_manipulation_in_wasm_using/>`_
       | `Experiments with image manipulation in WASM using Go <http://agniva.me/wasm/2018/06/18/shimmer-wasm.html>`_
       |
       | `Web Assembly and Go: A look to the future : golang <https://www.reddit.com/r/golang/comments/8t2q1h/web_assembly_and_go_a_look_to_the_future/>`_
       | `Web Assembly and Go: A look to the future <https://brianketelsen.com/web-assembly-and-go-a-look-to-the-future/>`_
       | `Web Assembly and Go: A look to the future | Hacker News <https://news.ycombinator.com/item?id=17381816>`_
       | `Creating Web Component in Go + wasm <https://matthewphillips.info/programming/wasm-golang-ce.html>`_
       |
       | `GopherWasm - A wrapper for GopherJS (\`gopherjs/js\`) and Wasm (\`syscall/js\`) : golang <https://www.reddit.com/r/golang/comments/8tjtc1/gopherwasm_a_wrapper_for_gopherjs_gopherjsjs_and/>`_
       | `GitHub - gopherjs/gopherwasm: A wrapper for GopherJS (\`gopherjs/js\`) and Wasm (\`syscall/js\`) <https://github.com/gopherjs/gopherwasm>`_
       |
       | `Go and Goto in WebAssembly : golang <https://www.reddit.com/r/golang/comments/8usnpe/go_and_goto_in_webassembly/>`_
       | `Please Support Arbitrary Labels and Gotos. · Issue #796 · WebAssembly/design · GitHub <https://github.com/WebAssembly/design/issues/796#issuecomment-401310366>`_
       |
       | `go1.11 webassembly experiments : golang <https://www.reddit.com/r/golang/comments/8vrhld/go111_webassembly_experiments/>`_
       | `GitHub - stdiopt/gowasm-experiments: go1.11 webassembly experiments <https://github.com/stdiopt/gowasm-experiments>`_
       |
       | `Go 1.11: WebAssembly for the gophers : golang <https://www.reddit.com/r/golang/comments/8vgmpj/go_111_webassembly_for_the_gophers/>`_
       | `Go 1.11: WebAssembly for the gophers – Zenika <https://medium.zenika.com/go-1-11-webassembly-for-the-gophers-ae4bb8b1ee03>`_
       |
       | `Using Go for WebAssembly Applications : golang <https://www.reddit.com/r/golang/comments/8wi14y/using_go_for_webassembly_applications/>`_
       | `Using Go (Golang) for WebAssembly Applications · Sebastian Holstein <https://sebastian-holstein.de/post/2018-07-05-go-wasm-application/>`_
       |
       | `Go WebAssembly: Binding structures to JS references : golang <https://www.reddit.com/r/golang/comments/8wkd5r/go_webassembly_binding_structures_to_js_references/>`_
       | `Go WebAssembly: Binding structures to JS references <https://medium.com/@nlepage/go-webassembly-binding-structures-to-js-references-4eddd6fd4d23>`_

.. [2] | `Go1.11和WebAssembly <https://mp.weixin.qq.com/s/jqISsdQ9tDzy9Zg6g6u5xw>`_
       | `GitHub - chai2010/awesome-wasm-zh: WebAssembly(wasm)资源精选 <https://github.com/chai2010/awesome-wasm-zh>`_
       | `WebAssembly · golang/go Wiki · GitHub <https://github.com/golang/go/wiki/WebAssembly>`_

.. [3] `WebAssembly = Writing JS in Go? • r/golang <https://old.reddit.com/r/golang/comments/9gtg2h/webassembly_writing_js_in_go/>`_
.. [4] `A web Go REPL via WebAssembly : golang <https://old.reddit.com/r/golang/comments/9hsaui/a_web_go_repl_via_webassembly/>`_
.. [5] `Writing a Frontend Web Framework with WebAssembly And Go : golang <https://old.reddit.com/r/golang/comments/9tuuep/writing_a_frontend_web_framework_with_webassembly/>`_
.. [6] `Run JavaScript using Otto in Go in WASM in the Browser : golang <https://old.reddit.com/r/golang/comments/9velz5/run_javascript_using_otto_in_go_in_wasm_in_the/>`_
.. [7] `This is a WASM-based Vue.js wrapper written in Go : golang <https://old.reddit.com/r/golang/comments/9w4nk5/this_is_a_wasmbased_vuejs_wrapper_written_in_go/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _official website of Go: https://golang.org/
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _Bulma: https://bulma.io/
.. _Vue.js: https://vuejs.org/
.. _Go Playground: https://play.golang.org/
.. _godom: https://github.com/siongui/godom
.. _go1.11beta2: https://groups.google.com/forum/#!msg/golang-announce/RVR0FzIKBsU/PAxl4-ZVCAAJ
.. _WebAssembly: https://duckduckgo.com/?q=webassembly
.. _supported officially in Go: https://tip.golang.org/doc/go1.11#wasm
.. _go1.11beta2: https://groups.google.com/forum/#!msg/golang-announce/RVR0FzIKBsU/PAxl4-ZVCAAJ
.. _syscall/js: https://tip.golang.org/pkg/syscall/js/
.. _alert: https://www.w3schools.com/jsref/met_win_alert.asp
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _wasm_exec.js: https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js
.. _in my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/wasm/001-hello-world
