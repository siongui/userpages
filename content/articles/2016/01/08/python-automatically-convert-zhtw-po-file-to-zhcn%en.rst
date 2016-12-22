[Python] Automatically Convert Traditional Chinese PO file to Simplified Chinese
################################################################################

:date: 2016-01-08T00:22+08:00
:tags: Python, String Manipulation, i18n, Locale, gettext, File Input/Output,
       Conversion of Traditional and Simplified Chinese, OpenCC
:category: Python
:summary: Automatically convert `Traditional Chinese`_ (`zh_TW`_) PO_ file to
          `Simplified Chinese`_ (`zh_CN`_) by OpenCC_ and Python_ programming
          language.

In this post, we will write a Python script to automatically convert
`Traditional Chinese`_ (`zh_TW`_) PO_ file to `Simplified Chinese`_ (`zh_CN`_)
by OpenCC_ (Open Chinese Convert) and pyOpenCC_ (OpenCC Python binding). Please
read my previous post [1]_ to install OpenCC_ and pyOpenCC_ first.


Source Code
+++++++++++

The `zh_TW`_ PO_ file for test:

.. show_github_file:: siongui userpages content/code/python-zhtw-to-zhcn-po-file/locale/zh_TW/LC_MESSAGES/messages.po

The Python script:

.. show_github_file:: siongui userpages content/code/python-zhtw-to-zhcn-po-file/tw2cn.py


Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``, ``opencc 0.4.3-2build1``,
``pyopencc-0.4.2.2``.

----

References:

.. [1] `[Python] Conversion of Traditional and Simplified Chinese <{filename}../04/python-conversion-of-traditional-and-simplified-chinese%en.rst>`_

.. [2] `Python Regular Expressions  |  Google for Education  |  Google Developers <https://developers.google.com/edu/python/regular-expressions>`_

.. [3] `Regex replace (in Python) - a simpler way? - Stack Overflow <http://stackoverflow.com/questions/490597/regex-replace-in-python-a-simpler-way>`_

.. [4] `python - Import a module from a relative path - Stack Overflow <http://stackoverflow.com/questions/279237/import-a-module-from-a-relative-path>`_


.. _Python: https://www.python.org/
.. _pyOpenCC: https://github.com/cute/pyopencc
.. _OpenCC: http://opencc.byvoid.com/
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _Traditional Chinese: https://en.wikipedia.org/wiki/Traditional_Chinese_characters
.. _Simplified Chinese: https://en.wikipedia.org/wiki/Simplified_Chinese_characters
.. _zh_TW: https://docs.oracle.com/cd/E19455-01/806-0169/6j9hsml3g/index.html
.. _zh_CN: https://docs.oracle.com/cd/E19683-01/806-6642/new-tbl-72/index.html
