[Pelican] Create URL Based on File Path
#######################################

:date: 2018-10-17T20:14+08:00
:tags: Python, Pelican
:category: Python
:summary: *Pelican* static site generator - Create URL of *page* or *article*
          based on the full path relative to the content source directory.
:og_image: https://lazybatman.com/wp-content/uploads/2017/07/pelican.png
:adsu: yes


Pelican_ static site generator - Create URL of page_ or article_ based on the
full path relative to the content source directory.

**Problem**

for example, the following is the path of a page:

.. code-block:: txt

  content/pages/intro.rst

The URL of the page will be:

.. code-block:: txt

  /intro/

The following is the path of a article:

.. code-block:: txt

  content/articles/category/myproduct.rst

The URL of the article will be:

.. code-block:: txt

  /category/myproduct/

|

**Solution**

We will create a custom metadata_ field called *urlpath*. The value of *urlpath*
is extracted from article/page's path relative to the content source directory.
Then the value of *urlpath* will be used to determine the URL of the
article/page. You do not need to modify the file content of articles/pages.

The solution is to set *PATH_METADATA*, *PAGE_URL*, *PAGE_SAVE_AS*,
*ARTICLE_URL*, *ARTICLE_SAVE_AS* in `pelicanconf.py`_ as follows:

.. code-block:: python

  PATH_METADATA = '(articles|pages)/(?P<urlpath>[-a-zA-Z0-9/]*)\.rst'

  ARTICLE_URL = '{urlpath}/'
  ARTICLE_SAVE_AS = '{urlpath}/index.html'

  PAGE_URL = '{urlpath}/'
  PAGE_SAVE_AS = '{urlpath}/index.html'

----

Tested on: ``Ubuntu Linux 18.04``, ``Python 2.7.15rc1``, `Pelican 3.7.1`_.

----

.. adsu:: 2

References:

.. [1] `Settings — Pelican 3.7.1 documentation <http://docs.getpelican.com/en/stable/settings.html>`_

.. _Pelican: http://blog.getpelican.com/
.. _Pelican 3.7.1: http://docs.getpelican.com/en/3.7.1/
.. _page: http://docs.getpelican.com/en/stable/content.html#articles-and-pages
.. _article: http://docs.getpelican.com/en/stable/content.html#articles-and-pages
.. _metadata: http://docs.getpelican.com/en/stable/content.html#file-metadata
.. _pelicanconf.py: http://docs.getpelican.com/en/stable/settings.html
