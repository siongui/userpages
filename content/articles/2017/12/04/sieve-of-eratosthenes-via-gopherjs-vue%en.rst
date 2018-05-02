Online Sieve of Eratosthenes Demo via Go and Vue.js
###################################################

:date: 2018-05-01T23:51+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Vue.js, gopherjs-vue, Algorithm, Prime Number, Math
:category: Frontend Programming in Go
:summary: Online demo of *sieve of eratosthenes* via Go/GopherJS/gopherjs-vue.
:og_image: https://www.cut-the-knot.org/Curriculum/Arithmetic/Eratosthenes.jpg
:adsu: yes

Online `Sieve of Eratosthenes`_ demo via Go/GopherJS/gopherjs-vue_ (Go binding
for Vue.js_). `Sieve of Eratosthenes`_ is a method to find prime numbers up to a
given limit. The algorithm comes from GeeksforGeeks [1]_ and implemented in Go.
gopherjs-vue_ is used to build to UI interaction.

Although the following demo is written in JavaScript/Vue.js, the effect is the
same as that written in Go.

.. rubric:: `Demo <{filename}/articles/2018/05/02/vuejs-online-sieve-of-eratosthenes-demo%en.rst>`_
   :class: align-center

The link of full source code is available `on my GitHub repo`_.

**HTML**

.. code-block:: html

  <div id="vueapp">
    <input v-model="intn" placeholder="Input even integer > 2">
    <p>{{ result }}</p>
  </div>


**Go**: *go get* *github.com/gopherjs/gopherjs* and
*github.com/oskca/gopherjs-vue* before compiling the Go code.

.. code-block:: go

  package main

  import (
  	"strconv"

  	"github.com/gopherjs/gopherjs/js"
  	"github.com/oskca/gopherjs-vue"
  )

  type Model struct {
  	*js.Object        // this is needed for bidirectional data bindings
  	Intn       string `js:"intn"`
  	Result     string `js:"result"`
  }

  func SieveOfEratosthenes(n int) []int {
  	// Create a boolean array "prime[0..n]" and initialize
  	// all entries it as true. A value in prime[i] will
  	// finally be false if i is Not a prime, else true.
  	integers := make([]bool, n+1)
  	for i := 2; i < n+1; i++ {
  		integers[i] = true
  	}

  	for p := 2; p*p <= n; p++ {
  		// If integers[p] is not changed, then it is a prime
  		if integers[p] == true {
  			// Update all multiples of p
  			for i := p * 2; i <= n; i += p {
  				integers[i] = false
  			}
  		}
  	}

  	// return all prime numbers <= n
  	var primes []int
  	for p := 2; p <= n; p++ {
  		if integers[p] == true {
  			primes = append(primes, p)
  		}
  	}

  	return primes
  }

  func main() {
  	m := &Model{
  		Object: js.Global.Get("Object").New(),
  	}
  	// field assignment is required in this way to make data passing works
  	m.Intn = "4"
  	m.Result = ""

  	// create the VueJS viewModel using a struct pointer
  	app := vue.New("#vueapp", m)

  	opt := js.Global.Get("Object").New()
  	opt.Set("immediate", true)
  	app.Call("$watch", "intn", func(newVal, oldVal string) {
  		n, err := strconv.Atoi(newVal)
  		if err != nil {
  			m.Result = "input must be integer"
  			return
  		}
  		m.Result = ""

  		primes := SieveOfEratosthenes(n)
  		for _, prime := range primes {
  			m.Result += (strconv.Itoa(prime) + " ")
  		}
  	}, opt)
  }

The JavaScript equivalent is available on my previous post [2]_.

.. adsu:: 2

Tested on:

- ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``
- ``Go 1.10.1``
- ``GopherJS 1.10-3``

----

References:

.. [1] `Sieve of Eratosthenes - GeeksforGeeks <https://www.geeksforgeeks.org/sieve-of-eratosthenes/>`_
.. [2] `[JavaScript] Sieve of Eratosthenes <{filename}/articles/2018/04/28/javascript-sieve-of-eratosthenes%en.rst>`_
.. [3] `[Golang] Sieve of Eratosthenes <{filename}/articles/2017/04/17/go-sieve-of-eratosthenes%en.rst>`_
.. [4] `[Vue.js] Online Goldbach's Conjecture Demo <{filename}/articles/2018/04/29/vuejs-online-goldbach-conjecture-demo%en.rst>`_
.. [5] `vue - GoDoc <https://godoc.org/github.com/oskca/gopherjs-vue#ViewModel.Watch>`_
.. [6] `GitHub - lifeng1335/gopherjs-vue-examples: Vue.js official examples written in go language with gopherjs <https://github.com/lifeng1335/gopherjs-vue-examples>`_
.. [7] `vm.$watch( expOrFn, callback, [options] ) - API — Vue.js <https://vuejs.org/v2/api/#vm-watch>`_


.. _gopherjs-vue: https://github.com/oskca/gopherjs-vue
.. _Vue.js: https://vuejs.org/
.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
.. _on my GitHub repo: https://github.com/siongui/frontend-programming-in-go/tree/master/030-sieve-of-eratosthenes-gopherjs-vue
