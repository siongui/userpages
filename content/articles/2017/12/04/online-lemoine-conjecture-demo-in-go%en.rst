Online Lemoine’s Conjecture Demo in Go
######################################

:date: 2018-04-24T23:06+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Math, Algorithm, Prime Number, Type Casting, Type Conversion
:category: Frontend Programming in Go
:summary: 2n + 1 = p + 2q always has a solution in primes p and q (not
          necessarily distinct) for n > 2. This online demo finds p and q for
          given odd number greater than 5. The demo is written in Go and
          compiled to JavaScript using GopherJS.
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/b/b9/Lemoine.jpg/250px-Lemoine.jpg
:adsu: yes


Lemoine’s conjecture says 2n + 1 = p + 2q always has a solution in primes p and
q (not necessarily distinct) for n > 2. This online demo finds p and q for given
odd number greater than 5. The demo is written in Go and compiled to JavaScript
using GopherJS.

.. rubric:: `Demo <{filename}/articles/2018/04/22/online-lemoine-conjecture-demo%en.rst>`_
   :class: align-center

The following are source code.

**HTML**

.. code-block:: html

  Input Odd Number Greater Than 5: <input type="text" name="n"><br><br>
  <button type="button" id="getLemoine">Get Lemoine’s Conjecture</button><br><br>
  <textarea id="result" rows="20" cols="50"></textarea>

**Go (lemoine.go)**: This file can be used not only online, but also in your
local program. This is the power of GopherJS, which compiles almost all local
Go sources without I/O to JavaScript, and makes your code shareable both on
frontend and backend.

.. code-block:: go

  package main

  import (
  	"errors"
  )

  func IsPrime(n int) bool {
  	if n < 2 {
  		return false
  	}

  	for i := 2; i*i <= n; i++ {
  		if n%i == 0 {
  			return false
  		}
  	}
  	return true
  }

  func Lemoine(n int) (map[int]int, error) {
  	if n <= 5 || n%2 == 0 {
  		return nil, errors.New("n must be greater than 5 and must be odd")
  	}

  	m := make(map[int]int)

  	for q := 1; q <= n/2; q++ {
  		p := n - 2*q

  		if IsPrime(p) && IsPrime(q) {
  			m[p] = q
  		}
  	}
  	return m, nil
  }

**Go (appdom.go)**: This file use *github.com/siongui/godom*. Please *go get* it
before compiling this file.

.. code-block:: go

  package main

  import (
  	. "github.com/siongui/godom"
  	"strconv"
  )

  func main() {
  	elmn := Document.QuerySelector("input[name='n']")
  	btn := Document.QuerySelector("#getLemoine")
  	resultElm := Document.QuerySelector("#result")

  	btn.AddEventListener("click", func(e Event) {
  		n, err := strconv.Atoi(elmn.Value())
  		if err != nil {
  			resultElm.SetValue(err.Error())
  			return
  		}
  		pqs, err := Lemoine(n)
  		if err != nil {
  			resultElm.SetValue(err.Error())
  			return
  		}

  		text := ""
  		for p, q := range pqs {
  			text += elmn.Value() + " = " + strconv.Itoa(p) + " + ( 2 * " + strconv.Itoa(q) + " )\n"
  			resultElm.SetValue(text)
  		}
  	})
  }

The JavaScript equivalent is also available on the demo page [1]_.

The link of full source code is available `on my GitHub repo`_.

.. adsu:: 2

Tested on:

- ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``
- ``Go 1.10.1``
- ``GopherJS 1.10-3``

----

References:

.. [1] `Online Lemoine’s Conjecture Demo <{filename}/articles/2018/04/22/online-lemoine-conjecture-demo%en.rst>`_

.. _on my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/028-lemoine-conjecture
