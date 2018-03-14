[JavaScript] Union of Two Arrays
################################

:date: 2018-03-13T23:19+08:00
:tags: JavaScript, Algorithm, Set Operation
:category: JavaScript
:summary: Find the set of all elements (union) in two arrays in JavaScript.
:og_image: http://hideoushumpbackfreak.com/image.axd?picture=figure3-1_thumb.png
:adsu: yes


See demo first:

.. raw:: html

  Array A: <input type="text" name="arra" value="[1, 2, 3, 4, 5]"><br>
  Array B: <input type="text" name="arrb" value="[2, 3, 5, 7, 11]"><br>
  <br>
  <button type="button" id="getUnion">Get Union</button>
  <input type="text" name="union"><br>

.. raw:: html

  <script>
  var arra = document.querySelector("input[name='arra']");
  var arrb = document.querySelector("input[name='arrb']");
  var union = document.querySelector("input[name='union']");
  var btn = document.querySelector("#getUnion");

  btn.addEventListener("click", function(e) {
    var arrayA = eval(arra.value);
    var arrayB = eval(arrb.value);

    var hashTable = {};
    for (var i = 0; i < arrayA.length; i++) {
      hashTable[arrayA[i]] = true;
    }

    for (var i = 0; i < arrayB.length; i++) {
      if (!hashTable.hasOwnProperty(arrayB[i])) {
        arrayA.push(arrayB[i]);
      }
    }

    union.value = JSON.stringify(arrayA);
  });
  </script>

Find the set of all elements in two arrays, i.e., union of two arrays. This is
the continued work of my post yesterday [1]_, which is intersection of two
arrays.

The idea is to convert one array to the data structure of key-value pairs, i.e.,
hash table. The hash table in JavaScript is built-in object_ type. Then we check
if items of the other array is in the hash table. If not, append the item to the
first array, and return the first array after finish.

The following is the code for the demo:

**HTML**:

.. code-block:: html

  Array A: <input type="text" name="arra" value="[1, 2, 3, 4, 5]"><br>
  Array B: <input type="text" name="arrb" value="[2, 3, 5, 7, 11]"><br>
  <br>
  <button type="button" id="getUnion">Get Union</button>
  <input type="text" name="union"><br>

**JavaScript**:

.. code-block:: javascript

  var arra = document.querySelector("input[name='arra']");
  var arrb = document.querySelector("input[name='arrb']");
  var union = document.querySelector("input[name='union']");
  var btn = document.querySelector("#getUnion");

  btn.addEventListener("click", function(e) {
    var arrayA = eval(arra.value);
    var arrayB = eval(arrb.value);

    var hashTable = {};
    for (var i = 0; i < arrayA.length; i++) {
      hashTable[arrayA[i]] = true;
    }

    for (var i = 0; i < arrayB.length; i++) {
      if (!hashTable.hasOwnProperty(arrayB[i])) {
        arrayA.push(arrayB[i]);
      }
    }

    union.value = JSON.stringify(arrayA);
  });

.. adsu:: 2

Tested on: `Chromium 64.0.3282.167 on Ubuntu 17.10 (64-bit)`

----

References:

.. [1] `[JavaScript] Intersection of Two Arrays <{filename}/articles/2018/03/12/javascript-match-common-element-in-two-array%en.rst>`_

.. _object: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
