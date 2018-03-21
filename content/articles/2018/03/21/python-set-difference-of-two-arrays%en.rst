[Python] Set Difference of Two Arrays
#####################################

:date: 2018-03-21T23:07+08:00
:tags: Python, Algorithm, Set Operation
:category: Python
:summary: Find set differecne of two arrays, i.e., the elements in one array but
          not in the other, in Python.
:og_image: https://cdn.programiz.com/sites/tutorial2program/files/set-difference.jpg
:adsu: yes


Find the elements in one array but not in the other, i.e., set difference of two
arrays. In mathematical term:

  :math:`A-B=\{x\in A \ and \ x \notin B\}`

The idea is to convert the array B to the data structure of key-value pairs,
i.e., hash table. The hash table in Python_ is built-in dictionary_ type. Then
we check if items in array A is in the hash table. If not, append the item to
the difference array, and return the difference array after finish.

The following is the implementation of above idea.

.. rubric:: `Run Code on repl.it <https://repl.it/repls/TotalSlategreyRecovery>`__
   :class: align-center

.. code-block:: go

  # A - B
  def difference(a, b):
  	map = {}

  	for x in b:
  		map[x] = True

  	diff = []
  	for x in a:
  		if x not in map:
  			diff.append(x)

  	return diff

  if __name__ == "__main__":
  	a = [1, 2, 3, 4, 5]
  	b = [2, 3, 5, 7, 11]
  	print(difference(a, b))

The result will be `[1, 4]`.

.. adsu:: 2

Tested on:

- ``Ubuntu 17.10``, ``Python 3.6.3``
- `repl.it - Python3 Compiler`_

----

References:

.. [1] `[Python] Intersection of Two Arrays <{filename}/articles/2018/03/17/python-match-common-element-in-two-array%en.rst>`_
.. [2] `[Python] Union of Two Arrays <{filename}/articles/2018/03/20/python-set-of-all-elements-in-two-arrays%en.rst>`_

.. _Python: https://www.python.org/
.. _dictionary: https://www.tutorialspoint.com/python/python_dictionary.htm
.. _repl.it - Python3 Compiler: https://repl.it/languages/python3
