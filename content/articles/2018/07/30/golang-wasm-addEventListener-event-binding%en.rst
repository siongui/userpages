[Go WebAssembly] Event Binding - addEventListener Example
#########################################################

:date: 2018-10-16T23:21+08:00
:tags: Go, Golang, Go WebAssembly, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Go WebAssembly Event Binding - *addEventListener* example.
:og_image: https://golang.org/doc/gopher/gophercolor.png
:adsu: yes


In this post, we will show you how to use addEventListener_ method in Go wasm.

First install godom_ to help DOM manipulation:

.. code-block:: go

  $ GOARCH=wasm GOOS=js go get -u github.com/siongui/godom/wasm

Click the text in the following demo several times to see what happens.

.. rubric:: `Go Wasm addEventListener Demo <https://siongui.github.io/frontend-programming-in-go/wasm/004-addEventListener/demo/>`__
   :class: align-center

The following is source code of demo:

**Go**:

.. code-block:: go

  package main

  import (
  	"fmt"
  	"syscall/js"

  	. "github.com/siongui/godom/wasm"
  )

  var signal = make(chan int)

  func keepAlive() {
  	for {
  		<-signal
  	}
  }

  func main() {
  	foo := Document.GetElementById("foo")
  	count := 0

  	cb := js.NewCallback(func(args []js.Value) {
  		count++
  		foo.Set("textContent", fmt.Sprintf("I am clicked %d time", count))
  	})
  	foo.Call("addEventListener", "click", cb)

  	keepAlive()
  }

We use NewCallback_ method in `syscall/js`_ package to create the callback for
the click event of the *div* element. The most tricky and strange part of the
above code is the use of *keepAlive* func to prevent *main* func from exiting.
If you do not do this, you will see the following error message in the developer
console when clicking on the text in the demo:

.. code-block:: txt

  bad callback: Go program has already exited

In the beginning I did not use *keepAlive* func and always saw the above error
message when clicking on the text. After some searches, I found the answer in
[2]_, which uses *keepAlive* func to prevent *main* from exiting.


**HTML**:

.. code-block:: html

  <!doctype html>
  <html>
  <head>
    <meta charset="utf-8">
    <title>Go wasm - addEventListener</title>
  </head>
  <body>
    <div id="foo">Click Here</div>

    <!-- https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js -->
    <script src="wasm_exec.js"></script>
    <script>
  	const go = new Go();
  	let mod, inst;
  	WebAssembly.instantiateStreaming(
  		fetch("click.wasm", {cache: 'no-cache'}), go.importObject).then((result) => {
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

Nothing special in above HTML code.
Most of the HTML code is to load compiled wasm module. If you have no idea what
it means, see [1]_.

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
.. [2] | `Lazy Hacker Babble: Go with WebAssembly Early Examples <https://blog.lazyhacker.com/2018/06/go-with-webassembly-early-examples.html>`_
       | `go wasm event binding at DuckDuckGo <https://duckduckgo.com/?q=go+wasm+event+binding>`_
.. [3] `Event Binding - addEventListener in Go <{filename}/articles/2017/12/04/addEventListener-event-binding-in-go%en.rst>`_ (GopherJS)

.. _addEventListener: https://duckduckgo.com/?q=addEventListener
.. _Go Playground: https://play.golang.org/
.. _godom: https://github.com/siongui/godom
.. _WebAssembly: https://duckduckgo.com/?q=webassembly
.. _syscall/js: https://tip.golang.org/pkg/syscall/js/
.. _in my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/wasm/004-addEventListener
.. _NewCallback: https://godoc.org/syscall/js#NewCallback
