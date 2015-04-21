如何透過GitHub在本站寫文章
##########################

:date: 2015-04-21 19:17
:tags: Web開發
:category: Web開發
:summary: 透過GitHub平台協同撰寫網站


流程
++++

1. 首先註冊 GitHub_ 帳號

2. fork_ 本站 repo_, 利用 rst_ 格式寫文章，然後 `pull request`_ 提交。


如何寫文章
++++++++++

透過範例學習最快，本站每篇文章都有一個 ``在GitHub上編輯`` 的連結，
可連回到原始檔案，然後在點選 ``raw`` 看原始 rst_ 格式。

舉例來說，2015年4月21日的文章，就在
**content/2015/04/21/** 目錄(若該目錄不存在則自己創建)下新增一個該文章的檔案
，檔名取為

**SLUG%LANG.rst** ，其中 *SLUG* 是網址上的顯示名稱， *LANG* 是語言

舉例來說，若新增一個檔案名為 **content/articles/2015/04/21/post-on-this-site%zh.rst**
則該文章的網址是 **/zh/2015/04/21/post-on-this-site/** ，該文章的語言是傳統中文

若新增一個檔案名為 **content/articles/2015/04/21/post-on-this-site%en.rst**
則該文章的網址是 **/2015/04/21/post-on-this-site/** ，該文章的語言是英文

文章內容則長得像:

.. code-block:: txt

  如何透過GitHub在本站寫文章
  ##########################

  :date: 2015-04-21 19:17
  :author: 某某某
  :tags: Web開發
  :category: Web開發
  :summary: 透過GitHub平台協同撰寫網站

  ...(your main content here)

*category* 只可以有一個， tag可以有好幾個， date可加可不加，
不加的話則是用目錄裡的日期當成此文章日期

rst_ 格式怎樣寫可參考 [1]_ ，至於用LaTeX寫數學，可看 [2]_


注意事項：

- 每行建議不超過80個字母，一個中文算兩個字母。（非硬性規定）

- 若不確定LaTex顯示出來如何，可先在 [3]_ 輸入看結果。
  (參考 [4]_ 輸入數學符號)

- 可參考 [5]_ 來寫 rst_


預覽寫好的文章
++++++++++++++

本站目前只能在 Ubuntu Linux 上compile出來並預覽，詳情請看：
`README <https://github.com/siongui/userpages/blob/master/README.rst>`_ 。
Windows上理論上應該可以compile出來並預覽，但從沒試過。

----

參考：

.. [1] `reStructuredText <http://docutils.sourceforge.net/rst.html>`_

.. [2] `reStructuredText Directives - math <http://docutils.sourceforge.net/docs/ref/rst/directives.html#math>`_

.. [3] `Online LaTeX Equation Editor - create, integrate and download <http://www.codecogs.com/latex/eqneditor.php>`_

.. [4] `LaTeX/Mathematics - Wikibooks, open books for an open world <http://en.wikibooks.org/wiki/LaTeX/Mathematics>`_

.. [5] `7. 附录：轻量级标记语言 — GotGitHub <http://www.worldhello.net/gotgithub/appendix/markups.html>`_


.. _GitHub: https://github.com/
.. _fork: https://help.github.com/articles/fork-a-repo/
.. _repo: https://github.com/siongui/userpages
.. _rst: http://docutils.sourceforge.net/rst.html
.. _pull request: https://help.github.com/articles/using-pull-requests/
