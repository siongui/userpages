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


  def htmlTable2Rst(filepath):
    with open(filepath, 'r') as f:
      document = BeautifulSoup(f)
      table = document.find("table")
      title = document.find("h1")
      with open("tmp.rst", 'w') as fo:
        table_title = ""
        for string in title.stripped_strings:
          table_title += string
          fo.write(string.encode("utf-8") + "\n")
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


  if __name__ == '__main__':
    url = "http://nanda.online-dhamma.net/Tipitaka/Sutta/Khuddaka/Dhammapada/DhP_Chap01.htm"
    filename = os.path.basename(url)
    if os.path.exists(filename):
      htmlTable2Rst(filename)
    else:
      urllib.urlretrieve(url, filename)


----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `舊文移植：Dhammapada 法句經(第一：雙品) · twnanda/twnanda@65e4d01 · GitHub <https://github.com/twnanda/twnanda/commit/65e4d017b682a17db79fed60d734b391aeaeb1f1>`_

.. [2] `BeautifulSoup 4 Documentation <http://www.crummy.com/software/BeautifulSoup/bs4/doc/>`__

.. _Python: https://www.python.org/
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _list-table: http://docutils.sourceforge.net/docs/ref/rst/directives.html#list-table
.. _bs4: http://www.crummy.com/software/BeautifulSoup/bs4/doc/
.. _Beautiful Soup 4: http://www.crummy.com/software/BeautifulSoup/bs4/doc/
.. _HTML table: http://www.w3schools.com/html/html_tables.asp
