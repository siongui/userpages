[Bash] List Directory Size in Descending and Ascending Order
############################################################

:date: 2018-03-22T22:27+08:00
:tags: Bash, Commandline, du command, sort command
:category: Bash
:summary: List folder size in descending and ascending order via *du* and *sort*
          command in *bash*.
:og_image: https://www.tecmint.com/wp-content/uploads/2016/01/Find-Largest-Files-Directories-Size-in-Linux.png
:adsu: yes

Use du_ and sort_ command to list directory size in descending and ascending
order.

List Folder Size (Descending)
++++++++++++++++++++++++++++++

.. code-block:: bash

  $ du -s * | sort -rg


List Folder Size (Ascending)
++++++++++++++++++++++++++++

.. code-block:: bash

  $ du -s * | sort -g


.. adsu:: 2

----

References:

.. [1] `how to list folder size in descending order ? <https://www.linuxquestions.org/questions/linux-server-73/how-to-list-folder-size-in-descending-order-787889/>`_

.. _du: https://linux.die.net/man/1/du
.. _sort: http://man7.org/linux/man-pages/man1/sort.1.html
