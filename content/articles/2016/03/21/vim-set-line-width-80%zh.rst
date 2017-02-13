Vim編輯器設定寬度80的分隔線
###########################

:date: 2016-03-21T03:46+08:00
:tags: Linux, Vim
:category: Linux
:summary: Vim_ 編輯器設定寬度80的分隔線
:adsu: yes


Vim_ 編輯器設定寬度80的分隔線，
從 `SO的回答 <http://stackoverflow.com/a/3765575>`_ [1]_
，在 ``.vimrc`` 裡加入：

.. code-block:: vim

  if exists('+colorcolumn')
    set colorcolumn=80
  else
    au BufWinEnter * let w:m2=matchadd('ErrorMsg', '\%>80v.\+', -1)
  endif

.. adsu:: 2

----

References:

.. [1] | `vim set line width 80 <https://www.google.com/search?q=vim+set+line+width+80>`_
       | `coding style - Vim 80 column layout concerns - Stack Overflow <http://stackoverflow.com/questions/235439/vim-80-column-layout-concerns>`_

.. _Vim: http://www.vim.org/
