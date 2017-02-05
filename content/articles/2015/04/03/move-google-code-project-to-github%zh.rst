將Google Code上的專案移到Github
###############################

:date: 2015-04-03 02:13
:tags: Web開發
:category: Web開發
:summary: 如何將Google Code上的專案移到Github。
:adsu: yes


`Google Code`_ 要關了，想把原本放在Google Code上的 git_ 專案移到
GitHub_ 上，該怎樣做呢？

經過 Google 搜尋 `clone google code to github`_ ，發現 [1]_
裡頭有很清楚的步驟解說，大致步驟為：

1. 先在GitHub上 `建一個新repository`_
   (假設新repository名稱為 *obsoleted-pali*)

2. 將原本在Google Code上的專案clone下來：
   (假設專案名稱為 *pali*)

   .. code-block:: bash

     $ git clone https://code.google.com/p/pali/

3. 進入clone下來的專案，設定remote origin url：

   .. code-block:: bash

     $ git remote set-url origin https://github.com/USERNAME/obsoleted-pali.git

4. push到GitHub上：

   .. code-block:: bash

     $ git push origin master

.. adsu:: 2

----

參考：

.. [1] `Migrating a git based project from Google Code to Github · Issue #347 · holman/feedback · GitHub <https://github.com/holman/feedback/issues/347>`_


.. _clone google code to github: https://www.google.com/search?q=clone+google+code+to+github

.. _Google Code: https://code.google.com/

.. _git: http://git-scm.com/

.. _GitHub: https://github.com/

.. _建一個新repository: https://help.github.com/articles/creating-a-new-repository/
