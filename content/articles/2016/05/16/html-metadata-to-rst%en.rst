Guess Metadata from HTML and Converted to reStructuredText
##########################################################

:date: 2016-05-16T22:40+08:00
:tags: Go, Golang, Golang template, Go net/http, Web Scrape, reStructuredText,
       Go net/html, Go text/width, DOM
:category: Go
:summary: Guess metadata from HTML_ webpage and convert it to reStructuredText_
          format.


Guess metadata from HTML_ webpage and convert it to reStructuredText_ format.
Currently the following metadata extraction (if available) is supported:

- title
- keywords (tags)
- description (summary)
- author
- og:image

Usage:

.. show_github_file:: siongui html2rst metadata_test.go

Check `guess metadata from HTML`_ commit in html2rst_ repo for details of source
code.

----

Tested on: ``Ubuntu Linux 16.04``, ``Go 1.6.2``.

----

References:

.. [1] `[Golang] HTML to reStructuredText <{filename}../12/go-html-to-rst%en.rst>`_

.. [2] `Online Taobao Item to reStructuredText Image on Google App Engine Go <{filename}../14/gae-go-online-taobao-item-to-rst%en.rst>`_

.. [3] `[Golang] Create reStructuredText Metadata via text/template Package <{filename}../../04/22/go-rst-metadata-via-text-template%en.rst>`_

.. [4] `[Golang] Extract Title, Image, and URL via goquery <{filename}../../03/31/go-parse-buy123-webpage-to-rst%en.rst>`_


.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _HTML: https://www.google.com/search?q=HTML
.. _html2rst: https://github.com/siongui/html2rst
.. _guess metadata from HTML: https://github.com/siongui/html2rst/commit/167287af21e99504edb00a766aa4f4e74e1cfa18
