Check if Value of HTML Input Text Field is Integer in Go
########################################################

:date: 2018-04-25T23:04+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Type Casting, Type Conversion
:category: Frontend Programming in Go
:summary: Check if the value of HTML input text field is integer in Go. Compiled
          to JavaScript using GopherJS.
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/b/b9/Lemoine.jpg/250px-Lemoine.jpg
:adsu: yes


Check if the value of HTML input text field is integer in Go. Compiled to
JavaScript using GopherJS.

.. rubric:: `Demo <{filename}/articles/2018/04/23/javascript-check-if-input-text-value-is-integer%en.rst>`_
   :class: align-center

Note that above demo is written in JavaScript, and is not 100% the same as the
following demo written in Go. You have to compile the Go code to JavaScript
using GopherJS. The JavaScript demo is not very correct, because if you input
*2.1*, it will tell you it's integer *2* because of JavaScript *parseInt()*
method treats *2.1* as 2. The Go demo is correct, and it will tell you *2.1* is
invalid.

The following are source code in Go and HTML.

**HTML**

.. code-block:: html

  <input type="text" name="n">
  <button type="button" id="check">Check Integer</button><br><br>
  <div id="result"></div>


**Go (appdom.go)**: This file use *github.com/siongui/godom*. Please *go get* it
before compiling this file using GopherJS.

We use `strconv.Atoi`_ in Go standard library to convert the string to integer
for us. If the input is not integer, *error* will be returned.

.. code-block:: go

  package main

  import (
  	. "github.com/siongui/godom"
  	"strconv"
  )

  func main() {
  	elmn := Document.QuerySelector("input[name='n']")
  	btn := Document.QuerySelector("#check")
  	info := Document.QuerySelector("#result")

  	btn.AddEventListener("click", func(e Event) {
  		_, err := strconv.Atoi(elmn.Value())
  		if err != nil {
  			info.SetInnerHTML(err.Error())
  			return
  		}
  		info.SetInnerHTML("input is integer")
  	})
  }

The link of full source code is available `on my GitHub repo`_.

.. adsu:: 2

Tested on:

- ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``
- ``Go 1.10.1``
- ``GopherJS 1.10-3``

----

References:

.. [1] `[JavaScript] Check if Value of Input Text Field is Integer <{filename}/articles/2018/04/23/javascript-check-if-input-text-value-is-integer%en.rst>`_

.. _strconv.Atoi: https://golang.org/pkg/strconv/#Atoi
.. _on my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/029-input-text-integer-check
