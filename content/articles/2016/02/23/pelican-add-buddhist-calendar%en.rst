Pelican Add Buddhist Calendar
#############################

:date: 2016-02-23T14:28+08:00
:tags: Python, Pelican, Jinja2
:category: Python
:summary: Pelican_ static site generator - add `Buddhist calendar`_
          in your theme_.
:adsu: yes


Use Pelican_ static site generator to add `Buddhist calendar`_ in your theme_.

Usually the datetime format (`Gregorian calendar`_) in the theme looks like:

.. code-block:: html

  {{ article.date | strftime('%B %d, %Y') }}

Now we will change `Gregorian calendar`_ to `Buddhist calendar`_.
First define a Jinja2_ macro_:

.. code-block:: html

  {% macro Buddhist_calendar(datetime) -%}
  {{ datetime | strftime('%B %d') }}, {{ datetime.year + 543 }}
  {%- endmacro %}

Use the macro in your theme by:

.. code-block:: html

  {{ Buddhist_calendar(article.date) }}

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``, `Pelican 3.6.3`_.

----

References:

.. [1] `add Buddhist calendar · siongui/pelican-template@4616e29 · GitHub <https://github.com/siongui/pelican-template/commit/4616e2945507cd8bab1658ac9e21acdb5120de4d>`_



.. _Python: https://www.python.org/
.. _Pelican: http://blog.getpelican.com/
.. _Pelican 3.6.3: http://docs.getpelican.com/en/3.6.3/
.. _theme: http://docs.getpelican.com/en/latest/themes.html
.. _Buddhist calendar: https://en.wikipedia.org/wiki/Buddhist_calendar
.. _Gregorian calendar: https://en.wikipedia.org/wiki/Gregorian_calendar
.. _Jinja2: http://jinja.pocoo.org/docs/dev/
.. _macro: http://jinja.pocoo.org/docs/dev/templates/#macros
