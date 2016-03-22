[Python] Export PIXNET Blog to reStructuredText Files
#####################################################

:date: 2016-02-17T07:30+08:00
:tags: Python, Web Scrape, reStructuredText, remove trailing newline, html, DOM,
       String Manipulation, Beautiful Soup, List Files in Directory
:category: Python
:summary: Python_ scripts to export PIXNET_ blog posts to files of
          reStructuredText_ format.


Python_ scripts [1]_ to export PIXNET_ blog posts to reStructuredText_ (rst)
files. bs4_ (Beautiful Soup 4) are used to do DOM manipulation of HTML, such as
find elements in DOM or access the content of a HTML element.

Usage
+++++

.. show_github_file:: siongui pixnet2rst demo.py

Source Code
+++++++++++

Get urls of all blog posts of a user account:

.. show_github_file:: siongui pixnet2rst allPostsUrls.py

Get HTML and convert the metadata and content of posts to rst format:

.. show_github_file:: siongui pixnet2rst html2rst.py

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `GitHub - siongui/pixnet2rst: Export your PIXNET blog posts to reStructuredText (rst) files. <https://github.com/siongui/pixnet2rst>`_

.. [2] `BeautifulSoup 4 Documentation <http://www.crummy.com/software/BeautifulSoup/bs4/doc/>`__

.. [3] `小信' Blog <http://playbear.github.io/>`_

.. [5] `小趴趴--知乎精华回答的非专业大数据统计 <http://www.jianshu.com/p/6d53b34165d2>`_
       (`伯樂在線轉錄 <http://python.jobbole.com/84524/>`__,
       `GitHub - SmileXie/zhihu_crawler <https://github.com/SmileXie/zhihu_crawler>`__)

.. [6] `使用python进行web抓取 -  磁针石的个人空间 - 开源中国社区 <http://my.oschina.net/u/1433482/blog/620858>`_
       (`伯樂在線轉錄 <http://python.jobbole.com/84523/>`__)

.. [7] `关于背单词软件,你不知道的惊人真相 <http://www.jianshu.com/p/b57e55cb5941>`_
       (`伯樂在線轉錄 <http://python.jobbole.com/84526/>`__,
       `GitHub <https://github.com/twocucao/DataScience/>`__)

.. _Python: https://www.python.org/
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _PIXNET: https://www.pixnet.net/
.. _bs4: http://www.crummy.com/software/BeautifulSoup/bs4/doc/
