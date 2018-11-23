[Golang] Login Facebook and Take Screenshot Programmatically
############################################################

:date: 2017-05-03T04:41+08:00
:tags: Go, Golang, Web Scrape, Scraping JavaScript Web Page
:category: Go
:summary: Web scraping JavaScript rendered webpages by Chrome Debugging Protocol
          . Write Go code to programmatically login Facebook and then take
          screenshot.
:og_image: http://www.cuelogic.com/blog/wp-content/uploads/2014/11/GO.jpg
:adsu: yes

I heard the news of headless Chrome from HN [1]_, and found that it is possible
to write a Go program [3]_ to automatically web-scraping JavaScript rendered web
pages via *Chrome Debugging Protocol* [2]_. I have tried to use Python dryscrape
to programmatically login a website before [4]_, and now I will try to use Go
*chromedp* package to login Facebook and take screenshot programmatically.

Install Go *chromedp* package for programmatically drive the browser and
simulate human user actions:

.. code-block:: bash

  $ go get -u github.com/knq/chromedp

Install Go package for read password from console [5]_:

.. code-block:: bash

  $ go get -u golang.org/x/crypto/ssh/terminal

**Login Facebook and Take Screenshot**:

.. show_github_file:: siongui userpages content/code/go/fb-login-screenshot/main.go
.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.04``, ``Go 1.8.1``.

----

References:

.. [1] `Getting Started with Headless Chrome | Hacker News <https://news.ycombinator.com/item?id=14239194>`_
.. [2] | `chrome debugging protocol golang - Google search <https://www.google.com/search?q=chrome+debugging+protocol+golang>`_
       | `chrome debugging protocol golang - DuckDuckGo search <https://duckduckgo.com/?q=chrome+debugging+protocol+golang>`_
       | `chrome debugging protocol golang - Ecosia search <https://www.ecosia.org/search?q=chrome+debugging+protocol+golang>`_
       | `chrome debugging protocol golang - Qwant search <https://www.qwant.com/?q=chrome+debugging+protocol+golang>`_
       | `chrome debugging protocol golang - Bing search <https://www.bing.com/search?q=chrome+debugging+protocol+golang>`_
       | `chrome debugging protocol golang - Yahoo search <https://search.yahoo.com/search?p=chrome+debugging+protocol+golang>`_
       | `chrome debugging protocol golang - Baidu search <https://www.baidu.com/s?wd=chrome+debugging+protocol+golang>`_
       | `chrome debugging protocol golang - Yandex search <https://www.yandex.com/search/?text=chrome+debugging+protocol+golang>`_
.. [3] `GitHub - knq/chromedp: Package chromedp is a faster, simpler way to drive browsers (Chrome, Edge, Safari, Android, etc) without external dependencies (ie, Selenium, PhantomJS, etc) using the Chrome Debugging Protocol. <https://github.com/knq/chromedp>`_
.. [4] `[Python] Web Scrape JavaScript Webpage by dryscrape <{filename}../../../2016/01/06/python-web-scrape-javascript-webpage-by-dryscrape%en.rst>`_
.. [5] | `[Golang] Read Yes/No From Console <{filename}../../../2016/04/23/go-read-yes-no-from-console%en.rst>`_
       | `golang read password - Google search <https://www.google.com/search?q=golang+read+password>`_
.. [6] | `初探 Headless Chrome - WEB前端 - 伯乐在线 <http://web.jobbole.com/91489/>`_
       | `初探 Headless Chrome - 知乎专栏 <https://zhuanlan.zhihu.com/p/27100187>`_
.. [7] `What's the recommended encryption library for Go? : golang <https://old.reddit.com/r/golang/comments/9zhqga/whats_the_recommended_encryption_library_for_go/>`_

.. _Web scrape: https://en.wikipedia.org/wiki/Web_scraping
.. _Python: https://www.python.org/
.. _dryscrape: https://github.com/niklasb/dryscrape
.. _Requests: http://docs.python-requests.org/
.. _web scraping: https://en.wikipedia.org/wiki/Web_scraping
.. _iframe: http://www.w3schools.com/tags/tag_iframe.asp
