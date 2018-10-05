Load and Run Go WebAssembly Module
##################################

:date: 2018-10-06T04:34+08:00
:tags: Go, Golang, Go WebAssembly, Frontend Programming in Go
:category: Frontend Programming in Go
:summary: Show how to load and run Go WebAssembly code.
:og_image: https://golang.org/doc/gopher/gophercolor.png
:adsu: yes


When I tried to write and run my first Go_ WebAssembly_ module, it seemed to be
a easy task. I searched `load wasm module`_ that led me to the tutorial on MDN
[1]_. It looks easy, just use *WebAssembly.instantiateStreaming()* method and
everything is done. But life is not so easy. It did not work. So I took a look
at what other people did in their first Go wasm try (see references in [2]_) and
found out the solution.

To load and run compiled Go wasm module, we need a JavaScript loader provided in
Go source. We can copy the JavaScript loader to current directory by the
following shell command:

.. code-block:: bash

  $ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

You can also download *wasm_exec.js* from `Go GitHub repo`_.

Next, put *wasm_exec.js*, compiled Go wasm module (assume the name of the module
is *mymodule.wasm*) and HTML file in the same directory. Insert the following
code in the HTML file.

.. code-block:: html

  <script src="wasm_exec.js"></script>
  <script>
  	const go = new Go();
  	let mod, inst;
  	WebAssembly.instantiateStreaming(
  		fetch("mymodule.wasm", {cache: 'no-cache'}), go.importObject).then((result) => {
  		mod = result.module;
  		inst = result.instance;
  		run();
  	});
  	async function run() {
  		await go.run(inst);
  	};
  </script>

Open the HTML with your browser and now you will see Go wasm code loaded and
running!

.. adsu:: 2

----

Tested on:

- ``Ubuntu Linux 18.04``
- ``Go 1.11.1``
- ``Chromium Version 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

References:

.. [1] `Loading and running WebAssembly code - WebAssembly | MDN <https://developer.mozilla.org/en-US/docs/WebAssembly/Loading_and_running>`_
.. [2] `[Go WebAssembly] First Wasm Program - Hello World <{filename}golang-wasm-hello-world%en.rst>`_
.. [3] `WebAssembly · golang/go Wiki · GitHub <https://github.com/golang/go/wiki/WebAssembly>`_

.. _Go: https://golang.org/
.. _WebAssembly: https://duckduckgo.com/?q=webassembly
.. _load wasm module: https://duckduckgo.com/?q=load+wasm+module
.. _Go GitHub repo: https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js
