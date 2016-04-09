[Python] Convert PO file to JSON Format
#######################################

:date: 2016-01-26T20:12+08:00
:tags: Python, i18n, Locale, JSON, Regular Expression, gettext,
       File Input/Output
:category: Python
:summary: Convert PO_ files to JSON_ format via Python_. The data of JSON format
          can be passed to front-end by web servers to translate a text string
          into the user's native language. You can use the JSON data to
          implement `gettext function`_ in browsers.


Introduction
++++++++++++

Write a Python_ program to convert PO_ files to JSON_ format. The data of JSON
format can be passed to front-end by web servers to translate a text string into
the user's native language. You can use the JSON data from PO_ files to
implement `gettext function`_ in browsers.


Sample PO_ files
++++++++++++++++

In this example, we support two locale_, *zh_TW (Traditional Chinese)* and
*vi_VN (Vietnamese)*. The zh_TW PO file are located at
``locale/zh_TW/LC_MESSAGES/messages.po`` and vi_VN PO file are located at
``locale/vi_VN/LC_MESSAGES/messages.po``.

zh_TW PO file ``locale/zh_TW/LC_MESSAGES/messages.po``:

.. show_github_file:: siongui userpages content/code/python-i18n-gettext/locale/zh_TW/LC_MESSAGES/messages.po

vi_VN PO file ``locale/vi_VN/LC_MESSAGES/messages.po``:

.. show_github_file:: siongui userpages content/code/python-i18n-gettext/locale/vi_VN/LC_MESSAGES/messages.po

Source Code
+++++++++++

Convert PO_ files to JSON_ format:

.. show_github_file:: siongui userpages content/code/python-i18n-gettext/po2json.py


Output of Demo
++++++++++++++

.. code-block:: txt

  {"zh_TW": {"Home": "\u9996\u9801", "About": "\u95dc\u65bc", "Setting": "\u8a2d\u5b9a", "Canon": "\u7d93\u5178", "Translation": "\u7ffb\u8b6f"}, "vi_VN": {"Home": "Trang ch\u00ednh", "About": "Gi\u1edbi thi\u1ec7u", "Setting": "Thi\u1ebft l\u1eadp", "Canon": "Kinh \u0111i\u1ec3n", "Translation": "D\u1ecbch"}}


Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `Python Regular Expressions  |  Google for Education  |  Google Developers <https://developers.google.com/edu/python/regular-expressions>`_

.. [2] `Regex replace (in Python) - a simpler way? - Stack Overflow <http://stackoverflow.com/questions/490597/regex-replace-in-python-a-simpler-way>`_


.. _gettext: https://www.gnu.org/software/gettext/
.. _locale: https://en.wikipedia.org/wiki/Locale
.. _Python: https://www.python.org/
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _JSON: https://www.google.com/search?q=JSON
.. _gettext function: http://linux.die.net/man/3/gettext
