Google App Engine Go - HTML Link to reStructuredText
####################################################

:date: 2016-05-11T22:31+08:00
:tags: Go, Golang, Golang template, Go net/http, Web Scrape, reStructuredText,
       Copy to Clipboard, Go net/html, Google App Engine
:category: Google App Engine
:summary: Given a webpage URL. Fetch the title of the webpage and output
          `reStructuredText link`_ on `Google App Engine Go`_.


Given a webpage URL. Fetch the title of the webpage and output
`reStructuredText link`_ on `Google App Engine Go`_.

.. rubric:: `Demo <http://v1.golden-operator-130720.appspot.com/>`_
   :class: align-center

*Makefile*: Edit the variable in Makefile according to your development
environment.

.. show_github_file:: siongui userpages content/code/go-gae-link2rst/Makefile

*GAE Go Environment Setting*:

.. show_github_file:: siongui userpages content/code/go-gae-link2rst/app.yaml

Source code:

.. show_github_file:: siongui userpages content/code/go-gae-link2rst/link2rst.go

.. show_github_file:: siongui userpages content/code/go-gae-link2rst/title.go

.. show_github_file:: siongui userpages content/code/go-gae-link2rst/fetch.go

----

Tested on: ``Ubuntu Linux 16.04``, ``Google App Engine SDK for Go 1.9.37``.

----

References:

.. [1] `[Golang] Server Get Form POST Value <{filename}../../03/27/go-server-get-form-post-value%en.rst>`_

.. [2] `[Golang] Get HTML Title via net/html Package <{filename}..//10/go-get-html-title-via-net-html%en.rst>`_

.. [3] `[Golang] Hacker News Link to reStructuredText <{filename}../../04/04/go-hacker-news-link-to-rst%en.rst>`_

.. [4] `Issuing HTTP(S) Requests - Go — Google Cloud Platform <https://cloud.google.com/appengine/docs/go/issue-requests>`_


.. _reStructuredText link: http://docutils.sourceforge.net/docs/user/rst/quickref.html#hyperlink-targets
.. _Google App Engine Go: https://cloud.google.com/appengine/docs/go/
