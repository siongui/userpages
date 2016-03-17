[Python] Use Regular Expression to Find Strings Marked For Internationalization (i18n)
######################################################################################

:date: 2016-01-01T19:12+08:00
:tags: Python, Regular Expression, String Manipulation, Locale, i18n, gettext,
       find command
:category: Python
:summary: Use regular expression in Python to search and find strings marked for
          i18n (Internationalization)

Python Regular Expression
+++++++++++++++++++++++++

An i18n (web) application usually mark strings to be translated as
``_("string")``. You can use xgettext_ in `GNU gettext utilities`_ to extract
translatable strings from given input files. This post, however, will use
`regular expression`_ in Python to do the work.

The basic pattern to search ``_("string")`` is:

.. code-block:: python

  def searchI18n(string):
    # only first match and longest match
    # i.e., the string {{_("ddd")}}12345{{_("sss")}} will return
    # {{_("ddd")}}12345{{_("sss")}}, not return {{_("ddd")}}
    return re.search(r'{{\s*_\(\s*(.+)\s*\)\s*}}', string)

A more advanced pattern is:

.. code-block:: python

  def getAllMatchesInFile(filepath):
    with open(filepath, 'r') as f:
      # [^)] to prevent {{_("ddd")}}12345{{_("sss")}}
      return re.findall(r'{{\s*_\(\s*([^)]+)\s*\)\s*}}', f.read())

The above function will return all matched strings in a file.


Alternative (Use xgettext_)
+++++++++++++++++++++++++++

You can also use the following command line in Linux console to extract strings:
(assume your strings are in HTML files in ``.`` directory)

.. code-block:: bash

  xgettext --no-wrap --from-code=UTF-8 --keyword=_ --output=messages.pot `find . -name *.html`

xgettext_ will save the strings in the file named *messages.pot*.

----

References:

.. [1] `Python Regular Expressions  |  Google for Education  |  Google Developers <https://developers.google.com/edu/python/regular-expressions>`_

.. [2] `Regex replace (in Python) - a simpler way? - Stack Overflow <http://stackoverflow.com/questions/490597/regex-replace-in-python-a-simpler-way>`_

.. [3] `python - Import a module from a relative path - Stack Overflow <http://stackoverflow.com/questions/279237/import-a-module-from-a-relative-path>`_

.. [4] `Internationalize a Python application - maemo.org wiki <http://wiki.maemo.org/Internationalize_a_Python_application>`_

.. [5] `Python localization made easy «  Supernifty – nifty stuff <http://www.supernifty.org/blog/2011/09/16/python-localization-made-easy/>`_

.. [6] `GNU gettext utilities <http://www.gnu.org/software/gettext/manual/gettext.html>`_


.. _xgettext: https://www.gnu.org/software/gettext/manual/html_node/xgettext-Invocation.html
.. _regular expression: https://www.google.com/search?q=Regular+Expression
.. _GNU gettext utilities: http://www.gnu.org/software/gettext/manual/gettext.html
