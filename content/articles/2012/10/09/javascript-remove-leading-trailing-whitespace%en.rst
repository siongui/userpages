[JavaScript] Remove Leading and Trailing Whitespaces
####################################################

:date: 2012-10-09 00:52
:modified: 2015-02-22 21:49
:tags: JavaScript, String Manipulation
:category: JavaScript
:summary: Remove leading and trailing whitespaces of a string in JavaScript.
:adsu: yes


This post is a brief summary of links in the references.

Remove (Trim) Leading Whitespace_ (similar to `Python lstrip`_)
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.. code-block:: javascript

  function lstrip(str) {
    return str.replace(/^\s+/g, "");
  }

Remove (Trim) Trailing Whitespace_ (similar to `Python rstrip`_)
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.. adsu:: 2

.. code-block:: javascript

  function rstrip(str) {
    return str.replace(/\s+$/g, "");
  }

Remove (Trim) Leading and Trailing Whitespace_ (similar to `Python strip`_)
+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

.. code-block:: javascript

  function strip(str) {
    return str.replace(/^\s+|\s+$/g, "");
  }

.. adsu:: 3

----

References:

.. [1] `Trimming a string in JavaScript <http://www.javascripter.net/faq/trim.htm>`_

.. [2] `trim - String strip() for JavaScript? - Stack Overflow <http://stackoverflow.com/questions/1418050/string-strip-for-javascript>`_

.. [3] `Faster JavaScript Trim <http://blog.stevenlevithan.com/archives/faster-trim-javascript>`_

.. [4] `String.prototype.trim() - JavaScript | MDN <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/Trim>`_

.. [5] `trim Method (String) (JavaScript) <https://msdn.microsoft.com/en-us/library/windows/apps/ff679971(v=vs.94).aspx>`_

.. [6] `JavaScript RegExp \s Metacharacter - W3Schools <http://www.w3schools.com/jsref/jsref_regexp_whitespace.asp>`_

.. _Whitespace: http://www.w3schools.com/jsref/jsref_regexp_whitespace.asp

.. _Python lstrip: https://docs.python.org/2/library/stdtypes.html#str.lstrip

.. _Python rstrip: https://docs.python.org/2/library/stdtypes.html#str.rstrip

.. _Python strip: https://docs.python.org/2/library/stdtypes.html#str.strip
