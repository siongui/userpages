[JavaScript] Generate Random String From [a-z0-9]
#################################################

:date: 2017-01-14T07:36+08:00
:tags: JavaScript, String Manipulation, Random Number
:category: JavaScript
:summary: Generate a random string from [a-z0-9] in JavaScript_.


.. code-block:: javascript

  /**
   * Returns a random integer between min (inclusive) and max (inclusive)
   * Using Math.round() will give you a non-uniform distribution!
   * Code from http://stackoverflow.com/a/1527820
   */
  function getRandomInt(min, max) {
      return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  /**
   * Input: the length of random string
   * Output: a random string from [a-z0-9]
   */
  function RandomString(strlen) {
      const chars = "abcdefghijklmnopqrstuvwxyz0123456789";
      var result = "";
      for (var i=0; i<strlen; i++) {
          result += chars[getRandomInt(0,35)];
      }
      return result;
  }

  // function call to generate a random string of length 10
  RandomString(10);

Run `code <https://repl.it/FJIx/0>`_ online:

.. raw:: html

  <script src="//repl.it/embed/FJIx/0.js"></script>

----

Tested on: ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``.

----

References:

.. [1] `javascript random integer - Google search <https://www.google.com/search?q=javascript+random+integer>`_

       `javascript random integer - DuckDuckGo search <https://duckduckgo.com/?q=javascript+random+integer>`_

       `javascript random integer - Bing search <https://www.bing.com/search?q=javascript+random+integer>`_

       `javascript random integer - Yahoo search <https://search.yahoo.com/search?p=javascript+random+integer>`_

       `javascript random integer - Baidu search <https://www.baidu.com/s?wd=javascript+random+integer>`_

       `javascript random integer - Yandex search <https://www.yandex.com/search/?text=javascript+random+integer>`_

.. [2] `Generating random whole numbers in JavaScript in a specific range? - Stack Overflow <http://stackoverflow.com/a/1527820>`_

.. [3] `javascript constants - Google search <https://www.google.com/search?q=javascript+constants>`_

       `javascript constants - DuckDuckGo search <https://duckduckgo.com/?q=javascript+constants>`_

       `javascript constants - Bing search <https://www.bing.com/search?q=javascript+constants>`_

       `javascript constants - Yahoo search <https://search.yahoo.com/search?p=javascript+constants>`_

       `javascript constants - Baidu search <https://www.baidu.com/s?wd=javascript+constants>`_

       `javascript constants - Yandex search <https://www.yandex.com/search/?text=javascript+constants>`_

.. [4] `How to declare string constants in JavaScript? - Stack Overflow <http://stackoverflow.com/questions/5786054/how-to-declare-string-constants-in-javascript>`_

.. [5] `[Golang] Generate Random String From [a-z0-9] <{filename}../../../2015/04/13/go-generate-random-string%en.rst>`_


.. _JavaScript: https://www.google.com/search?q=JavaScript
