[Go WebAssembly] querySelectorAll Example
#########################################

:date: 2018-10-05T22:14+08:00
:tags: Go, Golang, Go WebAssembly, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Go WebAssembly *querySelectorAll* example.
:og_image: https://golang.org/doc/gopher/gophercolor.png
:adsu: yes


In this post, we will show you how to use querySelectorAll_ method in Go wasm.
godom_ package is used for DOM manipulation instead of using `syscall/js`_
directly. First install godom_:

.. code-block:: go

  $ GOARCH=wasm GOOS=js go get -u github.com/siongui/godom/wasm

We will select div elements by querySelectorAll_, and set the value of its
*innerHTML* to **hi**. See demo first:

.. rubric:: `Go Wasm querySelectorAll Demo <https://siongui.github.io/frontend-programming-in-go/wasm/003-querySelectorAll/demo/>`__
   :class: align-center

The following is source code of demo:

**Go**:

.. code-block:: go

  package main

  import (
  	. "github.com/siongui/godom/wasm"
  )

  func main() {
  	testdivs := Document.QuerySelectorAll("#testdivs > div")
  	for _, testdiv := range testdivs {
  		testdiv.Set("innerHTML", "hi")
  	}
  }

If you have experience of JavaScript DOM manipulation, the above code is
self-evident. If you do not know how to compile the code, see [1]_.


**HTML**:

.. code-block:: html

  <!doctype html>
  <html>
  <head>
    <meta charset="utf-8">
    <title>Go wasm - querySelectorAll</title>
  </head>
  <body>
    <div id="testdivs">
      <div></div>
      <div></div>
      <div></div>
    </div>

    <!-- https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js -->
    <script src="wasm_exec.js"></script>
    <script>
  	const go = new Go();
  	let mod, inst;
  	WebAssembly.instantiateStreaming(
  		fetch("query.wasm", {cache: 'no-cache'}), go.importObject).then((result) => {
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

We manipulate the *div* elements inside the element whose id is *testdivs*.
Most of the HTML code is to load compiled wasm module. If you have no idea what
it means, see [1]_.

For querySelector_ example, see [2]_.

.. adsu:: 2

The full source code is also available `in my GitHub repo`_.

----

Tested on:

- ``Ubuntu Linux 18.04``
- ``Go 1.11.1``
- ``Chromium Version 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

References:

.. [1] `[Go WebAssembly] First Wasm Program - Hello World <{filename}golang-wasm-hello-world%en.rst>`_
.. [2] `[Go WebAssembly] querySelector Example <{filename}golang-wasm-querySelector%en.rst>`_

.. _querySelector: https://duckduckgo.com/?q=querySelector
.. _querySelectorAll: https://duckduckgo.com/?q=querySelectorAll
.. _Go Playground: https://play.golang.org/
.. _godom: https://github.com/siongui/godom
.. _WebAssembly: https://duckduckgo.com/?q=webassembly
.. _syscall/js: https://tip.golang.org/pkg/syscall/js/
.. _in my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/wasm/003-querySelectorAll
