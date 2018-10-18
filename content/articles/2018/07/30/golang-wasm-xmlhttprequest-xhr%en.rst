[Go WebAssembly] XMLHttpRequest (XHR)
#####################################

:date: 2018-10-19T01:47+08:00
:tags: Go, Golang, Go WebAssembly, Frontend Programming in Go, XMLHttpRequest,
       Go net/http, CORS, HTTP GET
:category: Frontend Programming in Go
:summary: Go WebAssembly *XMLHttpRequest (XHR)* - Issue HTTP requests to
          exchange data between browsers (client) and servers.
:og_image: https://golang.org/doc/gopher/gophercolor.png
:adsu: yes


This post shows how to interact with servers from browsers. Retrieve data from
URL from your browser via Go WebAssembly.
See demo first:

.. rubric:: `Go Wasm XMLHttpRequest (XHR) Demo <https://siongui.github.io/frontend-programming-in-go/wasm/005-xmlhttprequest-xhr/demo/>`__
   :class: align-center

Using XMLHttpRequest_ is the most common way to issue HTTP requests for data
exchange between browsers and servers. In Go Wasm, we do not need to deal with
XMLHttpRequest_ at all for data exchange. We can write idomatic Go code and use
`net/http`_ package in Go standard library to issue HTTP requests. You do not
even have to know about XMLHttpRequest_!

The following is the source code of the demo:

**Go**:

.. code-block:: go

  package main

  import (
  	"bytes"
  	"fmt"
  	"net/http"
  )

  func GetJson(url string) (json string, err error) {
  	resp, err := http.Get(url)
  	if err != nil {
  		return
  	}
  	defer resp.Body.Close()

  	if resp.StatusCode != 200 {
  		err = fmt.Errorf("response status code: %d", resp.StatusCode)
  		return
  	}

  	var buf bytes.Buffer
  	_, err = buf.ReadFrom(resp.Body)
  	if err != nil {
  		return
  	}
  	json = buf.String()
  	return
  }

  func main() {
  	json, err := GetJson("https://siongui.github.io/xemaauj9k5qn34x88m4h/sacca.json")
  	if err != nil {
  		panic(err)
  	}
  	fmt.Println(json)
  }

**HTML**:

.. code-block:: html

  <!doctype html>
  <html>
  <head>
    <meta charset="utf-8">
    <title>Go wasm - XMLHttpRequest (XHR)</title>
  </head>
  <body>
    Please open developer console to see the retrieved JSON data from server.

    <!-- https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js -->
    <script src="wasm_exec.js"></script>
    <script>
  	const go = new Go();
  	let mod, inst;
  	WebAssembly.instantiateStreaming(
  		fetch("xhr.wasm", {cache: 'no-cache'}), go.importObject).then((result) => {
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
.. [2] `XMLHttpRequest (XHR) in Go <{filename}/articles/2017/12/04/xmlhttprequest-xhr-in-go%en.rst>`_ (GopherJS)

.. _XMLHttpRequest: https://duckduckgo.com/?q=XMLHttpRequest
.. _net/http: https://golang.org/pkg/net/http/
.. _in my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/wasm/005-xmlhttprequest-xhr
