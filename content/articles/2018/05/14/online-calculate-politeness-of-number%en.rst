Online Calcualte Politeness of Number
#####################################

:date: 2018-05-14T22:32+08:00
:tags: Vue.js, JavaScript, Algorithm, Prime Number, Math
:category: Math
:summary: Calculate politeness of a number online, i.e., the number of ways it
          can be expressed as the sum of consecutive integers.
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/0/0f/Young_456_French.svg/640px-Young_456_French.svg.png
:adsu: yes


.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <div id="vueapp">
    positive integer: <input v-model="intn" placeholder="Input positive integer">
    <p>Politeness: {{ politeness }}</p>
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

  // Algorithm from wiki: An easy way of calculating the politeness of a positive
  // number is that of decomposing the number into its prime factors, taking the
  // powers of all prime factors greater than 2, adding 1 to all of them,
  // multiplying the numbers thus obtained with each other and subtracting 1.
  function CalculatePoliteness(n) {
    var pfs = PrimeFactors(n);

    // key: prime
    // value: prime exponent
    var pe = {};
    for (var i=0; i < pfs.length; i++) {
      var prime = pfs[i];
      if (pe.hasOwnProperty(prime)) {
        pe[prime] += 1;
      } else {
        pe[prime] = 1;
      }
    }

    var politeness = 1;
    for (var prime in pe) {
      if (prime > 2) {
        var exponent = pe[prime];
        politeness *= (exponent+1);
      }
    }
    politeness -= 1;

    return politeness;
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

          this.politeness = CalculatePoliteness(n);
        }
      }
    }
  });
  </script>


Calculate politeness of a number online.

The following description comes from Wikipedia:

  The politeness of a positive number is defined as the number of ways it can be
  expressed as the sum of consecutive integers. For instance, the politeness of
  9 is 2 because it has two odd divisors, 3 and itself, and two polite
  representations

    9 = 2 + 3 + 4 = 4 + 5;

The easy way to calculate the politeness also comes from Wikipedia:

  An easy way of calculating the politeness of a positive number is that of
  decomposing the number into its prime factors, taking the powers of all prime
  factors greater than 2, adding 1 to all of them, multiplying the numbers thus
  obtained with each other and subtracting 1.

The following is implementation of above algorithm in JavaScript/Vue.js_.

**HTML**:

.. code-block:: html

  <div id="vueapp">
    positive integer: <input v-model="intn" placeholder="Input positive integer">
    <p>Politeness: {{ politeness }}</p>
  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>


Given an input from user, we check if the input is an positive integer. Then we
get all prime factors of the positive integer and hence the politeness of the
number.


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

  // Algorithm from wiki: An easy way of calculating the politeness of a positive
  // number is that of decomposing the number into its prime factors, taking the
  // powers of all prime factors greater than 2, adding 1 to all of them,
  // multiplying the numbers thus obtained with each other and subtracting 1.
  function CalculatePoliteness(n) {
    var pfs = PrimeFactors(n);

    // key: prime
    // value: prime exponent
    var pe = {};
    for (var i=0; i < pfs.length; i++) {
      var prime = pfs[i];
      if (pe.hasOwnProperty(prime)) {
        pe[prime] += 1;
      } else {
        pe[prime] = 1;
      }
    }

    var politeness = 1;
    for (var prime in pe) {
      if (prime > 2) {
        var exponent = pe[prime];
        politeness *= (exponent+1);
      }
    }
    politeness -= 1;

    return politeness;
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

          this.politeness = CalculatePoliteness(n);
        }
      }
    }
  });

.. adsu:: 2

----

Tested on:

- ``Chromium 66.0.3359.139 on Ubuntu 18.04 (64-bit)``
- ``Vue.js 2.5.16``

----

References:

.. [1] `Calcualte Politeness of Number in Golang <{filename}/articles/2018/05/12/calculate-politeness-of-number-in-go%en.rst>`_

.. _Vue.js: https://vuejs.org/
