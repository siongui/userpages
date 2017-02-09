[Golang] Convert PO file to JSON Format
#######################################

:date: 2016-01-27T06:36+08:00
:tags: Go, Golang, i18n, Locale, JSON, Regular Expression, gettext
:category: Go
:summary: Convert PO_ files to JSON_ format via Go_. The data of JSON format can
          be passed to front-end by web servers to translate a text string into
          the user's native language. You can use the JSON data to implement
          `gettext function`_ in browsers.
:adsu: yes


Introduction
++++++++++++

Write a Go_ program to convert PO_ files to JSON_ format. The data of JSON
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

.. show_github_file:: siongui userpages content/code/go-gettext/locale/zh_TW/LC_MESSAGES/messages.po
.. adsu:: 2

vi_VN PO file ``locale/vi_VN/LC_MESSAGES/messages.po``:

.. show_github_file:: siongui userpages content/code/go-gettext/locale/vi_VN/LC_MESSAGES/messages.po
.. adsu:: 3

Source Code
+++++++++++

Convert PO_ files to JSON_ format:

.. show_github_file:: siongui userpages content/code/go-gettext/po2json.go
.. adsu:: 4
.. show_github_file:: siongui userpages content/code/go-gettext/po2json_test.go
.. adsu:: 5


Output of Demo
++++++++++++++

.. code-block:: txt

  === RUN   TestAll
  --- PASS: TestAll (0.00s)
          po2json_test.go:10: locale/zh_TW/LC_MESSAGES/messages.po
          po2json_test.go:11: locale/vi_VN/LC_MESSAGES/messages.po
          po2json_test.go:13: {"vi_VN":{"About":"Giới thiệu","Canon":"Kinh điển","Home":"Trang chính","Setting":"Thiết lập","Translation":"Dịch"},"zh_TW":{"About":"關於","Canon":"經典","Home":"首頁","Setting":"設定","Translation":"翻譯"}}
  PASS
  ok      command-line-arguments  0.003s


Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `golang regular expression <https://www.google.com/search?q=golang+regular+expression>`_

.. [2] `regexp - The Go Programming Language <https://golang.org/pkg/regexp/#Regexp.FindAllStringSubmatch>`_

.. [3] `golang read file to string <https://www.google.com/search?q=golang+read+file+to+string>`_

.. [4] `go - How Can i read a whole file into a string variable in golang? - Stack Overflow <http://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable-in-golang>`_

.. [5] `json - The Go Programming Language <https://golang.org/pkg/encoding/json/#example_Marshal>`_


.. _gettext: https://www.gnu.org/software/gettext/
.. _locale: https://en.wikipedia.org/wiki/Locale
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _PO: https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html
.. _JSON: https://www.google.com/search?q=JSON
.. _gettext function: http://linux.die.net/man/3/gettext
