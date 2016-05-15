Basic HTML Table to reStructuredText Support for HTML2RST service on GAE Go
###########################################################################

:date: 2016-05-15T22:10+08:00
:tags: Go, Golang, Golang template, Go net/http, Web Scrape, reStructuredText,
       Copy to Clipboard, Go net/html, Google App Engine, DOM
:category: Google App Engine
:summary: Add basic `HTML table`_ to reStructuredText_ support for
          `online HTML to reStructuredText service`_ on `Google App Engine Go`_.


Add basic `HTML table`_ to reStructuredText_ support for
`online HTML to reStructuredText service`_ on `Google App Engine Go`_.
Based on my previous work [1]_ and [2]_.

.. rubric:: `Online HTML to reStructuredText <http://html2rst.golden-operator-130720.appspot.com/>`_
   :class: align-center

For `HTML table`_ to reStructuredText_ implementation, see `this commit`_.
For complete source code, source code, see GitHub html2rst_ repo.

----

Tested on: ``Ubuntu Linux 16.04``, ``Google App Engine SDK for Go 1.9.37``.

----

References:

.. [1] `[Golang] Unrobust HTML Table to reStructuredText list-table <{filename}../../04/13/go-unrobust-html-table-to-rst-list-table%en.rst>`_

.. [2] `Online HTML to reStructuredText on Google App Engine Go <{filename}../13/gae-go-online-html-to-rst%en.rst>`_

.. [3] `[Golang] Convert File Encoding From Big5 to UTF-8 <{filename}../../03/21/go-convert-file-encoding-from-big5-to-utf8%en.rst>`_

.. [4] `[Golang] HTML to reStructuredText <{filename}../12/go-html-to-rst%en.rst>`_


.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _HTML: https://www.google.com/search?q=HTML
.. _Google App Engine Go: https://cloud.google.com/appengine/docs/go/
.. _gae directory: https://github.com/siongui/html2rst/tree/master/gae
.. _html2rst: https://github.com/siongui/html2rst
.. _HTML table: http://www.w3schools.com/html/html_tables.asp
.. _online HTML to reStructuredText service: http://html2rst.golden-operator-130720.appspot.com/
.. _this commit: https://github.com/siongui/html2rst/commit/bb3ef68014c7b9bd145386f3af85ad7a945911cc
