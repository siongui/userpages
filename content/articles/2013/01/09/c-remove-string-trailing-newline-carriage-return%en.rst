[C] Remove String Trailing Newline (Carriage Return)
####################################################

:tags: C, remove trailing newline, remove carriage return, String Manipulation,
       Makefile
:category: C
:summary: Remove trailing newline (carriage return) of a string in
          `C programming language`_.
:adsu: yes


The following code shows how to remove trailing newline_ (carriage return) of a
string in C_:

.. show_github_file:: siongui userpages content/code/c/remove-trailing-newline/remove-string-trailing-newline.c
.. adsu:: 2

The following test code shows how to use above function:

.. show_github_file:: siongui userpages content/code/c/remove-trailing-newline/test-remove-string-trailing-newline.c
.. adsu:: 3

Makefile_ for building above code:

.. show_github_file:: siongui userpages content/code/c/remove-trailing-newline/Makefile

Run **make** to build the executable ``main``. Then run the executable ``main``
in your terminal and the output will be like:

.. code-block:: bash

  $ ./main 
  abcd
  abcd$ 

Because the trailing newline_ of the string is removed, the terminal prompt sign
**$** is shown on the right of *abcd* instead of below *abcd*.

.. _C: https://www.google.com/search?q=C+programming
.. _Makefile: https://www.google.com/search?q=Makefile
.. _C programming language: https://www.google.com/search?q=C+programming+language
.. _newline: https://en.wikipedia.org/wiki/Newline
