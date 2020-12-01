[Pelican] Add Build Timestamp to Pelican Site Automatically
###########################################################

:date: 2020-11-30T01:52+08:00
:tags: Pelican, Python
:category: Python
:summary: Add build timestamp to Pelican site automatically during generating
          the site.
:og_image: https://miro.medium.com/max/878/1*MbgGn92JzfkPgz-mPm0ufw.png
:adsu: yes

When we generate the website using Python Pelican static site generator, we may
want to add "Last Updated" time to our website at the footer or somewhere else
automatically. This post will show you how to do it.

In ``pelicanconf.py``, add following code:

.. code-block:: python

  from datetime import datetime
  import pytz

  # modify TIMEZONE to your timezone
  TIMEZONE = 'Asia/Taipei'
  BUILD_TIME = datetime.now(pytz.timezone(TIMEZONE))

In the template html, you can use the *BUILD_TIME* as follows:

.. code-block:: html

  Last Updated: {{ BUILD_TIME | strftime("%b %d, %Y, %H:%M:%S") }} ({{ TIMEZONE }})

If you do not know about Python strftime, check
`this post <https://www.programiz.com/python-programming/datetime/strftime>`_.

Tested on: ``Ubuntu Linux 20.04``, ``Python 3.8.5``.

----

References:

.. [1] | `pelican jinja2 current date - Google search <https://www.google.com/search?q=pelican+jinja2+current+date>`_
       | `pelican jinja2 current date - DuckDuckGo search <https://duckduckgo.com/?q=pelican+jinja2+current+date>`_
       | `pelican jinja2 current date - Ecosia search <https://www.ecosia.org/search?q=pelican+jinja2+current+date>`_
       | `pelican jinja2 current date - Qwant search <https://www.qwant.com/?q=pelican+jinja2+current+date>`_
       | `pelican jinja2 current date - Bing search <https://www.bing.com/search?q=pelican+jinja2+current+date>`_
       | `pelican jinja2 current date - Yahoo search <https://search.yahoo.com/search?p=pelican+jinja2+current+date>`_
       | `pelican jinja2 current date - Baidu search <https://www.baidu.com/s?wd=pelican+jinja2+current+date>`_
       | `pelican jinja2 current date - Yandex search <https://www.yandex.com/search/?text=pelican+jinja2+current+date>`_

.. [2] `How to keep your copyright year up-to-date • bernhard.scheirle.de <https://bernhard.scheirle.de/posts/2016/February/29/how-to-keep-your-copyright-year-up-to-date-using-jinja-filters/>`_

.. [3] `python - Print the last modification of a jinja2 template in pelican - Stack Overflow <https://stackoverflow.com/questions/20766692/print-the-last-modification-of-a-jinja2-template-in-pelican>`_

.. [4] `How do I get a value of datetime.today() in Python that is "timezone aware"? - Stack Overflow <https://stackoverflow.com/a/16660476>`_

.. _Pelican: https://github.com/getpelican/pelican
