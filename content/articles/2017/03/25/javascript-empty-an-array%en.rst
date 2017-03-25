JavaScript Empty an Array
#########################

:date: 2017-03-25T22:32+08:00
:tags: html, JavaScript, DOM
:category: JavaScript
:summary: My favorite way to empty an array in JavaScript and why.
:og_image: http://www.htmlgoodies.com/img/2007/04/array-comparison.jpg
:adsu: yes


My favorite answer is:

.. code-block:: javascript

  while (A.length > 0) {
    A.pop();
  }

I strongly suggest to read the discussions of the Stack Overflow question [2]_.
The reason why I like the answer above is that long time ago, I had to
`remove all child nodes of a DOM element`_ [3]_, the bug-free way to do this is:

.. code-block:: javascript

  while (elm.hasChildNodes()) {
    elm.removeChild(elm.lastChild);
  }

Compare the idea of *remove all child nodes* and *empty an array* in above code,
it is very clearly that the idea is the same. This is an interesting finding so
I make it a post.

.. adsu:: 2

----

References:

.. [1] | `javascript empty array - Google search <https://www.google.com/search?q=javascript+empty+array>`_
       | `javascript empty array - DuckDuckGo search <https://duckduckgo.com/?q=javascript+empty+array>`_
       | `javascript empty array - Ecosia search <https://www.ecosia.org/search?q=javascript+empty+array>`_
       | `javascript empty array - Qwant search <https://www.qwant.com/?q=javascript+empty+array>`_
       | `javascript empty array - Bing search <https://www.bing.com/search?q=javascript+empty+array>`_
       | `javascript empty array - Yahoo search <https://search.yahoo.com/search?p=javascript+empty+array>`_
       | `javascript empty array - Baidu search <https://www.baidu.com/s?wd=javascript+empty+array>`_
       | `javascript empty array - Yandex search <https://www.yandex.com/search/?text=javascript+empty+array>`_

.. [2] `How do I empty an array in JavaScript? - Stack Overflow <http://stackoverflow.com/questions/1232040/how-do-i-empty-an-array-in-javascript>`_

.. [3] `JavaScript Remove All Children of a DOM Element <{filename}../../../2012/09/26/javascript-remove-all-children-of-dom-element%en.rst>`_

.. _remove all child nodes of a DOM element: {filename}../../../2012/09/26/javascript-remove-all-children-of-dom-element%en.rst
