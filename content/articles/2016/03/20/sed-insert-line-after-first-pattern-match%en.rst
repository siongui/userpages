[sed] Insert Line After First Pattern Match
###########################################

:date: 2016-03-20T04:09+08:00
:tags: Bash, Commandline, sed, String Manipulation, Regular Expression,
       reStructuredText
:category: Bash
:summary: Insert a line after first pattern match via sed_ stream editor.
:adsu: yes


Insert a line after first pattern match via sed_ stream editor.

  In the files with names starting with **dhp-story** under current directory,
  find the line that starts with **:oldurl:**, insert the following line after
  the matched line:

  **:author: 護法法師**

.. code-block:: bash

  sed -i '/:oldurl:/a :author: 護法法師' dhp-story*

----

Tested on ``Ubuntu Linux 15.10``, ``sed 4.2.2-6.1``.

----

References:

.. [1] `sed insert line after match <https://www.google.com/search?q=sed+insert+line+after+match>`_

       `shell - Insert line after first match using sed - Stack Overflow <http://stackoverflow.com/questions/15559359/insert-line-after-first-match-using-sed>`_

.. [2] `[Python] Insert Line With Matched Pattern <{filename}../01/python-insert-line-with-matched-pattern%en.rst>`_

.. _sed: http://www.grymoire.com/Unix/Sed.html
