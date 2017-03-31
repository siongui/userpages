[Makefile] Publish to GitHub Pages Automatically
################################################

:date: 2017-03-31T23:21+08:00
:tags: Bash, Makefile, Commandline, GitHub Pages
:category: Bash
:summary: Create *gh-pages* branch and publish your static website to GitHub
          Pages in Makefile_.
:adsu: yes


Assume we have the following repo:

.. code-block:: txt

  https://github.com/myname/myrepo

We want to automate the publish of static website to GitHub Pages.

Here is what I do:

First create a ``gh-pages`` branch:

.. code-block:: txt

  # clone repo
  $ git clone https://github.com/myname/myrepo.git
  $ cd myrepo
  # create *gh-pages* branch
  $ git checkout -b gh-pages
  # push *gh-pages* on GitHub
  $ git push origin gh-pages

Then use the following code to push static contents to GitHub Pages:

.. code-block:: makefile

  REPO=github.com/myname/myrepo
  REPO_DIR=/tmp/$(REPO)

  publish:
  	@echo "\033[92mClone $(REPO) ...\033[0m"
  	rm -rf $(REPO_DIR)
  	git clone https://$(REPO).git $(REPO_DIR) --depth=1
  	@echo "\033[92mGit checkout gh-pages ...\033[0m"
  	cd $(REPO_DIR); git checkout gh-pages; rm -rf *
  	# Build you contents here
  	@echo "\033[92mPush to Remote Repo ...\033[0m"
  	cd $(REPO_DIR); git add .; git commit -m "update site"; git push

.. adsu:: 2

----

Tested on `Ubuntu Linux 16.10`_, `GNU make 4.1-9`_.

----

References:

.. [1] | `gh-pages push - Google search <https://www.google.com/search?q=gh-pages+push>`_
       | `gh-pages push - DuckDuckGo search <https://duckduckgo.com/?q=gh-pages+push>`_
       | `gh-pages push - Ecosia search <https://www.ecosia.org/search?q=gh-pages+push>`_
       | `gh-pages push - Qwant search <https://www.qwant.com/?q=gh-pages+push>`_
       | `gh-pages push - Bing search <https://www.bing.com/search?q=gh-pages+push>`_
       | `gh-pages push - Yahoo search <https://search.yahoo.com/search?p=gh-pages+push>`_
       | `gh-pages push - Baidu search <https://www.baidu.com/s?wd=gh-pages+push>`_
       | `gh-pages push - Yandex search <https://www.yandex.com/search/?text=gh-pages+push>`_

.. [2] | `git create branch - Google search <https://www.google.com/search?q=git+create+branch>`_
       | `git create branch - DuckDuckGo search <https://duckduckgo.com/?q=git+create+branch>`_
       | `git create branch - Ecosia search <https://www.ecosia.org/search?q=git+create+branch>`_
       | `git create branch - Qwant search <https://www.qwant.com/?q=git+create+branch>`_
       | `git create branch - Bing search <https://www.bing.com/search?q=git+create+branch>`_
       | `git create branch - Yahoo search <https://search.yahoo.com/search?p=git+create+branch>`_
       | `git create branch - Baidu search <https://www.baidu.com/s?wd=git+create+branch>`_
       | `git create branch - Yandex search <https://www.yandex.com/search/?text=git+create+branch>`_

.. [3] `Create a new branch with git and manage branches · Kunena/Kunena-Forum Wiki · GitHub <https://github.com/Kunena/Kunena-Forum/wiki/Create-a-new-branch-with-git-and-manage-branches>`_
.. adsu:: 3
.. [4] | `makefile change directory - Google search <https://www.google.com/search?q=makefile+change+directory>`_
       | `makefile change directory - DuckDuckGo search <https://duckduckgo.com/?q=makefile+change+directory>`_
       | `makefile change directory - Ecosia search <https://www.ecosia.org/search?q=makefile+change+directory>`_
       | `makefile change directory - Qwant search <https://www.qwant.com/?q=makefile+change+directory>`_
       | `makefile change directory - Bing search <https://www.bing.com/search?q=makefile+change+directory>`_
       | `makefile change directory - Yahoo search <https://search.yahoo.com/search?p=makefile+change+directory>`_
       | `makefile change directory - Baidu search <https://www.baidu.com/s?wd=makefile+change+directory>`_
       | `makefile change directory - Yandex search <https://www.yandex.com/search/?text=makefile+change+directory>`_

.. _Makefile: https://www.google.com/search?q=Makefile
.. _Ubuntu Linux 16.10: http://releases.ubuntu.com/16.10/
.. _GNU make 4.1-9: https://www.gnu.org/software/make/
