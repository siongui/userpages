[Golang] Auto-Detect and Convert Encoding of HTML to UTF-8
##########################################################

:date: 2018-10-27T04:05+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n, Web Scrape, iconv command,
       Go net/html, Go net/http
:category: Go
:summary: Given an URL, auto-detect and convert the encoding of the HTML
          document to UTF-8 if it is not UTF-8 encoded in *Go*.
:og_image: http://1.bp.blogspot.com/-8J4Yz4LWYPs/U_yr6-loLnI/AAAAAAAABzg/HBGWbD6A7Vo/s1600/Character%2BEncoding%2C%2BConverting%2BByte%2Barray%2Bto%2BString%2Bin%2BJava.png
:adsu: yes


Given an URL, auto-detect and convert the encoding of the HTML document to UTF-8
if it is not UTF-8 encoded in Go_. Package `golang.org/x/net/html`_ is used to
determine the encoding of the HTML document, and package `golang.org/x/text`_ is
used for encoding conversion.

Install the packages first:

.. code-block:: bash

  $ go get -u golang.org/x/text
  $ go get -u golang.org/x/net/html

The following code shows how to determine the encoding of an HTML document given
the URL, and convert the encoding to utf8 if it is not utf8 encoded:

.. show_github_file:: siongui userpages content/code/go/utf8-encoding/url.go

Note that *resp.Body* stream cannot be read twice, so it is read into bytes
first and used later for detection and conversion.

Usage of the above code:

.. show_github_file:: siongui userpages content/code/go/utf8-encoding/url_test.go

----

Tested on: ``Ubuntu 18.04``, ``Go 1.11.1``

----

.. adsu:: 2

References:

.. [1] | `爬蟲自動轉碼，獲取城市列表：城市名稱+URL - 掃文資訊 <https://tw.saowen.com/a/a57e654a2d2b091f5582a429948b101693a3540c765128998a52316597b7a9a5>`_
       | `爬虫自动转码，获取城市列表：城市名称+URL  - Go语言中文网 - Golang中文社区 <https://studygolang.com/articles/14418>`_
       | `爬虫自动转码，获取城市列表：城市名称+URL - 简书 <https://www.jianshu.com/p/16981659bd8d>`_
.. [2] `golang 用/x/net/html写的小爬虫，爬小说 - 简书 <https://www.jianshu.com/p/9b31ecb0d5ab>`_
.. [3] `[Golang] Determine Encoding of HTML Document <{filename}/articles/2018/10/26/determine-encoding-of-html-document-in-go%en.rst>`_

.. _Go: https://golang.org/
.. _golang.org/x/net/html: https://godoc.org/golang.org/x/net/html
.. _golang.org/x/text: https://godoc.org/golang.org/x/text
