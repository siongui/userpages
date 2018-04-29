[Vue.js] Online Goldbach's Conjecture Demo
##########################################

:date: 2018-04-29T23:40+08:00
:tags: Vue.js, JavaScript, Algorithm, Prime Number, Math
:category: Vue.js
:summary: Goldbach's conjecture - Every even integer greater than 2 can be
          written as the sum of two primes. Given a positive even integer, this
          online demo returns the two primes.
:og_image: https://i.ytimg.com/vi/ES4U-MCl3wY/maxresdefault.jpg
:adsu: yes


Goldbach's conjecture - Every even integer greater than 2 can be written as the
sum of two primes. Given a positive even integer, this online demo returns the
two primes. The UI interaction is written using Vue.js_, and the algorithm is
implemented using JavaScript.

.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <div id="vueapp">
    <input v-model="intn" placeholder="Input even integer > 2">
    <p>{{ result }}</p>
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

  function Goldbach(n) {
    var primes = SieveOfEratosthenes(n);
    for (var i=0; i < primes.length; i++) {
      var p1 = primes[i];
      var p2 = n - p1;
      if (primes.indexOf(p2) != -1) {
        return [p1, p2];
      }
    }
  }

  new Vue({
    el: '#vueapp',
    data: {
      intn: 4,
      result: ""
    },
    watch: {
      intn: {
        immediate: true,
        handler(val, oldVal) {
          var n = parseInt(val);
          if (isNaN(n)) {
            this.result = "please input integer";
            return;
          }
          if (n < 3) {
            this.result = "input integer must > 2";
            return;
          }
          if (n%2 == 1) {
            this.result = "input integer must be even";
            return;
          }
          var pp = Goldbach(n);
          this.result = (val + " = " + pp[0] + " + " + pp[1]);
        }
      }
    }
  });
  </script>

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <input v-model="intn" placeholder="Input even integer > 2">
    <p>{{ result }}</p>
  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>

Given an input from user, we check if the input is an integer, if the input is
greater than 2, and if the input is even. Then we use `Sieve of Eratosthenes`_
to find out all primes up to the user input integer, and find that the two
primes that sum of which is user input integer.


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

  function Goldbach(n) {
    var primes = SieveOfEratosthenes(n);
    for (var i=0; i < primes.length; i++) {
      var p1 = primes[i];
      var p2 = n - p1;
      if (primes.indexOf(p2) != -1) {
        return [p1, p2];
      }
    }
  }

  new Vue({
    el: '#vueapp',
    data: {
      intn: 4,
      result: ""
    },
    watch: {
      intn: {
        immediate: true,
        handler(val, oldVal) {
          var n = parseInt(val);
          if (isNaN(n)) {
            this.result = "please input integer";
            return;
          }
          if (n < 3) {
            this.result = "input integer must > 2";
            return;
          }
          if (n%2 == 1) {
            this.result = "input integer must be even";
            return;
          }
          var pp = Goldbach(n);
          this.result = (val + " = " + pp[0] + " + " + pp[1]);
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

.. [1] `[JavaScript] Sieve of Eratosthenes <{filename}/articles/2018/04/28/javascript-sieve-of-eratosthenes%en.rst>`_
.. [2] `[Golang] Goldbach's conjecture <{filename}/articles/2017/06/06/go-goldbach-conjecture%en.rst>`_
.. [3] `[Vue.js] Input Text Element Change Event <{filename}/articles/2017/02/03/vuejs-input-change-event%en.rst>`_
.. [4] `[JavaScript] Check if Value of Input Text Field is Integer <{filename}/articles/2018/04/23/javascript-check-if-input-text-value-is-integer%en.rst>`_
.. [5] `Watchers not triggered on initialization - Get Help - Vue Forum <https://forum.vuejs.org/t/watchers-not-triggered-on-initialization/12475>`_

.. _Vue.js: https://vuejs.org/
.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
