xgettext Extract Translatable Strings From Golang html/template
###############################################################

:date: 2016-01-19T21:09+08:00
:tags: Go, Golang, i18n, Locale, gettext, html, Web application, web, sed
:category: Go
:summary: Use xgettext_ (one of GNU gettext_ utilities) to extract translatable
          strings from HTML_ templates of Go_ `html/template`_.


GNU gettext_ tools are great for creating i18n_ (web) applications. xgettext_ is
one of GNU gettext_ utilities, which extracts translatable strings from given
input files, and outputs PO_ template (*POT* file).

If HTML_ templates are written by Go_ `html/template`_ (read [3]_), you will
find that xgettext_ cannot extracts translatable strings directly from the
templates, because Go_ `html/template`_ uses different syntax.

We can write a custom *gettext* function and use this function in templates by
Go_ `html/template`_ (read [3]_), but the *gettext* function must be called with
the following syntax:

.. code-block:: html

  {{gettext "Home"}}

But xgettext_ can not extract above strings. The idiomatic syntax used for
xgettext_ to extract strings are:

.. code-block:: html

  {{gettext("Home")}}

So how to extract the translatable strings in syntax of Go_ `html/template`_?
After some googling [4]_, I found [5]_ provides me the idea of solution. The
solution is to uses sed_ to convert the Go_ `html/template`_ syntax to idiomatic
syntax by the following command:

.. code-block:: bash

  sed "s/{{gettext \(".*"\)}}/{{gettext(\1)}}/g" html.go | xgettext --no-wrap --language=c --from-code=UTF-8 --output=locale/messages.pot -

After conversion of syntax, the output are feeded to xgettext_ to extract
strings. Handlebars_ templates uses syntax silimiar to Go_ `html/template`_,
so the solution here is also applied to Handlebars_ templates.


Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `Internationalization (i18n) of Web Application by GNU gettext Tools <{filename}../07/i18n-web-application-by-gnu-gettext-tools%en.rst>`_

.. [2] `[Golang] Internationalization (i18n) of Go Application by GNU gettext Tools <{filename}../08/golang-i18n-go-application-by-gnu-gettext%en.rst>`_

.. [3] `i18n Golang Web Application by gettext and html/template <{filename}i18n-go-web-application-by-gettext-html-template%en.rst>`_

.. [4] `xgettext example <https://www.google.com/search?q=xgettext+example>`_

.. [5] `php - Let xgettext find keywords in comments - Stack Overflow <http://stackoverflow.com/questions/7645319/let-xgettext-find-keywords-in-comments>`_

.. [6] `gmarty/xgettext · GitHub <https://github.com/gmarty/xgettext>`_ (Extract translatable strings from Handlebars templates.)

.. [7] `arendjr/grunt-xgettext: Grunt xgettext plugin for JavaScript and Handlebars <https://github.com/arendjr/grunt-xgettext>`_


.. _HTML: http://www.w3schools.com/html/
.. _gettext: https://www.gnu.org/software/gettext/
.. _i18n: https://en.wikipedia.org/wiki/Internationalization_and_localization
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _html/template: https://golang.org/pkg/html/template/
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _xgettext: https://www.gnu.org/software/gettext/manual/html_node/xgettext-Invocation.html
.. _sed: http://www.grymoire.com/Unix/Sed.html
.. _Handlebars: http://handlebarsjs.com/
