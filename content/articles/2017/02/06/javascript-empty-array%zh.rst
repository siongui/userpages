[JavaScript] 清空陣列(array,數組)
#################################

:date: 2017-02-06T09:36+08:00
:tags: JavaScript, Web開發
:category: JavaScript
:summary: JavaScript_ 程式語言如何清空 `陣列`_ (array_,數組)
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


根據 [2]_ 的回答，最好方式為：

.. code-block:: javascript

  // 假設陣列(數組)名字為A
  while (A.length > 0) {
    A.pop();
  }

----

測試環境：

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

參考：

.. [1] `javascript empty array - Google search <https://www.google.com/search?q=javascript+empty+array>`_

       `javascript empty array - DuckDuckGo search <https://duckduckgo.com/?q=javascript+empty+array>`_

       `javascript empty array - Bing search <https://www.bing.com/search?q=javascript+empty+array>`_

       `javascript empty array - Yahoo search <https://search.yahoo.com/search?p=javascript+empty+array>`_

       `javascript empty array - Baidu search <https://www.baidu.com/s?wd=javascript+empty+array>`_

       `javascript empty array - Yandex search <https://www.yandex.com/search/?text=javascript+empty+array>`_
.. adsu:: 2
.. [2] `How do I empty an array in JavaScript? - Stack Overflow <http://stackoverflow.com/a/1232046>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _陣列: https://zh.wikipedia.org/wiki/%E6%95%B0%E7%BB%84
.. _array: https://en.wikipedia.org/wiki/Array_data_structure
