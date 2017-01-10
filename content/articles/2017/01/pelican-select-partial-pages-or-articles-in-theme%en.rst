[Pelican] Get Partials Pages or Articles in Theme
#################################################

:date: 2017-01-10T22:04+08:00
:tags: Python, Pelican, Jinja2
:category: Python
:summary: Pelican_ static site generator - Get partial pages_ or articles_ with
          specific metadata_ by Jinja2_ built-in selectattr_ filter  in theme_.


Inspired by the answer to re-order pages in [1]_, this post shows how to get
partial pages_ or articles_ with specific metadata_ by Jinja2_ built-in
selectattr_ in the theme_.

Illustrate by example, assume we want to get pages of which ``order`` is ``2``
in ``index.html`` in theme_:

(metadata_ ``order`` is ``2`` in reStructuredText_):

.. code-block:: rst

  :order: 2

(metadata_ ``order`` is ``2`` in Markdown_):

.. code-block:: rst

  Order: 2

You can select by Jinja2_ built-in selectattr_ filter in `index.html`_ in
theme_ as follows:

.. code-block:: html

  {% for page in pages|selectattr("order", "equalto", "2") %}
    <div>Title: {{ page.title }}</div>
  {% endfor %}

----

Tested on: ``Ubuntu Linux 16.10``, ``Python 2.7.12+``, `Pelican 3.7.0`_.

----

References:

.. [1] `pelican page order - Google search <https://www.google.com/search?q=pelican+page+order>`_

       `Ordering content - Settings — Pelican latest documentation <http://docs.getpelican.com/en/latest/settings.html#ordering-content>`_

       `python - How can I control the order of pages from within a pelican article category? - Stack Overflow <http://stackoverflow.com/questions/18520046/how-can-i-control-the-order-of-pages-from-within-a-pelican-article-category>`_

.. [2] `TemplateRuntimeError: no test named - Google search <https://www.google.com/search?q=TemplateRuntimeError:+no+test+named>`_

       `python - How do I filter a collection in a jinja2 template ? - Stack Overflow <http://stackoverflow.com/questions/34974691/how-do-i-filter-a-collection-in-a-jinja2-template>`_


.. _Pelican: http://blog.getpelican.com/
.. _page: http://docs.getpelican.com/en/latest/content.html#articles-and-pages
.. _pages: http://docs.getpelican.com/en/latest/content.html#articles-and-pages
.. _article: http://docs.getpelican.com/en/latest/content.html#articles-and-pages
.. _articles: http://docs.getpelican.com/en/latest/content.html#articles-and-pages
.. _metadata: http://docs.getpelican.com/en/latest/content.html#file-metadata
.. _theme: http://docs.getpelican.com/en/latest/themes.html
.. _Jinja2: https://www.google.com/search?q=jinja2
.. _selectattr: http://jinja.pocoo.org/docs/2.9/templates/#selectattr
.. _custom Jinja2 filter: http://jinja.pocoo.org/docs/latest/api/#custom-filters
.. _pelicanconf.py: http://docs.getpelican.com/en/latest/settings.html
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _Markdown: https://www.google.com/search?q=Markdown
.. _index.html: http://docs.getpelican.com/en/latest/themes.html#index-html
.. _Pelican 3.7.0: http://docs.getpelican.com/en/3.7.0/
