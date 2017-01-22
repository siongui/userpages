[sed] Replace reStructuredText Link by Regular Expression
#########################################################

:date: 2016-03-30T21:21+08:00
:tags: Bash, Commandline, sed, String Manipulation, Regular Expression,
       reStructuredText
:category: Bash
:summary: Make reStructuredText_ links of the same name anonymous via
          `regular expression`_ by sed_ stream editor.
:adsu: yes


Make reStructuredText_ links of the same name anonymous via
`regular expression`_ by sed_ stream editor.

```頁4 <http://www.chilin.edu.hk/edu/report_section_detail.asp?section_id=59&id=274&page_id=156:0>`_``
=>
```頁4 <http://www.chilin.edu.hk/edu/report_section_detail.asp?section_id=59&id=274&page_id=156:0>`__``

.. code-block:: bash

  $ sed -r 's/(`頁4 <[^>]+>`)_/\1__/g' -i content/articles/tipitaka/sutta/diigha/dn22/contrast-reading%zh.rst

Result:

.. show_github_file:: siongui userpages content/articles/2016/03/30/link.diff

Note that use ``'s/(`頁4 <.+>`)_/\1__/g'`` instead of
``'s/(`頁4 <[^>]+>`)_/\1__/g'`` will match only one link because the match is
greedy. See [1]_ for more details.

----

Tested on ``Ubuntu Linux 15.10``, ``sed 4.2.2-6.1``.

----

References:

.. [1] `sed match non greedy <https://www.google.com/search?q=sed+match+non+greedy>`_

       `Non greedy regex matching in sed? - Stack Overflow <http://stackoverflow.com/questions/1103149/non-greedy-regex-matching-in-sed>`_


.. _sed: http://www.grymoire.com/Unix/Sed.html
.. _regular expression: https://www.google.com.tw/search?q=regular+expression
.. _reStructuredText: https://www.google.com.tw/search?q=reStructuredText
