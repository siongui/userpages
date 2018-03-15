[JavaScript] Set Difference of Two Arrays
#########################################

:date: 2018-03-15T22:39+08:00
:tags: JavaScript, Algorithm, Set Operation
:category: JavaScript
:summary: Find set differecne of two arrays, i.e., the elements in one array but
          not in the other, in JavaScript.
:og_image: https://cdn.programiz.com/sites/tutorial2program/files/set-difference.jpg
:adsu: yes


Find the elements in one array but not in the other, i.e., set difference of two
arrays. In mathematical term:

  :math:`A-B=\{x\in A \ and \ x \notin B\}`

Try the following demo to get sense of the set difference:

.. raw:: html

  Array A: <input type="text" name="arra" value="[1, 2, 3, 4, 5]"><br>
  Array B: <input type="text" name="arrb" value="[2, 3, 5, 7, 11]"><br>
  <br>
  <button type="button" id="getDiff">Get A - B</button>
  <input type="text" name="diff"><br>

.. raw:: html

  <script>
  var arra = document.querySelector("input[name='arra']");
  var arrb = document.querySelector("input[name='arrb']");
  var diff = document.querySelector("input[name='diff']");
  var btn = document.querySelector("#getDiff");

  btn.addEventListener("click", function(e) {
    var arrayA = eval(arra.value);
    var arrayB = eval(arrb.value);
    var arrayDiff = [];

    var hashTable = {};
    for (var i = 0; i < arrayB.length; i++) {
      hashTable[arrayB[i]] = true;
    }

    for (var i = 0; i < arrayA.length; i++) {
      if (!hashTable.hasOwnProperty(arrayA[i])) {
        arrayDiff.push(arrayA[i]);
      }
    }

    diff.value = JSON.stringify(arrayDiff);
  });
  </script>

The idea is to convert the array B to the data structure of key-value pairs,
i.e., hash table. The hash table in JavaScript is built-in object_ type. Then we
check if items in array A is in the hash table. If not, append the item to the
difference array, and return the difference array after finish.

The following is the code for the demo:

**HTML**:

.. code-block:: html

  Array A: <input type="text" name="arra" value="[1, 2, 3, 4, 5]"><br>
  Array B: <input type="text" name="arrb" value="[2, 3, 5, 7, 11]"><br>
  <br>
  <button type="button" id="getDiff">Get A - B</button>
  <input type="text" name="diff"><br>

**JavaScript**:

.. code-block:: javascript

  var arra = document.querySelector("input[name='arra']");
  var arrb = document.querySelector("input[name='arrb']");
  var diff = document.querySelector("input[name='diff']");
  var btn = document.querySelector("#getDiff");

  btn.addEventListener("click", function(e) {
    var arrayA = eval(arra.value);
    var arrayB = eval(arrb.value);
    var arrayDiff = [];

    var hashTable = {};
    for (var i = 0; i < arrayB.length; i++) {
      hashTable[arrayB[i]] = true;
    }

    for (var i = 0; i < arrayA.length; i++) {
      if (!hashTable.hasOwnProperty(arrayA[i])) {
        arrayDiff.push(arrayA[i]);
      }
    }

    diff.value = JSON.stringify(arrayDiff);
  });

.. adsu:: 2

Tested on: `Chromium 64.0.3282.167 on Ubuntu 17.10 (64-bit)`

----

References:

.. [1] `[JavaScript] Intersection of Two Arrays <{filename}/articles/2018/03/12/javascript-match-common-element-in-two-array%en.rst>`_
.. [2] `[JavaScript] Union of Two Arrays <{filename}/articles/2018/03/13/javascript-set-of-all-elements-in-two-arrays%en.rst>`_

.. _object: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
