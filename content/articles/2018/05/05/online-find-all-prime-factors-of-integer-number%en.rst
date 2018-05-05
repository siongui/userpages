Online Prime Factorization
##########################

:date: 2018-05-05T23:18+08:00
:tags: Vue.js, JavaScript, Algorithm, Prime Number, Math
:category: Math
:summary: Online tool that helps you do prime factorization. The algorithm is
          implemented in JavaScript and UI in Vue.js.
:og_image: https://www.mathsisfun.com/numbers/images/factor-tree-48.gif
:adsu: yes


.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <div id="vueapp">
    positive integer: <input v-model="intn" placeholder="Input positive integer">
    <p>{{ primes }}</p>
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

  new Vue({
    el: '#vueapp',
    data: {
      intn: 6,
      primes: ""
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

          this.primes = PrimeFactors(n);
        }
      }
    }
  });
  </script>

Online `prime factorization`_ tool. The algorithm comes from GeeksforGeeks [1]_
and implemented in JavaScript. The UI is implemented using `Vue.js`_. The
following is complete source code.

**HTML**:

.. code-block:: html

  <div id="vueapp">
    positive integer: <input v-model="intn" placeholder="Input positive integer">
    <p>{{ primes }}</p>
  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>


Given an input from user, we check if the input is an integer. Then we run the
sieve algorithm to check if the number is prime.


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

  new Vue({
    el: '#vueapp',
    data: {
      intn: 6,
      primes: ""
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

          this.primes = PrimeFactors(n);
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

.. [1] `Efficient program to print all prime factors of a given number - GeeksforGeeks <https://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/>`_
.. [2] `[Golang] Get All Prime Factors of Integer Number <{filename}/articles/2017/05/09/go-find-all-prime-factors-of-integer-number%en.rst>`_
.. [3] `[Vue.js] Online Sieve of Eratosthenes Demo <{filename}/articles/2018/05/02/vuejs-online-sieve-of-eratosthenes-demo%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _prime factorization: https://www.google.com/search?q=prime+factorization
