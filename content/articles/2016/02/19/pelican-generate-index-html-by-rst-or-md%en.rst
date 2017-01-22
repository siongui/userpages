Pelican Generate index.html by reStructuredText or Markdown
###########################################################

:date: 2016-02-19T06:52+08:00
:tags: Python, reStructuredText, Pelican, Jinja2
:category: Python
:summary: Pelican_ static site generator - Generate index.html_ by writing
          reStructuredText_ or Markdown_.
:adsu: yes


Use Pelican_ static site generator to generate index.html_ by writing
reStructuredText_ or Markdown_.

In Pelican_, you can `create your own theme`_. In this post, we will show you
how to wrtie rst_ or md_ to generate index.html_ in your customized theme.


Create `hidden pages`_
++++++++++++++++++++++

First we create `hidden pages`_ by rst or md, in which the content will be our
index.html. We create the following **index.rst** (or **index.md** if you use
Markdown):

.. code-block:: rst

  :slug: index
  :lang: en
  :status: hidden
  :summary: English index.html
  :og_image: YOUR og:image URL

  .. write your English index.html content here

If you use i18n_subsites_ plugin, you can create hidden pages with different
*lang* metadata for each supported language. For example, you can create
**index_zh.rst** for Chinese language with the following content:

.. code-block:: rst

  :slug: index
  :lang: zh
  :status: hidden
  :summary: Chinese index.html
  :og_image: YOUR og:image URL

  .. write your Chinese index.html content here

`Jinja2 filter`_ in pelicanconf.py_
+++++++++++++++++++++++++++++++++++

We create a custom `Jinja2 filter`_ to select the hidden pages with *slug* is
*index*.

.. code-block:: python

  # custom Jinja2 filter
  def hidden_pages_get_page_with_slug_index(hidden_pages):
      for page in hidden_pages:
          if page.slug == "index":
              return page

  JINJA_FILTERS = {
      "hidden_pages_get_page_with_slug_index": hidden_pages_get_page_with_slug_index,
  }


index.html_ in theme
++++++++++++++++++++

Then we modify the index.html in theme to do something like following:

.. code-block:: html

  {% extends "layout/layout.html" %}
  {% set index = (hidden_pages|hidden_pages_get_page_with_slug_index) %}

  <div class="title">{{ index.title }}</div>
  <div class="summary">{{ index.summary }}</div>
  <div class="main-content">{{ index.content }}</div>

Just the same as you are writing **page.html** or **article.html** in the
customized theme.

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``, `Pelican 3.6.3`_.

----

References:

.. [1] `pelican custom Jinja2 filter to let you write index.html in rst · siongui/pelican-template@e70b7ca · GitHub <https://github.com/siongui/pelican-template/commit/e70b7ca15937f54f174196e5096211dd75a8d2ac>`_

.. [2] JINJA_FILTERS in `Settings — Pelican documentation <http://docs.getpelican.com/en/latest/settings.html>`_

       `Jinja custom filters documentation <http://jinja.pocoo.org/docs/dev/api/#custom-filters>`_

.. [3] `GitHub - siongui/pelican-template <https://github.com/siongui/pelican-template>`_

.. [4] `python - Querying for specific articles (via tag/category) in Pelican themes - Stack Overflow <http://stackoverflow.com/questions/19283880/querying-for-specific-articles-via-tag-category-in-pelican-themes>`_


.. _Python: https://www.python.org/
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _rst: https://www.google.com/search?q=reStructuredText
.. _Markdown: https://www.google.com/search?q=Markdown
.. _md: https://www.google.com/search?q=Markdown
.. _Pelican: http://blog.getpelican.com/
.. _Pelican 3.6.3: http://docs.getpelican.com/en/3.6.3/
.. _i18n_subsites: https://github.com/getpelican/pelican-plugins/tree/master/i18n_subsites
.. _index.html: https://www.google.com/search?q=index.html
.. _create your own theme: http://docs.getpelican.com/en/latest/themes.html
.. _pelicanconf.py: http://docs.getpelican.com/en/latest/settings.html
.. _hidden pages: http://docs.getpelican.com/en/latest/themes.html
.. _Jinja2 filter: http://jinja.pocoo.org/docs/dev/api/#custom-filters
