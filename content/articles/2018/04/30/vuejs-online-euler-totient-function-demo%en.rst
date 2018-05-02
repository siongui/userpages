[Vue.js] Online Euler's Totient Function Demo
#############################################

:date: 2018-04-30T23:46+08:00
:tags: Vue.js, JavaScript, Algorithm, Prime Number, Math
:category: Vue.js
:summary: Euler's Totient Function φ(n) counts the positive integers that are
          relatively prime to n. This online demo use naive method to calculate
          φ(n) and positive integers coprime to n.
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/9/9b/EulerPhi.svg/1200px-EulerPhi.svg.png
:adsu: yes


`Euler's Totient Function`_ φ(n) counts the positive integers that are relatively
prime to n, i.e., the number of positive integers whose GCD (Greatest Common
Divisor) with n is 1.

This online demo use naive method to calculate φ(n) and positive integers
coprime to n. The UI interaction is written using Vue.js_, and the algorithm is
implemented using JavaScript.

.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <div id="vueapp">
    input positive integer: <input v-model="intn" placeholder="Input positive integer">
    <p>φ(n): {{ phi }}</p>
    <p>Numbers coprime to n: {{ coprimes }}</p>
    <p>{{ message }}</p>
  </div>

  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  // greatest common divisor (GCD) via Euclidean algorithm
  function GCD(a, b) {
      while (b != 0) {
              var t = b;
              b = a % b;
              a = t;
      }
      return a;
  }

  // Euler’s Totient Function
  function Phi(n) {
      var result = {phi:1, coprimes: [1]};
      for (var i = 2; i < n; i++) {
              if (GCD(i, n) == 1) {
                      result.phi += 1;
                      result.coprimes.push(i);
              }
      }
      return result;
  }

  new Vue({
    el: '#vueapp',
    data: {
      intn: 1,
      phi: "",
      coprimes: "",
      message: ""
    },
    watch: {
      intn: {
        immediate: true,
        handler(val, oldVal) {
          var n = parseInt(val);
          if (isNaN(n)) {
            this.message = "please input integer";
            return;
          }
          if (n < 1) {
            this.message = "input integer must >= 1";
            return;
          }
          this.message = "";

          var result = Phi(n);
          this.phi = result.phi;
          this.coprimes = result.coprimes;
        }
      }
    }
  });
  </script>

**HTML**:

.. code-block:: html

  <div id="vueapp">
    input positive integer: <input v-model="intn" placeholder="Input positive integer">
    <p>φ(n): {{ phi }}</p>
    <p>Numbers coprime to n: {{ coprimes }}</p>
    <p>{{ message }}</p>
  </div>

  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.16/dist/vue.js"></script>

Given an input from user, we check if the input is a positive integer. Then we
run GCD algorithm to check if the number is coprime to n. If yes, add the number
to our result.


**JavaScript**:

.. code-block:: javascript

  'use strict';

  // greatest common divisor (GCD) via Euclidean algorithm
  function GCD(a, b) {
      while (b != 0) {
              var t = b;
              b = a % b;
              a = t;
      }
      return a;
  }

  // Euler’s Totient Function
  function Phi(n) {
      var result = {phi:1, coprimes: [1]};
      for (var i = 2; i < n; i++) {
              if (GCD(i, n) == 1) {
                      result.phi += 1;
                      result.coprimes.push(i);
              }
      }
      return result;
  }

  new Vue({
    el: '#vueapp',
    data: {
      intn: 1,
      phi: "",
      coprimes: "",
      message: ""
    },
    watch: {
      intn: {
        immediate: true,
        handler(val, oldVal) {
          var n = parseInt(val);
          if (isNaN(n)) {
            this.message = "please input integer";
            return;
          }
          if (n < 1) {
            this.message = "input integer must >= 1";
            return;
          }
          this.message = "";

          var result = Phi(n);
          this.phi = result.phi;
          this.coprimes = result.coprimes;
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

.. [1] `[Golang] Euler's Totient Function <{filename}/articles/2017/06/04/go-euler-totient-function%en.rst>`_
.. [2] `[Vue.js] Online Goldbach's Conjecture Demo <{filename}/articles/2018/04/29/vuejs-online-goldbach-conjecture-demo%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _Euler's Totient Function: https://www.google.com/search?q=Euler's+totient+function
