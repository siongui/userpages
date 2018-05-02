[Vue.js] Online Sieve of Eratosthenes Demo
##########################################

:date: 2018-05-02T23:33+08:00
:tags: Vue.js, JavaScript, Algorithm, Prime Number, Math
:category: Vue.js
:summary: *Sieve of Eratosthenes* is a simple and ancient method to find prime
          numbers up to a given limit. Given a limit, this online demo returns
          all prime number below the limit.
:og_image: http://www.texample.net/media/tikz/examples/PNG/eratosthenes-sieve.png
:adsu: yes


`Sieve of Eratosthenes`_ is a simple and ancient method to find prime numbers up
to a given limit. Given a limit, this online demo returns all prime numbers
below the limit. The UI interaction is written using Vue.js_, and the algorithm
comes from GeeksforGeeks [1]_ and implemented using JavaScript.

.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <div id="vueapp">
    positive integer limit: <input v-model="intn" placeholder="Input positive integer">
    <p>{{ primes }}</p>
  </div>

  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  function SieveOfEratosthenes(n) {
  	// Create a boolean array "prime[0..n]" and initialize
  	// all entries it as true. A value in prime[i] will
  	// finally be false if i is Not a prime, else true.
  	var integers = [];
  	for (var i = 2; i < n+1; i++) {
  		integers[i] = true;
  	}

  	for (var p = 2; p*p <= n; p++) {
  		// If integers[p] is not changed, then it is a prime
  		if (integers[p] == true) {
  			// Update all multiples of p
  			for (var i = p * 2; i <= n; i += p) {
  				integers[i] = false;
  			}
  		}
  	}

  	// return all prime numbers <= n
  	var primes = [];
  	for (var p = 2; p <= n; p++) {
  		if (integers[p] == true) {
  			primes.push(p);
  		}
  	}

  	return primes;
  }

  new Vue({
    el: '#vueapp',
    data: {
      intn: 5,
      primes: ""
    },
    watch: {
      intn: {
        immediate: true,
        handler(val, oldVal) {
          var n = parseInt(val);
          if (isNaN(n)) {
            this.primes = "please input integer";
            return;
          }

          this.primes = SieveOfEratosthenes(n);
        }
      }
    }
  });
  </script>

**HTML**:

.. code-block:: html

  <div id="vueapp">
    positive integer limit: <input v-model="intn" placeholder="Input positive integer">
    <p>{{ primes }}</p>
  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>


Given an input from user, we check if the input is an integer. Then we run the
sieve algorithm to check if the number is prime.


**JavaScript**:

.. code-block:: javascript

  'use strict';

  function SieveOfEratosthenes(n) {
  	// Create a boolean array "prime[0..n]" and initialize
  	// all entries it as true. A value in prime[i] will
  	// finally be false if i is Not a prime, else true.
  	var integers = [];
  	for (var i = 2; i < n+1; i++) {
  		integers[i] = true;
  	}

  	for (var p = 2; p*p <= n; p++) {
  		// If integers[p] is not changed, then it is a prime
  		if (integers[p] == true) {
  			// Update all multiples of p
  			for (var i = p * 2; i <= n; i += p) {
  				integers[i] = false;
  			}
  		}
  	}

  	// return all prime numbers <= n
  	var primes = [];
  	for (var p = 2; p <= n; p++) {
  		if (integers[p] == true) {
  			primes.push(p);
  		}
  	}

  	return primes;
  }

  new Vue({
    el: '#vueapp',
    data: {
      intn: 5,
      primes: ""
    },
    watch: {
      intn: {
        immediate: true,
        handler(val, oldVal) {
          var n = parseInt(val);
          if (isNaN(n)) {
            this.primes = "please input integer";
            return;
          }

          this.primes = SieveOfEratosthenes(n);
        }
      }
    }
  });

.. adsu:: 2

----

Tested on:

- ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``
- ``Vue.js 2.5.16``

----

References:

.. [1] `Sieve of Eratosthenes - GeeksforGeeks <https://www.geeksforgeeks.org/sieve-of-eratosthenes/>`_
.. [2] `[Golang] Sieve of Eratosthenes <{filename}/articles/2017/04/17/go-sieve-of-eratosthenes%en.rst>`_
.. [3] `[JavaScript] Sieve of Eratosthenes <{filename}/articles/2018/04/28/javascript-sieve-of-eratosthenes%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
