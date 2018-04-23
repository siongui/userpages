Online Lemoine’s Conjecture Demo
################################

:date: 2018-04-22T23:41+08:00
:tags: JavaScript, Math, Algorithm, Prime Number, Type Casting, Type Conversion
:category: Algorithm
:summary: 2n + 1 = p + 2q always has a solution in primes p and q (not
          necessarily distinct) for n > 2. This online demo finds p and q for
          given odd number greater than 5.
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/b/b9/Lemoine.jpg/250px-Lemoine.jpg
:adsu: yes


`Lemoine’s Conjecture`_ says any odd integer greater than 5 can be represented
as the sum of an odd prime number and an even semiprime. Another statement which
is suitable for programming is that 2n + 1 = p + 2q always has a solution in
primes p and q (not necessarily distinct) for n > 2. This online demo will find
p and q for given odd number greater than 5.

.. raw:: html

  Input Odd Number Greater Than 5: <input type="text" name="n"><br><br>
  <button type="button" id="getLemoine">Get Lemoine’s Conjecture</button><br><br>
  <textarea id="result" rows="20" cols="50"></textarea>

  <script>
  var elmn = document.querySelector("input[name='n']");
  var btn = document.querySelector("#getLemoine");
  var resultElm = document.querySelector("#result");
  btn.addEventListener("click", function(e) {
    var n = parseInt(elmn.value);
    if (!Number.isInteger(n)) {
      resultElm.value = "input must be integer!";
      return;
    }
    if ( n<=5 || n%2 == 0) {
      resultElm.value = "n must be greater than 5 and must be odd!";
      return;
    }
    Lemoine(n);
  });

  function IsPrime(n) {
    if (n<2) {
      return false;
    }

    var i;
    for (i = 2; i*i <= n; i++) {
        if (n%i == 0) {return false;}
    }
    return true;
  }

  function Lemoine(n) {
        var pqPairs = {};

        var q;
  	for (q = 1; q <= n/2; q++) {
  		var p = n - 2*q;

  		if (IsPrime(p) && IsPrime(q)) {
  			pqPairs[p] = q;
  		}
  	}

        resultElm.value = "";
  	for (var p in pqPairs) {
                var q = pqPairs[p];
  		resultElm.value += (n.toString() + " = " + p.toString() + " + ( 2 * " + q.toString() + " ) \n");
  	}
  }
  </script>

**Source Code (HTML)**

.. code-block:: html

  Input Odd Number Greater Than 5: <input type="text" name="n"><br><br>
  <button type="button" id="getLemoine">Get Lemoine’s Conjecture</button><br><br>
  <textarea id="result" rows="20" cols="50"></textarea>


**Source Code (JavaScript)**

.. code-block:: javascript

  var elmn = document.querySelector("input[name='n']");
  var btn = document.querySelector("#getLemoine");
  var resultElm = document.querySelector("#result");
  btn.addEventListener("click", function(e) {
    var n = parseInt(elmn.value);
    if (!Number.isInteger(n)) {
      resultElm.value = "input must be integer!";
      return;
    }
    if ( n<=5 || n%2 == 0) {
      resultElm.value = "n must be greater than 5 and must be odd!";
      return;
    }
    Lemoine(n);
  });

  function IsPrime(n) {
    if (n<2) {
      return false;
    }

    var i;
    for (i = 2; i*i <= n; i++) {
        if (n%i == 0) {return false;}
    }
    return true;
  }

  function Lemoine(n) {
        var pqPairs = {};

        var q;
  	for (q = 1; q <= n/2; q++) {
  		var p = n - 2*q;

  		if (IsPrime(p) && IsPrime(q)) {
  			pqPairs[p] = q;
  		}
  	}

        resultElm.value = "";
  	for (var p in pqPairs) {
                var q = pqPairs[p];
  		resultElm.value += (n.toString() + " = " + p.toString() + " + ( 2 * " + q.toString() + " ) \n");
  	}
  }

.. adsu:: 2

Tested on: ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``

----

References:

.. [1] `Lemoine’s Conjecture <{filename}/articles/2018/04/21/lemoine-conjecture%en.rst>`_

.. _Lemoine's conjecture: https://www.google.com/search?q=Lemoine's+Conjecture
