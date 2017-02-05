[Golang] Internationalization (i18n) of Go Application by GNU gettext Tools
###########################################################################

:date: 2016-01-08T04:49+08:00
:tags: Go, Golang, i18n, Locale, gettext
:category: Go
:summary: Internationalization (i18n_) of Go_ applications by GNU gettext_
          tools. This post shows how to use PO_ and MO_ files in Golang_.
:adsu: yes


GNU gettext_ tools are great for creating i18n (web) applications. In this post,
assume that PO_ and MO_ files are ready and we will use the PO_ and MO_ files to
let Go_ applications speak local languages. If you do not know what PO_ or MO_
files are, or you do not know how to create them by gettext_ tools, please read
my previous post [1]_ first. If you want to know how to render HTML by
`html/template`_ with the translated texts in PO_ and MO_ files in your Go_ web
application, please refer to [12]_.


PO_ and MO_ files
+++++++++++++++++

I will show how to use MO_ and PO_ files by example. In the example, we support
two locale_, *zh_TW (Traditional Chinese)* and *vi_VN (Vietnamese)*. The zh_TW
PO file are located at ``locale/zh_TW/LC_MESSAGES/messages.po`` and vi_VN PO
file are located at ``locale/vi_VN/LC_MESSAGES/messages.po``.

.. adsu:: 2

zh_TW PO file ``locale/zh_TW/LC_MESSAGES/messages.po``:

.. show_github_file:: siongui userpages content/code/go-gettext/locale/zh_TW/LC_MESSAGES/messages.po

vi_VN PO file ``locale/vi_VN/LC_MESSAGES/messages.po``:

.. show_github_file:: siongui userpages content/code/go-gettext/locale/vi_VN/LC_MESSAGES/messages.po

Generate corresponding MO files by msgfmt_. Put MO files together with PO files.
So we have two PO files and two MO files:

.. code-block:: txt

  locale/zh_TW/LC_MESSAGES/messages.po
  locale/zh_TW/LC_MESSAGES/messages.mo
  locale/vi_VN/LC_MESSAGES/messages.po
  locale/vi_VN/LC_MESSAGES/messages.mo


Source Code
+++++++++++

Next, we need the `gettext-go`_ package to handle the PO_ and MO_ files for
us. Install `gettext-go`_ by:

.. code-block:: bash

  $ go get github.com/chai2010/gettext-go/gettext

Now we can let Go_ application speak local language:

.. show_github_file:: siongui userpages content/code/go-gettext/gettext.go
.. adsu:: 3

.. note::

  In line 17 of above code, we use ``gettext.PGettext("", input)`` instead of
  ``gettext.Gettext(input)`` because no *msgctxt* in our PO_ file (only *msgid*
  and *msgstr*), so specify context as ``""`` in **PGettext**. If you use
  **Gettext** in this case, original string will be returned. For furthur
  detail, see `issue #1`_ of `gettext-go`_ on GitHub.

.. show_github_file:: siongui userpages content/code/go-gettext/gettext_test.go

.. note::

  The *domain* in this case is **messages**. If you name your PO_ and MO_ file
  as **hello**, the *domain* becomes **hello**.

  .. code-block:: txt

    locale/zh_TW/LC_MESSAGES/hello.po
    locale/zh_TW/LC_MESSAGES/hello.mo
    locale/vi_VN/LC_MESSAGES/hello.po
    locale/vi_VN/LC_MESSAGES/hello.mo


Output of Above Code
++++++++++++++++++++

.. code-block:: txt

  === RUN   TestAll
  --- PASS: TestAll (0.00s)
          gettext_test.go:10: 首頁
          gettext_test.go:11: 經典
          gettext_test.go:12: 關於
          gettext_test.go:13: 設定
          gettext_test.go:14: 翻譯
          gettext_test.go:17: Trang chính
          gettext_test.go:18: Kinh điển
          gettext_test.go:19: Giới thiệu
          gettext_test.go:20: Thiết lập
          gettext_test.go:21: Dịch
  PASS


Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.2``.

----

References:

.. [1] `Internationalization (i18n) of Web Application by GNU gettext Tools <{filename}../07/i18n-web-application-by-gnu-gettext-tools%en.rst>`_

.. [2] Google Search `go gettext <https://www.google.com/search?q=go+gettext>`__

.. [3] `chai2010/gettext-go · GitHub <https://github.com/chai2010/gettext-go>`_
       |godoc1-png|

.. [4] `Go语言的国际化支持(资源文件翻译) - CHAI2010 <http://chai2010.github.io/blog/2014/01/27/gettext-go-intro-02/>`_

.. [5] `Go语言的国际化支持(基于gettext-go) - CHAI2010 <http://chai2010.github.io/blog/2014/01/07/gettext-go-intro/>`_

.. [6] `localization - I18n strategies for Go with App Engine - Stack Overflow <http://stackoverflow.com/questions/14124630/i18n-strategies-for-go-with-app-engine>`_

.. [7] `samuel/go-gettext · GitHub <https://github.com/samuel/go-gettext>`_
       |godoc2-png|

.. [8] Google Search `go i18n <https://www.google.com/search?q=go+i18n>`__

.. [9] `[Python] Internationalization (i18n) of Python Application by GNU gettext Tools <{filename}../14/python-i18n-py-application-by-gnu-gettext%en.rst>`_

.. [10] `i18n Python Web Application by gettext and Jinja2 <{filename}../17/i18n-python-web-application-by-gettext-jinja2%en.rst>`_

.. [11] `gosexy/gettext · GitHub <https://github.com/gosexy/gettext>`_

.. [12] `i18n Golang Web Application by gettext and html/template <{filename}../19/i18n-go-web-application-by-gettext-html-template%en.rst>`_


.. _gettext: https://www.gnu.org/software/gettext/
.. _i18n: https://en.wikipedia.org/wiki/Internationalization_and_localization
.. _locale: https://en.wikipedia.org/wiki/Locale
.. _Python: https://www.python.org/
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _MO: https://www.gnu.org/software/gettext/manual/html_node/MO-Files.html
.. _msgfmt: https://www.gnu.org/software/gettext/manual/html_node/msgfmt-Invocation.html
.. _gettext-go: https://github.com/chai2010/gettext-go
.. _issue #1: https://github.com/chai2010/gettext-go/issues/1
.. _html/template: https://golang.org/pkg/html/template/

.. |godoc1-svg| image:: https://godoc.org/github.com/chai2010/gettext-go/gettext?status.svg
   :target: https://godoc.org/github.com/chai2010/gettext-go/gettext

.. |godoc1-png| image:: https://godoc.org/github.com/chai2010/gettext-go/gettext?status.png
   :target: https://godoc.org/github.com/chai2010/gettext-go/gettext

.. |godoc2-svg| image:: https://godoc.org/github.com/samuel/go-gettext/gettext?status.svg
   :target: https://godoc.org/github.com/samuel/go-gettext/gettext

.. |godoc2-png| image:: https://godoc.org/github.com/samuel/go-gettext/gettext?status.png
   :target: https://godoc.org/github.com/samuel/go-gettext/gettext

.. python - Image grid in reStructuredText / Sphinx - Stack Overflow
   http://stackoverflow.com/questions/10219634/image-grid-in-restructuredtext-sphinx
   Google Search: rst image in table
