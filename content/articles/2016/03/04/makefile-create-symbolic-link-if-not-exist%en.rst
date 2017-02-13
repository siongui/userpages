[Makefile] Create Symbolic Link If Not Exist
############################################

:date: 2016-03-04T18:09+08:00
:tags: Bash, Makefile, Commandline
:category: Bash
:summary: Makefile_ - Create symbolic link if it does not exist.
          Otherwise do nothing.
:adsu: yes


Create a symbolic link to directory if the symlink does not exist in Makefile_.
Otherwise do nothing.

.. code-block:: bash

  [ -L tipitaka/common ] || (cd tipitaka; ln -s ../common/ common)

----

Tested on ``Ubuntu Linux 15.10``.

----

.. [1] `makefile check if symlink exists <https://www.google.com/search?q=makefile+check+if+symlink+exists>`_

       `bash - How to check if symlink exists - Stack Overflow <http://stackoverflow.com/questions/5767062/how-to-check-if-symlink-exists>`_

       `linux - Using makefile check if symlink exists and create if it doesn't - Stack Overflow <http://stackoverflow.com/questions/29072366/using-makefile-check-if-symlink-exists-and-create-if-it-doesnt>`_

       `Makefile target with wildcard to create symlink using target string - Stack Overflow <http://stackoverflow.com/questions/21064718/makefile-target-with-wildcard-to-create-symlink-using-target-string>`_
.. adsu:: 2
.. [2] `use Makefile for setup and development (unfinished) · siongui/pali@333e8c5 · GitHub <https://github.com/siongui/pali/commit/333e8c570959707f620c612e1b6494d3fe5696f7>`_

.. _Makefile: https://www.google.com/search?q=Makefile
