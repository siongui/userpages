[Python] Union of Two Arrays
############################

:date: 2018-03-20T23:25+08:00
:tags: Python, Algorithm, Set Operation
:category: Python
:summary: Find the set of all elements (union) in two arrays in Python.
:og_image: http://hideoushumpbackfreak.com/image.axd?picture=figure3-1_thumb.png
:adsu: yes


Find the set of all elements in two arrays, i.e., union of two arrays in
Python_.

The idea is to convert one array to the data structure of key-value pairs, i.e.,
hash table. The hash table in Python is built-in dictionary_ type. Then we check
if items of the other array is in the hash table. If not, append the item to the
first array, and return the first array after finish.

The following is the implementation of above idea.

.. rubric:: `Run Code on repl.it <https://repl.it/repls/BossySpiritedAnalysts>`__
   :class: align-center

.. code-block:: go

  def union(a, b):
  	map = {}

  	for x in a:
  		map[x] = True

  	for x in b:
  		if x not in map:
  			a.append(x)

  	return a


  if __name__ == "__main__":
  	a = [1, 2, 3, 4, 5]
  	b = [2, 3, 5, 7, 11]
  	print(union(a, b))

The result will be `[1, 2, 3, 4, 5, 7, 11]`.

The same idea can be used to find intersection of two arrays. See [1]_

.. adsu:: 2

Tested on:

- ``Ubuntu 17.10``, ``Python 3.6.3``
- `repl.it - Python3 Compiler`_

----

References:

.. [1] `[Python] Intersection of Two Arrays <{filename}/articles/2018/03/17/python-match-common-element-in-two-array%en.rst>`_

.. _Python: https://www.python.org/
.. _dictionary: https://www.tutorialspoint.com/python/python_dictionary.htm
.. _repl.it - Python3 Compiler: https://repl.it/languages/python3
