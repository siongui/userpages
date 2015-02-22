[C] Remove String Trailing Newline (Carriage Return)
####################################################

:tags: C, remove trailing newline, remove carriage return, String Manipulation, Makefile
:category: C
:summary: Remove trailing newline (carriage return) of a string in C programming language.


The following code shows how to remove trailing `newline <http://en.wikipedia.org/wiki/Newline>`_ (carriage return) of a string in C:

.. show_github_file:: siongui userpages content/articles/2013/01/09/remove-string-trailing-newline.c

The following test code shows how to use above function:

.. show_github_file:: siongui userpages content/articles/2013/01/09/test-remove-string-trailing-newline.c

Makefile for building above code:

.. show_github_file:: siongui userpages content/articles/2013/01/09/Makefile

Run **make** to build the executable ``main``. Then run the executable ``main`` in your terminal and the output will be like:

.. code-block:: bash

  $ ./main 
  abcd
  abcd$ 

Because the trailing newline of the string is removed, the terminal prompt sign **$** is shown on the right of *abcd* instead of below *abcd*.
