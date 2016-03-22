[Python] Convert HTML Table to reStructuredText list-table
##########################################################

:date: 2016-02-28T20:42+08:00
:tags: Python, Web Scrape, reStructuredText, remove trailing newline, html, DOM,
       String Manipulation, Beautiful Soup
:category: Python
:summary: Python_ script to convert `HTML table`_ to reStructuredText_
          `list-table`_ via `Beautiful Soup 4`_.


Python_ script to convert `HTML table`_ to reStructuredText_ `list-table`_ via
`Beautiful Soup 4`_ (bs4).

.. code-block:: python

  #!/usr/bin/env python
  # -*- coding:utf-8 -*-

  import os
  import urllib
  from bs4 import BeautifulSoup

  def string2rst(string):
    result = ''
    count = 0
    for string2 in string.strip().split("\r"):
      if count == 0:
        result += string2.strip()
      else:
        result += ("\n       " + string2.strip())
      count += 1
    return result


  def processTd(td, fo, count):
    if count == 0:
      count2 = 0
      for string in td.stripped_strings:
        if count2 == 0:
          fo.write("   * - " + string2rst(string).encode("utf-8") + "\n")
        else:
          fo.write("\n       " + string2rst(string).encode("utf-8") + "\n")
        count2 += 1
    else:
      count2 = 0
      for string in td.stripped_strings:
        if count2 == 0:
          fo.write("     - " + string2rst(string).encode("utf-8") + "\n")
        else:
          fo.write("\n       " + string2rst(string).encode("utf-8") + "\n")
        count2 += 1
    fo.write("\n")


  def htmlTable2Rst(filepath, url):
    num = url[-6:-4]
    dstpath = "../content/articles/tipitaka/sutta/khuddaka/dhammapada/dhp-chap{}%zh.rst".format(num)
    tmppath = "/tmp/tmp.rst"
    print(dstpath)
    with open(filepath, 'r') as f:
      document = BeautifulSoup(f)
      table = document.find("table")
      title = document.find("h1")

      with open(tmppath, 'w') as fo:
        table_title = ""
        rst_title = ""
        for index, string in enumerate(title.stripped_strings):
          table_title += string
          if index == 0:
            rst_title = "Dhammapada 法句經({})".format(string.encode("utf-8"))

        fo.write(rst_title)
        fo.write("\n")
        fo.write("=" * len(rst_title))
        fo.write("\n\n")
        fo.write(":date: 2004-04-03\n")
        fo.write(":modified: 2004-04-03\n")
        fo.write(":tags: 法句經, Dhammapada, The Buddha's Path of Wisdom, The Path of Dhamma, The Word of the Doctrine, 法集要頌經, 法句譬喻經, 出曜經, DHP, Dhp\n")
        fo.write(":category: 巴利三藏小部\n")
        fo.write(":summary: 法句經, Dhammapada, The Buddha's Path of Wisdom, The Path of Dhamma, The Word of the Doctrine\n")
        fo.write("\n")
        fo.write("\n")
        fo.write(".. list-table:: " + table_title.encode("utf-8") + "\n")
        fo.write("   :header-rows: 1\n")
        fo.write("   :class: contrast-reading-table\n\n")

        trs = table.find_all("tr")

        ct = 0
        for th in trs[0].find_all("th"):
          processTd(th, fo, ct)
          ct += 1

        for tr in trs[1:]:
          tds = tr.find_all("td")
          count = 0
          for td in tds:
            processTd(td, fo, count)
            count += 1

        fo.write('備註：英譯可參考 "佛學數位圖書館暨博物館"中 巴利語教學 `經文選讀 (英) <http://buddhism.lib.ntu.edu.tw/DLMBS/lesson/pali/lesson_pali3.jsp>`_\n')
        fo.write("\n----\n\n參考：\n\n.. [a] ")
        fo.write("`舊網頁 <http://nanda.online-dhamma.net/Tipitaka/Sutta/Khuddaka/Dhammapada/DhP_Chap{}.htm>`_".format(num))

    with open(tmppath, "r") as f:
      with open(dstpath, 'w') as fo:
        fo.write(f.read().replace("偈\n\n       頌\n\n       次", "偈\n       頌\n       次").replace("(\n\n       典故\n\n       )", "(\n       典故\n       )"))


  if __name__ == '__main__':
    for i in range(5, 27):
      url = "http://nanda.online-dhamma.net/Tipitaka/Sutta/Khuddaka/Dhammapada/DhP_Chap{0:02d}.htm".format(i)
      filepath = os.path.join("/tmp", os.path.basename(url))
      urllib.urlretrieve(url, filepath)
      htmlTable2Rst(filepath, url)
    """
    if os.path.exists(filepath):
      htmlTable2Rst(filepath, url)
    else:
      urllib.urlretrieve(url, filepath)
    """

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `舊文移植：Dhammapada 法句經(第一：雙品) · twnanda/twnanda@65e4d01 · GitHub <https://github.com/twnanda/twnanda/commit/65e4d017b682a17db79fed60d734b391aeaeb1f1>`_

.. [2] `BeautifulSoup 4 Documentation <http://www.crummy.com/software/BeautifulSoup/bs4/doc/>`__

.. [3] `twnanda/olddhp.py at master · twnanda/twnanda · GitHub <https://github.com/twnanda/twnanda/blob/master/tool/olddhp.py>`_

.. [4] `小信' Blog <http://playbear.github.io/>`_

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
.. _list-table: http://docutils.sourceforge.net/docs/ref/rst/directives.html#list-table
.. _bs4: http://www.crummy.com/software/BeautifulSoup/bs4/doc/
.. _Beautiful Soup 4: http://www.crummy.com/software/BeautifulSoup/bs4/doc/
.. _HTML table: http://www.w3schools.com/html/html_tables.asp
