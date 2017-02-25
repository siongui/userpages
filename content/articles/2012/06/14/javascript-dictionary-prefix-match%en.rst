JavaScript Prefix Match (Dictionary Application)
################################################

:date: 2012-06-14 22:37
:modified: 2015-02-24 12:15
:tags: JavaScript, String Manipulation
:category: JavaScript
:summary: Given a set of strings (words), and an user input string, output a set
          of strings with prefix the same as the user input string
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


Recently I have a dictionary web application which needs to find prefix-matched
words in an array of strings. Here is the specification of my problem:

  Input: *Given a set of strings (words), and an user input string*

  Output: *a set of strings with prefix the same as the user input string*

My naive solution is as follows:
Because the set of given strings are huge, we first group them by the first
letter of the string. For example, string with prefix *a* goes to first group,
and strings with prefix *b* goes to second group, etc. Then we use the user
input sting to match the prefix of strings in the group by
`JavaScript String indexOf() Method`_. We don't need to search the entire given
strings. We just need to search the corresponding group which has the same
prefix as the user input string. This can save a lot of time, at the cost of
some additional space and code. If more efficient solution is needed, try
`Radix tree`_ to make your search faster [4]_.

The following is the demo and sample code of the naive solution I just mention.
Type some word that starts with *a*, *b*, or *c* in the demo, and you will see
the result.

.. rubric:: `Demo <{filename}/code/javascript-dictionary-prefix-match/naive.html>`_
      :class: align-center

.. adsu:: 2

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/javascript-dictionary-prefix-match/naive.html
.. adsu:: 3

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/code/javascript-dictionary-prefix-match/naive.js

----

References:

.. [1] Google Search Keyword: `javascript string prefix match <https://www.google.com/search?q=javascript+string+prefix+match>`_

.. [2] `JavaScript startsWith and endsWith Implementation for Strings <http://rickyrosario.com/blog/javascript-startswith-and-endswith-implementation-for-strings/>`_

.. [3] `JavaScript String match() Method <http://www.w3schools.com/jsref/jsref_match.asp>`_

.. [4] `The Most Efficient Algorithm to Find First Prefix-Match From a Sorted String Array? - Stack Overflow <http://stackoverflow.com/questions/457160/the-most-efficient-algorithm-to-find-first-prefix-match-from-a-sorted-string-arr>`_


.. _JavaScript String indexOf() Method: http://www.w3schools.com/jsref/jsref_indexof.asp

.. _Radix tree: http://en.wikipedia.org/wiki/Radix_tree
