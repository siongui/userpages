[Vim] Replace Footnote by Regular Expression
############################################

:date: 2016-03-28T06:30+08:00
:tags: Bash, String Manipulation, Regular Expression, Vim, reStructuredText
:category: Bash
:summary: Replace footnote in reStructuredText_ format via `regular expression`_
          in Vim_ editor.


Replace footnote in reStructuredText_ format via `regular expression`_ in Vim_
editor.

``(註1-1)`` => `` [1]\_ ``

.. code-block:: vim

  :%s/(註1-\(\d\{1}\))/ [\1]_ /gc


``〔註1-1〕：`` => `` .. [1] ``

.. code-block:: vim

  :%s/^〔註1-\(\d\{1}\)〕：/.. [\1] /gc


``[001]`` => ``[001]_``

.. code-block:: vim

  :%s/\[\(\d\{3}\)\]/\[\1\]_/gc


``001`` => ``[001]_``

.. code-block:: vim

  :%s/\(\d\{3}\)/\[\1\]_/gc
  :%s/\(0\d\{2}\)/\[\1\]_/gc


``)[001]_`` => ``) [001]_``

.. code-block:: vim

  :%s/)\(\[\d\{3}]_\)/) \1/gc


``〔註001〕`` => `` .. [001] ``

.. code-block:: vim

  :%s/^〔註\(0\d\{2}\)〕/.. [\1] /gc

----

Tested on ``Ubuntu Linux 15.10``, ``vim 2:7.4.712-2ubuntu4``.

----

References:

.. [1] `vim regex replace <https://www.google.com/search?q=vim+regex+replace>`_

       `Vim Regex Replace - Stack Overflow <http://stackoverflow.com/questions/11850033/vim-regex-replace>`_

.. _Vim: http://www.vim.org/
.. _regular expression: https://www.google.com.tw/search?q=regular+expression
.. _reStructuredText: https://www.google.com.tw/search?q=reStructuredText

.. ``(註1-1)`` => `` [1]_ ``
   :%s/(註1-\(\d\{1}\))/ [\1]_ /gc

.. ``〔註1-1〕：`` => `` .. [1] ``
   :%s/^〔註1-\(\d\{1}\)〕：/.. [\1] /gc
