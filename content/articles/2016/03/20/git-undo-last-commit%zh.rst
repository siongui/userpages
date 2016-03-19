git復原最後一次的commit
#######################

:date: 2016-03-20T05:19+08:00
:tags: Linux, git
:category: Linux
:summary: git復原最後一次的commit

從 [1]_ 的回答：

.. code-block:: bash

  # 重設上次修改
  $ git reset --soft HEAD~
  # 編輯並再次git add修改，完成後：
  $ git commit -c ORIG_HEAD

----

References:

.. [1] `git reset commit <https://www.google.com/search?q=git+reset+commit>`_

       `git - How do you undo the last commit? - Stack Overflow <http://stackoverflow.com/questions/927358/how-do-you-undo-the-last-commit>`_
