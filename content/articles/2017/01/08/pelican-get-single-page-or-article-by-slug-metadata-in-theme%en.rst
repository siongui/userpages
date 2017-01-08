[Pelican] Get Single Page or Article by slug Metadata in Theme
##############################################################

:date: 2017-01-08T21:49+08:00
:tags: Python, Pelican, Jinja2, i18n
:category: Python
:summary: Pelican_ static site generator - Get a single page_ or article_ by
          *slug* metadata_ in theme_.


Sometimes we want to get a single page_ or article_ by some metadata_ in theme_.
This post shows how to use a `custom Jinja2 filter`_ to get a single page_ or
article_ by *slug* metadata_ in the theme_.

First, add the following filter code in your `pelicanconf.py`_:

.. code-block:: python

  # custom Jinja2 filter
  def get_by_slug(objs, slug):
      for obj in objs:
          if obj.slug == slug:
              return obj

  JINJA_FILTERS = {
      "get_by_slug": get_by_slug,
  }


Then you can use the custom filter to select the page or article you need in the
theme_. For example, you can use as follows in ``index.html``:

.. code-block:: html

  {% set intro = (pages|get_by_slug("introduction")) %}
  {% set hello = (articles|get_by_slug("hello-world")) %}

  <div>Title of page: {{ intro.title }}</div>
  <div>Title of article: {{ hello.title }}</div>

----

Tested on: ``Ubuntu Linux 16.10``, ``Python 2.7.12+``, `Pelican 3.7.0`_.

----

References:

.. [1] `new arch: FIXME: URL issue · siongui/shenfang@33b9932 · GitHub <https://github.com/siongui/shenfang/commit/33b993216b41b86f2083f6c4cf7b23ae47ba858b>`_


.. _Pelican: http://blog.getpelican.com/
.. _Pelican 3.7.0: http://docs.getpelican.com/en/3.7.0/
.. _page: http://docs.getpelican.com/en/latest/content.html#articles-and-pages
.. _article: http://docs.getpelican.com/en/latest/content.html#articles-and-pages
.. _metadata: http://docs.getpelican.com/en/latest/content.html#file-metadata
.. _theme: http://docs.getpelican.com/en/latest/themes.html
.. _Jinja2: https://www.google.com/search?q=jinja2
.. _custom Jinja2 filter: http://jinja.pocoo.org/docs/latest/api/#custom-filters
.. _pelicanconf.py: http://docs.getpelican.com/en/latest/settings.html
