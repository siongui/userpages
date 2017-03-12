GitHub Pages的index.html無法顯示
################################

:date: 2017-03-13T05:42+08:00
:tags: Linux, git, Web開發
:category: Linux
:summary: 用 `GitHub Pages`_ 架站，但 index.html 無法顯示，一直回報 404
:adsu: yes

用 `GitHub Pages`_ 架站，但 index.html 無法顯示，一直回報 404 。 經過搜尋 [1]_
後，發現解答 [2]_ 可以解決這個問題：

.. code-block:: bash

  # 刪除GitHub上的 gh-pages branch
  $ git push origin --delete gh-pages
  # 刪除本機上的 gh-pages branch
  $ git branch -D gh-pages
  # 重新初始化 git
  $ git init
  # 重新初始化本機的 gh-pages branch
  $ git branch gh-pages
  # 將本機的 gh-pages push到 GitHub
  $ git push origin gh-pages

.. adsu:: 2

----

References:

.. [1] | `github pages index.html not working - Google search <https://www.google.com/search?q=github+pages+index.html+not+working>`_
       | `github pages index.html not working - DuckDuckGo search <https://duckduckgo.com/?q=github+pages+index.html+not+working>`_
       | `github pages index.html not working - Ecosia search <https://www.ecosia.org/search?q=github+pages+index.html+not+working>`_
       | `github pages index.html not working - Qwant search <https://www.qwant.com/?q=github+pages+index.html+not+working>`_
       | `github pages index.html not working - Bing search <https://www.bing.com/search?q=github+pages+index.html+not+working>`_
       | `github pages index.html not working - Yahoo search <https://search.yahoo.com/search?p=github+pages+index.html+not+working>`_
       | `github pages index.html not working - Baidu search <https://www.baidu.com/s?wd=github+pages+index.html+not+working>`_
       | `github pages index.html not working - Yandex search <https://www.yandex.com/search/?text=github+pages+index.html+not+working>`_

.. [2] `How to fix page 404 on Github Page? - Stack Overflow <http://stackoverflow.com/a/13812675>`_

.. _GitHub Pages: https://pages.github.com/
