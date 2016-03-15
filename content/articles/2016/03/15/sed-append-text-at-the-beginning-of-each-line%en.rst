[sed] Append Text at the Beginning of Each Line
###############################################

:date: 2016-03-15T19:42+08:00
:tags: Bash, Commandline, sed, String Manipulation
:category: Bash
:summary: Append text at the beginning of each line via sed_ stream editor.


Append text at the beginning of each line via sed_ stream editor.

Append ``| `` to the beginning of each line.

.. code-block:: bash

  sed 's/^/| /' tmp.txt > tmp2.txt


----

Tested on ``Ubuntu Linux 15.10``, ``sed 4.2.2-6.1``.

----

References:

.. [1] `sed append to end of line <https://www.google.com/search?q=sed+append+to+end+of+line>`_

       `sed - add text at the end of each line - Stack Overflow <http://stackoverflow.com/questions/15978504/add-text-at-the-end-of-each-line>`_


.. _sed: http://www.grymoire.com/Unix/Sed.html
