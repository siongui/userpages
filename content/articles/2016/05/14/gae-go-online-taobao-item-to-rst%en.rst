Online Taobao Item to reStructuredText Image on Google App Engine Go
####################################################################

:date: 2016-05-14T23:01+08:00
:tags: Go, Golang, Golang template, Go net/http, Web Scrape, reStructuredText,
       Copy to Clipboard, Go net/html, Google App Engine, DOM, Go net/url
:category: Google App Engine
:summary: Online service on `Google App Engine Go`_, which helps you extract
          title, image URL from Taobao_ item webpage, and output in
          reStructuredText_  format.


Online service on `Google App Engine Go`_, which helps you extract title, image
URL from Taobao_ item webpage, and output in reStructuredText_  format.

.. rubric:: `Online Taobao Item to reStructuredText <http://taobao-item2rst.golden-operator-130720.appspot.com/>`_
   :class: align-center

Source code:

.. show_github_file:: siongui userpages content/code/go-gae-taobao-item2rst/Makefile

.. show_github_file:: siongui userpages content/code/go-gae-taobao-item2rst/app.yaml

.. show_github_file:: siongui userpages content/code/go-gae-taobao-item2rst/taobaoitem2rst.go

.. show_github_file:: siongui userpages content/code/go-gae-taobao-item2rst/fetch.go

.. show_github_file:: siongui userpages content/code/go-gae-taobao-item2rst/urlnormalize.go

.. show_github_file:: siongui userpages content/code/go-gae-taobao-item2rst/iteminfo.go

----

Tested on: ``Ubuntu Linux 16.04``, ``Google App Engine SDK for Go 1.9.37``.

----

References:

.. [1] `Google App Engine Go - HTML Link to reStructuredText <{filename}../11/gae-go-html-link-to-rst%en.rst>`_

.. [2] `[Golang] Remove Query String From URL <{filename}../../03/26/go-remove-querystring-from-url%en.rst>`_

.. [3] `[Golang] Hacker News Link to reStructuredText <{filename}../../04/04/go-hacker-news-link-to-rst%en.rst>`_

.. [4] `[Golang] getElementById via net/html Package <{filename}../../04/15/go-getElementById-via-net-html-package%en.rst>`_

.. [5] `go http no redirect - Google search <https://www.google.com/search?q=go+http+no+redirect>`_

       `Query URL without redirect in Go - Stack Overflow <http://stackoverflow.com/questions/14420222/query-url-without-redirect-in-go>`_

.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _Google App Engine Go: https://cloud.google.com/appengine/docs/go/
.. _Taobao: https://www.taobao.com/
