[Python] Intersection of Two Arrays
###################################

:date: 2018-03-17T22:30+08:00
:tags: Python, Algorithm, Set Operation
:category: Python
:summary: Find common elements (matches, intersection) of two arrays in Python.
:og_image: http://4.bp.blogspot.com/-034xqHwdXwU/UqWAJEoDWKI/AAAAAAAAAHk/HyCOJbka4CU/s400/Intersection+of+two+arrays+java+coding+.jpg
:adsu: yes


Find common elements in two arrays, i.e., intersection of two arrays in Python_.

The idea is to convert one array to the data structure of key-value pairs, i.e.,
hash table. The hash table in Python is built-in dictionary_ type. Then we check
if items of the other array is in the hash table. If yes, the item is in the
intersection of the two arrays.

The following is the implementation of above idea.

.. rubric:: `Run Code on repl.it <https://repl.it/repls/LegalScholarlyPerl>`__
   :class: align-center

.. code-block:: go

  def intersection(a, b):
  	map = {}

  	for x in a:
  		map[x] = True

  	intersection = []
  	for x in b:
  		if x in map:
  			intersection.append(x)

  	return intersection


  if __name__ == "__main__":
  	a = [1, 2, 3, 4, 5]
  	b = [2, 3, 5, 7, 11]
  	print(intersection(a, b))

The result will be `[2 3 5]`.

The same idea can be used to find union of two arrays. See [1]_.

.. adsu:: 2

Tested on:

- ``Ubuntu 17.10``, ``Python 3.6.3``
- `repl.it - Python3 Compiler`_

----

References:

.. [1] `[Python] Union of Two Arrays <{filename}/articles/2018/03/20/python-set-of-all-elements-in-two-arrays%en.rst>`_

.. _Python: https://www.python.org/
.. _dictionary: https://www.tutorialspoint.com/python/python_dictionary.htm
.. _repl.it - Python3 Compiler: https://repl.it/languages/python3
