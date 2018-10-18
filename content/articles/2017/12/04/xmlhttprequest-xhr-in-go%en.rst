XMLHttpRequest (XHR) in Go
##########################

:date: 2017-12-14T23:28+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       XMLHttpRequest, Go net/http, Goroutine, CORS, HTTP GET
:category: Frontend Programming in Go
:summary: Send data to servers or retrieve data from servers on browsers - Use
          *XMLHttpRequest* method in Go/GopherJS.
:og_image: https://www.codeproject.com/KB/ajax/BackboneOfAjax/flow.jpg
:adsu: yes


This post shows how to interact with servers from browsers. Send data to URL or
retrieve data from URL from your browser via XMLHttpRequest_ method in
Go/GopherJS.
The full code example of this post is `on my GitHub`_.

.. contents:: **Table of Content**

Retrieve Data From URL
======================

In this example, we will show you how to get data from URL. We will get JSON
data from a URL and print the JSON data on the browser console.


JavaScript
++++++++++

.. code-block:: javascript

  var URL = "https://siongui.github.io/xemaauj9k5qn34x88m4h/sacca.json";

  function GetJSON() {
    console.log(this.responseText);
  }

  var req = new XMLHttpRequest();
  req.addEventListener("load", GetJSON);
  req.open("GET", URL);
  req.send();

This example is modified from `example on MDN`_. We get JSON file from the URL
and print the data on the console.

.. adsu:: 2


GopherJS
++++++++

The above code in Go/GopherJS is as follows:

.. code-block:: go

  package main

  import (
  	"github.com/gopherjs/gopherjs/js"
  )

  var URL = "https://siongui.github.io/xemaauj9k5qn34x88m4h/sacca.json"

  func GetJSON(url string) {
  	req := js.Global.Get("XMLHttpRequest").New()
  	req.Call("addEventListener", "load", func() {
  		println(req.Get("responseText").String())
  	})
  	req.Call("open", "GET", url, true)
  	req.Call("send")
  }

  func main() {
  	GetJSON(URL)
  }

.. rubric:: `Run Code on GopherJS Playground <https://gopherjs.github.io/playground/#/Q_zZMcJVK7>`__
   :class: align-center

.. adsu:: 3


Idomatic Go without import GopherJS
+++++++++++++++++++++++++++++++++++

Thanks to the great jobs done by GopherJS guys, you can use methods from
*net/http* package in Go standard library to retrieve data from a URL, just like
you are writing a local Go program. GopherJS will compile the code and translate
the Go code to the corresponding JavaScript XMLHttpRequest for you! You do not
even have to know about XMLHttpRequest!

.. code-block:: go

  package main

  import (
  	"bytes"
  	"net/http"
  )

  var URL = "https://siongui.github.io/xemaauj9k5qn34x88m4h/sacca.json"

  func GetJSON(url string) {
  	resp, err := http.Get(url)
  	if err != nil {
  		return
  	}
  	defer resp.Body.Close()
  	if resp.StatusCode != 200 {
  		return
  	}

  	buf := new(bytes.Buffer)
  	buf.ReadFrom(resp.Body)
  	println(buf.String())
  }

  func main() {
  	GetJSON(URL)
  }

.. rubric:: `Run Code on GopherJS Playground <https://gopherjs.github.io/playground/#/iC-_yZM_iJ>`__
   :class: align-center

The result of above code is the same as the Go code in previous section. Just
compile the code with GopherJS, and the JavaScript code output from GopherJS
will run with the same result! Amazing!

.. adsu:: 4


----

References:

.. [1] `GopherJS XMLHttpRequest (XHR) and MakeFunc Example <{filename}../../../2016/02/18/gopherjs-XMLHttpRequest-XHR-and-MakeFunc-example%en.rst>`_
.. [2] `[Golang] XMLHttpRequest (XHR) HTTP POST JSON Data by GopherJS <{filename}../../../2016/01/21/go-xhr-http-post-json-by-gopherjs%en.rst>`_
.. [3] `[Golang] XMLHttpRequest (XHR) HTTP GET JSON Data by GopherJS <{filename}../../../2016/01/20/go-xhr-http-get-json-by-gopherjs%en.rst>`_
.. [4] `[Golang] Access HTTP Request Header by XHR getAllResponseHeaders() <{filename}../../../2016/01/25/go-http-request-header-by-xhr-getAllResponseHeaders%en.rst>`_
.. [5] `[Golang] Caveats of GopherJS Development <{filename}../../../2016/02/07/go-caveat-of-gopherjs-development%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Go: https://golang.org/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/008-xmlhttprequest-xhr
.. _XMLHttpRequest: https://www.google.com/search?q=XMLHttpRequest
.. _example on MDN: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/Using_XMLHttpRequest
