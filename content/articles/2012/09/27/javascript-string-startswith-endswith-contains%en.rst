[JavaScript] String startswith, endswith and contains Implementation
####################################################################

:date: 2012-09-27 19:59
:modified: 2015-02-20 23:11
:tags: html, String Manipulation, JavaScript, DOM
:category: JavaScript
:summary: JavaScript equivalent of Python string startswith, endswith, and contains.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


When we do programming with strings, it's very common to have the following
questions:

  1. how do I check whether the *string* starts with *prefix*?
  2. how do I check whether the *string* ends with *suffix*?
  3. how do I check whether the *string* contains *sub-string*?

We will show how to do these case by case in JavaScript:

String startswith
+++++++++++++++++

We can use JavaScript built-in `indexOf()`_ function to check whether a string
starts (or begins) with prefix. For alternative solution, please refer to
reference [1]_.

.. code-block:: javascript

  /**
   * The string starts with prefix?
   * @param {string} string The string starts with prefix?
   * @param {string} prefix The string starts with prefix?
   * @return {boolean} true if string starts with prefix, otherwise false
   */
  startswith = function(string, prefix) {
    return string.indexOf(prefix) == 0;
  };

.. adsu:: 2

String endswith
+++++++++++++++

Again we use JavaScript built-in `indexOf()`_ function to check whether a string
ends with suffix. For alternative solution, please refer to reference [2]_.

.. code-block:: javascript

  /**
   * The string ends with suffix?
   * @param {string} string The string ends with suffix?
   * @param {string} suffix The string ends with suffix?
   * @return {boolean} true if string ends with suffix, otherwise false
   */
  endswith = function(string, suffix) {
    return string.indexOf(suffix, string.length - suffix.length) != -1;
  };

.. adsu:: 3

String contains
+++++++++++++++

The same trick of JavaScript built-in `indexOf()`_ function to check whether a
string contains another string. For alternative solution, please refer to
reference [3]_.

.. code-block:: javascript

  /**
   * The string contains substr?
   * @param {string} string The string contains substr?
   * @param {string} substr The string contains substr?
   * @return {boolean} true if string contains substr, otherwise false
   */
  contains = function(string, substr) {
    return string.indexOf(substr) != -1;
  };

----

References:

.. [1] `How to check if a string “StartsWith” another string? <http://stackoverflow.com/questions/646628/how-to-check-if-a-string-startswith-another-string>`_

.. [2] `endsWith in javascript <http://stackoverflow.com/questions/280634/endswith-in-javascript>`_

.. [3] `How can I check if one string contains another substring? <http://stackoverflow.com/questions/1789945/how-can-i-check-if-one-string-contains-another-substring>`_

.. [4] `JavaScript String indexOf() Method <http://www.w3schools.com/jsref/jsref_indexof.asp>`_

.. [5] `JavaScript basename() <{filename}../../10/02/javascript-basename%en.rst>`_

.. _indexOf(): http://www.w3schools.com/jsref/jsref_indexof.asp
