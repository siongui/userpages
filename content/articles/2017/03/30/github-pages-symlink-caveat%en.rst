GitHub Pages Symbolic Link Caveat
#################################

:date: 2017-03-30T23:11+08:00
:tags: GitHub Pages, web, Web application
:category: Web Development
:summary: GitHub Pages set *Content-Type* according to the name of symlink, not
          the file name that the symlink references to.
:adsu: yes

`GitHub Pages`_ set *Content-Type* according to the name of symlink, not the file
name that the symlink references to.

I have a URL of word as follows:

.. code-block:: txt

  /browse/s/sacca

I want this url linked to ``/``, so I make a symlink as follows:

.. code-block:: bash

  $ cd {{root of websit}}
  $ mkdir -p browse/s/
  $ ln -sf ../../index.html browse/s/sacca

But this does not work as expected. When I open ``/browse/s/sacca``, my browser
downloads the file instead of rendering it as HTML, because the *Content-Type*
is not ``text/html``.

My solution is to add a trailing slash to the URL of the word as follows:

.. code-block:: txt

  /browse/s/sacca/

And make symlink as follows:

.. code-block:: bash

  $ cd {{root of websit}}
  $ mkdir -p browse/s/sacca/
  $ ln -sf ../../../index.html browse/s/sacca/index.html

I open ``/browse/s/sacca/``, and my browser renders it as HTML!

.. adsu:: 2

----

References:

.. [1] | `github pages symlink - Google search <https://www.google.com/search?q=github+pages+symlink>`_
       | `github pages symlink - DuckDuckGo search <https://duckduckgo.com/?q=github+pages+symlink>`_
       | `github pages symlink - Ecosia search <https://www.ecosia.org/search?q=github+pages+symlink>`_
       | `github pages symlink - Qwant search <https://www.qwant.com/?q=github+pages+symlink>`_
       | `github pages symlink - Bing search <https://www.bing.com/search?q=github+pages+symlink>`_
       | `github pages symlink - Yahoo search <https://search.yahoo.com/search?p=github+pages+symlink>`_
       | `github pages symlink - Baidu search <https://www.baidu.com/s?wd=github+pages+symlink>`_
       | `github pages symlink - Yandex search <https://www.yandex.com/search/?text=github+pages+symlink>`_

.. [2] | `git gh-pages workflow - Google search <https://www.google.com/search?q=git+gh-pages+workflow>`_
       | `git gh-pages workflow - DuckDuckGo search <https://duckduckgo.com/?q=git+gh-pages+workflow>`_
       | `git gh-pages workflow - Ecosia search <https://www.ecosia.org/search?q=git+gh-pages+workflow>`_
       | `git gh-pages workflow - Qwant search <https://www.qwant.com/?q=git+gh-pages+workflow>`_
       | `git gh-pages workflow - Bing search <https://www.bing.com/search?q=git+gh-pages+workflow>`_
       | `git gh-pages workflow - Yahoo search <https://search.yahoo.com/search?p=git+gh-pages+workflow>`_
       | `git gh-pages workflow - Baidu search <https://www.baidu.com/s?wd=git+gh-pages+workflow>`_
       | `git gh-pages workflow - Yandex search <https://www.yandex.com/search/?text=git+gh-pages+workflow>`_

.. _GitHub Pages: https://pages.github.com/
