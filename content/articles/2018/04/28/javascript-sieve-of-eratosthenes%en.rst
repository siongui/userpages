[JavaScript] Sieve of Eratosthenes
##################################

:date: 2018-04-28T23:07+08:00
:tags: JavaScript, Algorithm, Prime Number, Math
:category: JavaScript
:summary: JavaScript implementation of *Sieve of Eratosthenes*.
:og_image: http://www.texample.net/media/tikz/examples/PNG/eratosthenes-sieve.png
:adsu: yes

JavaScript implementation of `Sieve of Eratosthenes`_. See from GeeksforGeeks
[1]_. The JavaScript code is ported from the C/C++ code of GeeksforGeeks.

.. raw:: html

  Input Integer Greater Than 2: <input type="text" name="n"><br><br>
  <button type="button" id="sieve">Run Sieve of Eratosthenes</button><br><br>
  <textarea id="result" rows="20" cols="50"></textarea>

  <script>
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

  var elmn = document.querySelector("input[name='n']");
  var btn = document.querySelector("#sieve");
  var resultElm = document.querySelector("#result");
  btn.addEventListener("click", function(e) {
    var n = parseInt(elmn.value);
    if (!Number.isInteger(n)) {
      resultElm.value = "input must be integer!";
      return;
    }
    if ( n<=2 ) {
      resultElm.value = "n must be greater than 2!";
      return;
    }
    result.value = SieveOfEratosthenes(n);
  });
  </script>


**Source Code (HTML)**:

.. code-block:: html

  Input Integer Greater Than 2: <input type="text" name="n"><br><br>
  <button type="button" id="sieve">Run Sieve of Eratosthenes</button><br><br>
  <textarea id="result" rows="20" cols="50"></textarea>


**Source Code (JavaScript)**:

.. code-block:: javascript

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

  var elmn = document.querySelector("input[name='n']");
  var btn = document.querySelector("#sieve");
  var resultElm = document.querySelector("#result");
  btn.addEventListener("click", function(e) {
    var n = parseInt(elmn.value);
    if (!Number.isInteger(n)) {
      resultElm.value = "input must be integer!";
      return;
    }
    if ( n<=2 ) {
      resultElm.value = "n must be greater than 2!";
      return;
    }
    result.value = SieveOfEratosthenes(n);
  });

.. adsu:: 2

Tested on: ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``

----

References:

.. [1] `Sieve of Eratosthenes - GeeksforGeeks <https://www.geeksforgeeks.org/sieve-of-eratosthenes/>`_
.. [2] `[Golang] Sieve of Eratosthenes <{filename}/articles/2017/04/17/go-sieve-of-eratosthenes%en.rst>`_
.. [3] `Online Lemoine’s Conjecture Demo <{filename}/articles/2018/04/22/online-lemoine-conjecture-demo%en.rst>`_

.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
