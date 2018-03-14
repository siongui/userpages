[JavaScript] Intersection of Two Arrays
#######################################

:date: 2018-03-12T23:21+08:00
:tags: JavaScript, Algorithm, Set Operation
:category: JavaScript
:summary: Find common elements (matches, intersection) of two arrays
          in JavaScript.
:og_image: http://4.bp.blogspot.com/-034xqHwdXwU/UqWAJEoDWKI/AAAAAAAAAHk/HyCOJbka4CU/s400/Intersection+of+two+arrays+java+coding+.jpg
:adsu: yes


See demo first:

.. raw:: html

  Array A: <input type="text" name="arra" value="[1, 2, 3, 4, 5]"><br>
  Array B: <input type="text" name="arrb" value="[2, 3, 5, 7, 11]"><br>
  <br>
  <button type="button" id="getIntersection">Get Intersection</button>
  <input type="text" name="intersection"><br>

.. raw:: html

  <script>
  var arra = document.querySelector("input[name='arra']");
  var arrb = document.querySelector("input[name='arrb']");
  var intersection = document.querySelector("input[name='intersection']");
  var btn = document.querySelector("#getIntersection");

  btn.addEventListener("click", function(e) {
    var arrayA = eval(arra.value);
    var arrayB = eval(arrb.value);
    var arrayU = [];

    var hashTable = {};
    for (var i = 0; i < arrayA.length; i++) {
      hashTable[arrayA[i]] = true;
    }

    for (var i = 0; i < arrayB.length; i++) {
      if (hashTable.hasOwnProperty(arrayB[i])) {
        arrayU.push(arrayB[i]);
      }
    }

    intersection.value = JSON.stringify(arrayU);
  });
  </script>

Find common elements in two arrays, i.e., intersection or matches in two arrays.
There are many ways to find the intersection [1]_ [3]_. Here we will implement
the method mentioned in [2]_.

The idea is to convert one array to the data structure of key-value pairs, i.e.,
hash table. The hash table in JavaScript is built-in object_ type. Then we check
if items of the other array is in the hash table. If yes, the item is in the
intersection of the two arrays.

The following is the code for the demo:

**HTML**:

.. code-block:: html

  Array A: <input type="text" name="arra" value="[1, 2, 3, 4, 5]"><br>
  Array B: <input type="text" name="arrb" value="[2, 3, 5, 7, 11]"><br>
  <br>
  <button type="button" id="getIntersection">Get Intersection</button>
  <input type="text" name="intersection"><br>

**JavaScript**:

.. code-block:: javascript

  var arra = document.querySelector("input[name='arra']");
  var arrb = document.querySelector("input[name='arrb']");
  var intersection = document.querySelector("input[name='intersection']");
  var btn = document.querySelector("#getIntersection");

  btn.addEventListener("click", function(e) {
    var arrayA = eval(arra.value);
    var arrayB = eval(arrb.value);
    var arrayU = [];

    var hashTable = {};
    for (var i = 0; i < arrayA.length; i++) {
      hashTable[arrayA[i]] = true;
    }

    for (var i = 0; i < arrayB.length; i++) {
      if (hashTable.hasOwnProperty(arrayB[i])) {
        arrayU.push(arrayB[i]);
      }
    }

    intersection.value = JSON.stringify(arrayU);
  });

.. adsu:: 2

For the same implementation in Go_, see [5]_.

Tested on: `Chromium 64.0.3282.167 on Ubuntu 17.10 (64-bit)`

----

References:

.. [1] | `match common element in two array - Google search <https://www.google.com/search?q=match+common+element+in+two+array>`_
       | `match common element in two array - DuckDuckGo search <https://duckduckgo.com/?q=match+common+element+in+two+array>`_
       | `match common element in two array - Ecosia search <https://www.ecosia.org/search?q=match+common+element+in+two+array>`_
       | `match common element in two array - Qwant search <https://www.qwant.com/?q=match+common+element+in+two+array>`_
       | `match common element in two array - Bing search <https://www.bing.com/search?q=match+common+element+in+two+array>`_
       | `match common element in two array - Yahoo search <https://search.yahoo.com/search?p=match+common+element+in+two+array>`_
       | `match common element in two array - Baidu search <https://www.baidu.com/s?wd=match+common+element+in+two+array>`_
       | `match common element in two array - Yandex search <https://www.yandex.com/search/?text=match+common+element+in+two+array>`_

.. [2] `efficiency - Quick algorithm to find matches between two arrays - Software Engineering Stack Exchange <https://softwareengineering.stackexchange.com/a/223477>`_

.. [3] | `intersection of two arrays - Google search <https://www.google.com/search?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - DuckDuckGo search <https://duckduckgo.com/?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - Ecosia search <https://www.ecosia.org/search?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - Qwant search <https://www.qwant.com/?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - Bing search <https://www.bing.com/search?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - Yahoo search <https://search.yahoo.com/search?p=intersection+of+two+arrays>`_
       | `intersection of two arrays - Baidu search <https://www.baidu.com/s?wd=intersection+of+two+arrays>`_
       | `intersection of two arrays - Yandex search <https://www.yandex.com/search/?text=intersection+of+two+arrays>`_

.. [4] | `Find Union and Intersection of two unsorted arrays - GeeksforGeeks <https://www.geeksforgeeks.org/find-union-and-intersection-of-two-unsorted-arrays/>`_
       | `Union and Intersection of two sorted arrays - GeeksforGeeks <https://www.geeksforgeeks.org/union-and-intersection-of-two-sorted-arrays-2/>`_

.. [5] `[Golang] Intersection of Two Arrays <{filename}/articles/2018/03/09/go-match-common-element-in-two-array%en.rst>`_

.. _object: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
.. _Go: https://golang.org/
