[JavaScript] Check if Value of Input Text Field is Integer
##########################################################

:date: 2018-04-23T23:38+08:00
:tags: JavaScript, Type Casting, Type Conversion
:category: JavaScript
:summary: check HTML text input field to see if the value is integer.
:og_image: https://i.ytimg.com/vi/oDUjP4N_MtQ/maxresdefault.jpg
:adsu: yes


Yesterday I was writing an online demo for Lemoine’s conjecture [3]_, I had to
check if the value of input text field is integer. If not, the value must be
discarded and shows a message to tell users that input must be integer.

After some searches, I found that `parseInt()`_ and `isNaN()`_ can be used to
do the job. The return value of text input field is *string*. First try
`parseInt()`_ to convert the value to integer. If the type casting fails, it
`parseInt()`_ will return *NaN*, so we can use `isNaN()`_ to check the result of
conversion.

The following is demo of above idea:

.. raw:: html

  <input type="text" name="n">
  <button type="button" id="check">Check Integer</button><br><br>
  <div id="result"></div>
  <script>
  document.querySelector("#check").addEventListener("click", function(e) {
    var n = parseInt(document.querySelector("input[name='n']").value);
    if (isNaN(n)) {
      document.querySelector("#result").innerHTML = "input must be integer";
      return;
    }
    document.querySelector("#result").innerHTML = "integer is " + n;
  });
  </script>

**Source Code (HTML)**

.. code-block:: html

  <input type="text" name="n">
  <button type="button" id="check">Check Integer</button><br><br>
  <div id="result"></div>

**Source Code (JavaScript)**

.. code-block:: javascript

  document.querySelector("#check").addEventListener("click", function(e) {
    var n = parseInt(document.querySelector("input[name='n']").value);
    if (isNaN(n)) {
      document.querySelector("#result").innerHTML = "input must be integer";
      return;
    }
    document.querySelector("#result").innerHTML = "integer is " + n;
  });

.. adsu:: 2

Tested on: ``Chromium 65.0.3325.181 on Ubuntu 17.10 (64-bit)``

----

References:

.. [1] | `js check integer - Google search <https://www.google.com/search?q=js+check+integer>`_
       | `js check integer - DuckDuckGo search <https://duckduckgo.com/?q=js+check+integer>`_
       | `js check integer - Ecosia search <https://www.ecosia.org/search?q=js+check+integer>`_
       | `js check integer - Qwant search <https://www.qwant.com/?q=js+check+integer>`_
       | `js check integer - Bing search <https://www.bing.com/search?q=js+check+integer>`_
       | `js check integer - Yahoo search <https://search.yahoo.com/search?p=js+check+integer>`_
       | `js check integer - Baidu search <https://www.baidu.com/s?wd=js+check+integer>`_
       | `js check integer - Yandex search <https://www.yandex.com/search/?text=js+check+integer>`_
.. [2] | `js read input - Google search <https://www.google.com/search?q=js+read+input>`_
       | `js read input - DuckDuckGo search <https://duckduckgo.com/?q=js+read+input>`_
       | `js read input - Ecosia search <https://www.ecosia.org/search?q=js+read+input>`_
       | `js read input - Qwant search <https://www.qwant.com/?q=js+read+input>`_
       | `js read input - Bing search <https://www.bing.com/search?q=js+read+input>`_
       | `js read input - Yahoo search <https://search.yahoo.com/search?p=js+read+input>`_
       | `js read input - Baidu search <https://www.baidu.com/s?wd=js+read+input>`_
       | `js read input - Yandex search <https://www.yandex.com/search/?text=js+read+input>`_
.. [3] `Online Lemoine’s Conjecture Demo <{filename}/articles/2018/04/22/online-lemoine-conjecture-demo%en.rst>`_
.. [4] | `javascript input text integer - Google search <https://www.google.com/search?q=javascript+input+text+integer>`_
       | `javascript input text integer - DuckDuckGo search <https://duckduckgo.com/?q=javascript+input+text+integer>`_
       | `javascript input text integer - Ecosia search <https://www.ecosia.org/search?q=javascript+input+text+integer>`_
       | `javascript input text integer - Qwant search <https://www.qwant.com/?q=javascript+input+text+integer>`_
       | `javascript input text integer - Bing search <https://www.bing.com/search?q=javascript+input+text+integer>`_
       | `javascript input text integer - Yahoo search <https://search.yahoo.com/search?p=javascript+input+text+integer>`_
       | `javascript input text integer - Baidu search <https://www.baidu.com/s?wd=javascript+input+text+integer>`_
       | `javascript input text integer - Yandex search <https://www.yandex.com/search/?text=javascript+input+text+integer>`_

.. _parseInt(): https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/parseInt
.. _isNaN(): https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/isNaN
