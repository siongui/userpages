[Bash] find Command Exclude Sub-directory
#########################################

:date: 2016-05-19T07:39+08:00
:tags: Bash, Commandline, find command, List Files in Directory
:category: Bash
:summary: Use find_ command to find all HTML_ files in a directory, but exclude
          the HTML files in some sub-directory.
:adsu: yes


Use find_ command to find all HTML_ files in a directory, but exclude the HTML
files in some sub-directory.

**Quesion**

Find all HTML files in ``output`` directory, but exclude the HTML files in
``output/extra`` and ``output/tag`` sub-directories.


**Answer**

.. code-block:: sh

  find output/ -not \( -path "output/extra/*" -o -path "output/tag/*" \) -name "*.html"

----

Tested on: ``Ubuntu Linux 16.04``, ``find (GNU findutils) 4.7.0-git``

----

References:

.. [1] `find exclude subdirectory - Google search <https://www.google.com/search?q=find+exclude+subdirectory>`_

       `linux - Exclude directory from find . command - Stack Overflow <http://stackoverflow.com/questions/4210042/exclude-directory-from-find-command>`_


.. _Bash: https://www.google.com/search?q=Bash
.. _HTML: https://www.google.com/search?q=HTML
.. _find: https://www.gnu.org/software/findutils/manual/html_mono/find.html
