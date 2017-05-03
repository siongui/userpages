在Windows下安裝Pelican靜態網站產生器
####################################

:date: 2017-05-03T17:22:00+08:00
:tags: Web開發
:category: Web開發
:summary: 在Windows作業系統下，安裝Python的Pelican靜態網站產生器。
:adsu: yes


1. 安裝 Cygwin_ （ `64位元安裝檔 <https://www.cygwin.com/setup-x86_64.exe>`_ ）

   安裝時，請將 **Python** 以及 **Devel** 選成 *Install* （不要選 *Default* ）

2. 安裝須花頗多時間，完成後，在桌面建立一個 Cygwin Terminal 的捷徑。

3. 點擊捷徑進入 Cygwin Terminal，進入後先檢查 Python_ 以及 git_ 是否安裝成功。

   先檢查 Python_ ：

   .. code-block:: bash

     $ python -V

   如果成功會看到 Python_ 的版本號碼。

   接著檢查 git_ ：

   .. code-block:: bash

     $ git --version

   如果成功會看到 git_ 的版本號碼。
   接著設定user名稱跟email：

   .. code-block:: bash

     $ git config --global user.name "你的GitHub使用者名稱"
     $ git config --global user.email "你的Email"

   請依照自己的帳號及email做設定。

4. 安裝 pip_ [2]_ ：

   .. code-block:: bash

     $ easy_install-2.7 pip

   移除有問題的套件：

   .. code-block:: bash

     $ pip uninstall stgit
     $ pip uninstall bzr-fastimport

5. 進入原始碼目錄，用 pip_ 安裝需要的 Python 套件：

   .. code-block:: bash

     $ pip install pelican
     $ pip install ghp-import
     $ pip install pyScss

----

附錄：

.. [1] `7. 附录：轻量级标记语言 — GotGitHub <http://www.worldhello.net/gotgithub/appendix/markups.html>`_
       (`GitHub <https://github.com/gotgit/gotgithub/blob/master/appendix/markups.rst>`__)

.. [2] `python - Installing Pip-3.2 on Cygwin - Stack Overflow <http://stackoverflow.com/a/30685412>`_

.. _GitHub Pages: https://pages.github.com/
.. _static site generator: https://www.google.com/search?q=static+site+generator
.. _Pelican: http://blog.getpelican.com/
.. _Pelican官方文件: http://docs.getpelican.com/en/3.5.0/content.html
.. _Cygwin: https://www.cygwin.com/
.. _Python: https://www.python.org/
.. _git: https://git-scm.com/
.. _pip: https://pypi.python.org/pypi/pip
