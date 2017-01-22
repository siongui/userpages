重設GitHub repository並移除所有commits
######################################

:date: 2016-03-19T03:22+08:00
:tags: Linux, git
:category: Linux
:summary: 重設GitHub repository並移除所有commits
:adsu: yes

從 [1]_ 的回答，假設 repository 位於 ``~/dev/myrepo`` ：

.. code-block:: bash

  # 移除所有紀錄
  $ cd ~/dev/myrepo
  $ rm -rf .git
  # 重新造git repository
  $ git init
  # 將所有檔案加入repository
  $ git add .
  $ git commit -m "initial commit"
  # 重設GitHub上的repository，此動作不可復原，做之前請先再次確認
  $ git remote add origin <url>
  $ git push --force --set-upstream origin master

----

References:

.. [1] `version control - How to reset a remote GIT repository to remove all commits? - Stack Overflow <http://stackoverflow.com/questions/2006172/how-to-reset-a-remote-git-repository-to-remove-all-commits>`_
