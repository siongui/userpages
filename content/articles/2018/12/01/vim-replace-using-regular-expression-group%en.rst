[Vim] Replace Using Regular Expression Group
############################################

:date: 2018-12-01T00:36+08:00
:tags: Bash, String Manipulation, Regular Expression, Vim
:category: Bash
:summary: Use regular expresion groups to replace texts in Vim.
:og_image: https://image.slidesharecdn.com/pycon2016-160814073914/95/pycon-apac-2016-regular-expressionaz-34-638.jpg
:adsu: yes

**Question**

  I want to turn the following texts:

  .. code-block:: txt

    C413 TYPE
    C414 TYPE
    C415 TYPE

  into:

  .. code-block:: txt

    平鐵板 C413
    平鐵板 C414
    平鐵板 C415

  How to do it in Vim_ editor?

**Answer**

  I know that `regular expression`_ groups can help me solve the problem, so I
  make some DuckDuckGo search [1]_ and found that someone have the same problem
  as me [2]_. I follow the example and come up with my own answer:

  .. code-block:: vim

    :%s/\(C\d\d\d\) TYPE/平鐵板 \1/gc

.. adsu:: 2

----

Tested on ``Ubuntu Linux 18.04``, ``vim 2:8.0.1453-1ubuntu1``.

----

References:

.. [1] `vim replace with named group at DuckDuckGo <https://duckduckgo.com/?q=vim+replace+with+named+group>`_
.. [2] `Vim Regex Capture Groups [bau -> byau : ceu -> cyeu] - Stack Overflow <https://stackoverflow.com/questions/19902089/vim-regex-capture-groups-bau-byau-ceu-cyeu>`_

.. _Vim: http://www.vim.org/
.. _regular expression: https://duckduckgo.com/?q=regular+expression

.. :%s/\(C\d\d\d\) TYPE/平鐵板 \1/gc
