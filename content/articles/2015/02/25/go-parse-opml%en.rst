[Golang] XML Parsing Example (5) - Parse OPML
#############################################

:date: 2015-02-25 15:05
:tags: Go, Golang, XML, html, OPML
:category: Go
:summary: How to read XML/HTML files in Go programming language (for newbie)
          - Parse OPML format.


Based on previous examples, this post shows how to parse OPML_ file. The
following OPML example comes from my `web feeds`_:

.. show_github_file:: siongui userpages content/code/go-xml/example-5.xml

Source code for parsing above OPML:

.. show_github_file:: siongui userpages content/code/go-xml/parse-5.go


The output result:

.. code-block:: txt

  {Hacker News Hacker News rss https://news.ycombinator.com/rss https://news.ycombinator.com/ https://news.ycombinator.com/favicon.ico}
  {CSDN最新资讯 CSDN最新资讯 rss http://www.csdn.net/article/rss_lastnews http://news.csdn.net http://s2.googleusercontent.com/s2/favicons?domain=csdn.net}
  {Solidot Solidot rss http://www.solidot.org/index.rss http://www.solidot.org http://s2.googleusercontent.com/s2/favicons?domain=solidot.org}
  {Slashdot Slashdot rss http://rss.slashdot.org/Slashdot/slashdot http://slashdot.org/ http://slashdot.org/favicon.ico}
  {InfoQ中文 InfoQ中文 rss http://www.infoq.com/cn/rss/rss.action?token=Apz6T7JrRXKONbiLoKzdoHiJ8uTgu5ef http://www.infoq.com/cn/ http://infoqstatic.com/favicon.ico}
  {外刊IT评论 外刊IT评论 rss http://www.aqee.net/feed/ http://www.aqee.net http://s2.googleusercontent.com/s2/favicons?domain=aqee.net}
  {博客 - 伯乐在线 博客 - 伯乐在线 rss http://blog.jobbole.com/feed/ http://blog.jobbole.com http://cdn2.jobbole.com/2013/10/favicon.png}


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}../21/go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}../24/go-parse-xml-example-4%en.rst>`_

.. [5] `[Golang] XML Parsing Example (5) - Parse OPML <{filename}go-parse-opml%en.rst>`_

.. [6] `[Golang] XML Parsing Example (6) - Parse OPML Concisely <{filename}../26/go-parse-opml-concisely%en.rst>`_

----

Reference:

.. [a] `OPML <http://en.wikipedia.org/wiki/OPML>`_

.. _OPML: http://en.wikipedia.org/wiki/OPML

.. _web feeds: http://en.wikipedia.org/wiki/Web_feed
