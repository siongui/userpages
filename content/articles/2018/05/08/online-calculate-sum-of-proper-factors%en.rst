Online Calculate Sum of Proper Divisors
#######################################

:date: 2018-05-08T23:34+08:00
:tags: Vue.js, JavaScript, Algorithm, Prime Number, Math
:category: Math
:summary: Online tool for prime factorization and calculating sum of proper
          divisors. The algorithm is implemented in JavaScript and UI in Vue.js.
:og_image: https://thinklikecomputerscientist.files.wordpress.com/2014/04/12.jpg
:adsu: yes


.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <div id="vueapp">
    positive integer: <input v-model="intn" placeholder="Input positive integer">
    <p>Prime Factors: {{ primes }}</p>
    <p>Sum of Proper Divisors: {{ sum }}</p>
  </div>

  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  // Get all prime factors of a given number n
  function PrimeFactors(n) {
    if (n==1) {return [1];}

    var pfs = [];
    // Get the number of 2s that divide n
    while (n%2 == 0) {
      pfs.push(2);
      n /= 2;
    }

    // n must be odd at this point. so we can skip one element
    // (note i = i + 2)
    for (var i=3; i*i <= n; i += 2) {
      // while i divides n, append i and divide n
      while (n%i == 0) {
        pfs.push(i);
        n /= i;
      }
    }

    // This condition is to handle the case when n is a prime number
    // greater than 2
    if ( n>2 ) {
      pfs.push(n);
    }

    return pfs;
  }

  function SumOfDivisors(pfs) {
    var m = {};
    for (var i = 0; i < pfs.length; i++) {
      var prime = pfs[i];
      if (m.hasOwnProperty(prime)) {
        m[prime] += 1;
      } else {
        m[prime] = 1;
      }
    }

    var sum = 1;
    for (var prime in m) {
      var exponents = m[prime];
      sum *= (Math.pow(prime, exponents+1)-1) / (prime-1);
    }
    return sum;
  }

  new Vue({
    el: '#vueapp',
    data: {
      intn: 284,
      primes: "",
      sum: ""
    },
    watch: {
      intn: {
        immediate: true,
        handler(val, oldVal) {
          var n = parseInt(val);
          if (isNaN(n) || n<1) {
            this.primes = "please input positive integer";
            return;
          }

          var pfs = PrimeFactors(n);
          this.primes = pfs;
          this.sum = (SumOfDivisors(pfs) - n);
        }
      }
    }
  });
  </script>

Online tool for `prime factorization`_ and calculating sum of proper divisors.
The following is source code of the tool.
For more details of the algorithm, see [1]_.

**HTML**:

.. code-block:: html

  <div id="vueapp">
    positive integer: <input v-model="intn" placeholder="Input positive integer">
    <p>Prime Factors: {{ primes }}</p>
    <p>Sum of Proper Divisors: {{ sum }}</p>
  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>


**JavaScript**:

.. code-block:: javascript

  'use strict';

  // Get all prime factors of a given number n
  function PrimeFactors(n) {
    if (n==1) {return [1];}

    var pfs = [];
    // Get the number of 2s that divide n
    while (n%2 == 0) {
      pfs.push(2);
      n /= 2;
    }

    // n must be odd at this point. so we can skip one element
    // (note i = i + 2)
    for (var i=3; i*i <= n; i += 2) {
      // while i divides n, append i and divide n
      while (n%i == 0) {
        pfs.push(i);
        n /= i;
      }
    }

    // This condition is to handle the case when n is a prime number
    // greater than 2
    if ( n>2 ) {
      pfs.push(n);
    }

    return pfs;
  }

  function SumOfDivisors(pfs) {
    var m = {};
    for (var i = 0; i < pfs.length; i++) {
      var prime = pfs[i];
      if (m.hasOwnProperty(prime)) {
        m[prime] += 1;
      } else {
        m[prime] = 1;
      }
    }

    var sum = 1;
    for (var prime in m) {
      var exponents = m[prime];
      sum *= (Math.pow(prime, exponents+1)-1) / (prime-1);
    }
    return sum;
  }

  new Vue({
    el: '#vueapp',
    data: {
      intn: 284,
      primes: "",
      sum: ""
    },
    watch: {
      intn: {
        immediate: true,
        handler(val, oldVal) {
          var n = parseInt(val);
          if (isNaN(n) || n<1) {
            this.primes = "please input positive integer";
            return;
          }

          var pfs = PrimeFactors(n);
          this.primes = pfs;
          this.sum = (SumOfDivisors(pfs) - n);
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

.. [1] `[Golang] Sum of the Proper Divisors (Factors) <{filename}/articles/2017/05/19/go-sum-of-proper-factors%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _prime factorization: https://www.google.com/search?q=prime+factorization
