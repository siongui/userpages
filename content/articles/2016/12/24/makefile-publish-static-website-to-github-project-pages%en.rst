[Makefile] Publish Static Website to GitHub Project Pages
#########################################################

:date: 2016-12-24T21:08+08:00
:tags: Bash, Makefile, Commandline, GitHub Pages, apt command
:category: Bash
:summary: Publish your static website to GitHub project pages via ghp-import_ in
          Makefile_.
:adsu: yes


First install ghp-import_ via pip_ (`Ubuntu Linux 16.10`_ as example):

.. code-block:: bash

  $ sudo apt-get install python-pip
  $ sudo pip install ghp-import

Then in Makefile_ (assume your static website are located in ``/home/usrname/dev/mywebsite/``):

.. code-block:: makefile

  publish_to_github_project_pages:
  	cd /home/usrname/dev/; ghp-import mywebsite; git push origin gh-pages

.. adsu:: 2

Note that I put several commands in one line.
If your Makefile is not located in ``/home/usrname/dev/``, the following will
not work:

.. code-block:: makefile

  publish_to_github_project_pages:
  	# wrong way
  	cd /home/usrname/dev/
  	ghp-import mywebsite
  	git push origin gh-pages

Because when GNU make executes every line of command, it executes the command at
current directory of of Makefile_. So when ghp-import_ try to find ``mywebsite``
directory in current directory, it will find nothing and report error.

.. adsu:: 3

----

Tested on `Ubuntu Linux 16.10`_, `GNU make 4.1-9`_.

----

References:

.. [1] | `makefile change directory - Google search <https://www.google.com/search?q=makefile+change+directory>`_
       | `makefile change directory - DuckDuckGo search <https://duckduckgo.com/?q=makefile+change+directory>`_
       | `makefile change directory - Ecosia search <https://www.ecosia.org/search?q=makefile+change+directory>`_
       | `makefile change directory - Bing search <https://www.bing.com/search?q=makefile+change+directory>`_
       | `makefile change directory - Yahoo search <https://search.yahoo.com/search?p=makefile+change+directory>`_
       | `makefile change directory - Baidu search <https://www.baidu.com/s?wd=makefile+change+directory>`_
       | `makefile change directory - Yandex search <https://www.yandex.com/search/?text=makefile+change+directory>`_
       | `linux - How do I write the 'cd' command in a makefile? - Stack Overflow <http://stackoverflow.com/questions/1789594/how-do-i-write-the-cd-command-in-a-makefile>`_
.. adsu:: 4
.. [2] `Publishing to GitHub - Tips — Pelican 3.7.1.dev0 documentation <http://docs.getpelican.com/en/latest/tips.html#publishing-to-github>`_


.. _Makefile: https://www.google.com/search?q=Makefile
.. _Ubuntu Linux 16.10: http://releases.ubuntu.com/16.10/
.. _GNU make 4.1-9: https://www.gnu.org/software/make/
.. _ghp-import: https://github.com/davisp/ghp-import
.. _pip: https://pypi.python.org/pypi/pip
