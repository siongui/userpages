[Python] Web Scrape JavaScript Webpage by dryscrape
###################################################

:date: 2016-01-06T02:03+08:00
:tags: Python, Web Scrape, JavaScript
:category: Python
:summary: `Web scrape`_ JavaScript rendered webpages by dryscrape_, a
          lightweight web scraping library for Python_.

Today JavaScript is heavily used to render the website content. Requests_, a
Python_ HTTP library, is not enough for `web scraping`_. In this post we will
try to use dryscrape_, a lightweight web scraping library for Python, to scrape
dynamically rendered webpages by JavaScript.

Install dryscrape_
++++++++++++++++++

.. code-block:: bash

  $ sudo apt-get install qt5-default libqt5webkit5-dev build-essential python-lxml python-pip xvfb
  $ sudo pip install dryscrape

Real World Example
++++++++++++++++++

We will write a Python script to visit a webpage with iframe_. Get the URL of
the iframe. Fill in the form in the iframe_ and submit the form.

.. show_github_file:: siongui userpages content/code/python-dryscrape-form/submit.py



Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``, ``dryscrape 1.0``.

----

References:

.. [1] Google Search: `python login script <https://www.google.com/search?q=python+login+script>`_

.. [2] Google Search: `requests python javascript <https://www.google.com/search?q=requests+python+javascript>`_

.. [3] `Web-scraping JavaScript page with Python - Stack Overflow <http://stackoverflow.com/questions/8049520/web-scraping-javascript-page-with-python>`_

.. [4] `How to submit a javascript-form using Python requests library? - Stack Overflow <http://stackoverflow.com/questions/20802108/how-to-submit-a-javascript-form-using-python-requests-library>`_

.. [5] `Ultimate guide for scraping  JavaScript rendered web pages | IMPYTHONIST <https://impythonist.wordpress.com/2015/01/06/ultimate-guide-for-scraping-javascript-rendered-web-pages/>`_

.. [6] `niklasb/dryscrape · GitHub <https://github.com/niklasb/dryscrape>`_
       (A lightweight Python library that uses Webkit to enable easy scraping of dynamic, Javascript-heavy web pages)

.. [7] Google Search: `python scrape javascript <https://www.google.com/search?q=python+scrape+javascript>`_

.. [8] `Scraping with JavaScript | Web Scraping with Python <http://pythonscraping.com/blog/javascript>`_

.. [9] `Selenium - Web Browser Automation <http://seleniumhq.org/>`_
       (`GitHub repo <https://github.com/SeleniumHQ/selenium/>`__)

.. [10] `Selenium with Python — Selenium Python Bindings 2 documentation <http://selenium-python.readthedocs.org/>`_

.. [11] `How to use Selenium with Python? - Stack Overflow <http://stackoverflow.com/questions/17540971/how-to-use-selenium-with-python>`_

.. [12] Google Search: `Selenium with Python <https://www.google.com/search?q=Selenium+with+Python>`_

.. [13] Google Search: `Selenium Python <https://www.google.com/search?q=Selenium+Python>`_

.. [14] `Web Scraping: Beyond BeautifulSoup : Python - Reddit <https://www.reddit.com/r/Python/comments/1xj39b/web_scraping_beyond_beautifulsoup/>`_

.. [15] `ghost.py: webkit web client written in python <http://jeanphix.me/Ghost.py/>`_
        (`GitHub repo <https://github.com/jeanphix/Ghost.py>`__)



.. _Web scrape: https://en.wikipedia.org/wiki/Web_scraping
.. _Python: https://www.python.org/
.. _dryscrape: https://github.com/niklasb/dryscrape
.. _Requests: http://docs.python-requests.org/
.. _web scraping: https://en.wikipedia.org/wiki/Web_scraping
.. _iframe: http://www.w3schools.com/tags/tag_iframe.asp