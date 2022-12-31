[Pelican] Remove Prefix in PATH_METADATA Using Named Regex Group
################################################################

:date: 2023-01-01T01:57+07:00
:tags: Python, Pelican, Regular Expression, String Manipulation
:category: Python
:summary: Remove prefix of Pelican PATH_METADATA using Python named regular
          expression group.
:og_image: https://i.ytimg.com/vi/WcT6MSLLE-M/maxresdefault.jpg
:adsu: yes


**Question**

In Pelican static website, we have the following pages under *content*
directory.

  1. pages/index%zh-hant.rst
  2. pages/about-us%zh-hant.rst
  3. pages/talk/thai-forest-tradition%zh-hant.rst
  4. pages/talk/thanissaro/how-to-fall%zh-hant.rst

We want to extract 3 metadata: *urlpath*, *slug*, and *lang* from the path of
the pages such that

  1. | *urlpath*: **empty string**
     | *slug*: **index**
     | *lang*: **zh-hant**
  2. | *urlpath*: empty string
     | *slug*: **about-us**
     | *lang*: **zh-hant**
  3. | *urlpath*: **talk/**
     | *slug*: **thai-forest-tradition**
     | *lang*: **zh-hant**
  4. | *urlpath*: **talk/thanissaro/**
     | *slug*: **how-to-fall**
     | *lang*: **zh-hant**

As you can see, the prefix **pages/** are removed from *urlpath*

----

**Solution**

In *pelicanconf.py*

.. code-block:: python

   PATH_METADATA = 'pages/(?P<urlpath>[-a-zA-Z0-9/]*/|)(?P<slug>[-a-zA-Z0-9]*)%(?P<lang>[-_a-zA-Z]{2,7})\.rst'
   PAGE_URL = '{urlpath}{slug}/'
   PAGE_SAVE_AS = '{urlpath}{slug}/index.html'

----

Tested on: ``Ubuntu Linux 22.04``, ``Python 3.10.6``.

----

References:

.. [1] `Settings - Pelican 4.8.0 <https://docs.getpelican.com/en/latest/settings.html#metadata>`_
.. [2] `Regex: match an empty string instead of nothing - Stack Overflow <https://stackoverflow.com/a/57893349>`_

.. _Pelican: https://getpelican.com/
.. _PATH_METADATA: https://docs.getpelican.com/en/stable/settings.html#metadata
