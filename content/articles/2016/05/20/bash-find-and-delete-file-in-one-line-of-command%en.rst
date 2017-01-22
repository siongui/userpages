[Bash] Find and Remove Files in One Line of Command
###################################################

:date: 2016-05-20T19:43+08:00
:tags: Bash, Commandline, find command
:category: Bash
:summary: Use find_ command to find all files with specific name and delete them
          in one line of command.
:adsu: yes


Use find_ command to find all files with specific name and delete them in one
line of command.

**Quesion**

Find all files with the name ``WS_FTP.LOG`` in ``content/extra`` directory, and
remove them at the same time.

**Answer**

.. code-block:: sh

  $ find content/extra/ -name "WS_FTP.LOG" | xargs rm

----

Tested on: ``Ubuntu Linux 16.04``, ``find (GNU findutils) 4.7.0-git``

----

References:

.. [1] `unix pipe delete file - Google search <https://www.google.com/search?q=unix+pipe+delete+file>`_

       `unix - Command line: piping find results to rm - Stack Overflow <http://stackoverflow.com/questions/11191475/command-line-piping-find-results-to-rm>`_


.. _Bash: https://www.google.com/search?q=Bash
.. _HTML: https://www.google.com/search?q=HTML
.. _find: https://www.gnu.org/software/findutils/manual/html_mono/find.html
