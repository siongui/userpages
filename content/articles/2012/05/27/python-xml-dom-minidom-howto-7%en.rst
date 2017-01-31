Python Library xml.dom.minidom Howto (7)
########################################

:modified: 2017-02-01T02:57+08:00
:tags: DOM, html, minidom, Python, XML
:category: Python
:summary: Python_ XML_/HTML_ manipulation primer of xml.dom.minidom_
:adsu: yes


PARSE XML/HTML FROM A FILE
==========================

This post gives a real-world example about how to parse and retrieve data from
a XML_/HTML_ file by the use of Python_ xml.dom.minidom_ library. The following
is a XML_ file which contains the explanation of a `Pāli`_ word *abbhāna*.
We want to parse the file and extract the information.

.. show_github_file:: siongui userpages content/articles/2012/05/27/example.xml
.. adsu:: 2

The following Python_ script parses the above XML_ file. In line 21, the script
parses the XML_ file first. In line 23, we get the item element by calling
*getElementsByTagName*. Then we parse each item one by one. Extract the content
of the text node in line 11, 12, 13. The result of each item is printed in line
15, 16, 17. The code is straight forward and easy to understand.

.. show_github_file:: siongui userpages content/articles/2012/05/27/minidom-howto-7.py
.. adsu:: 3

The result of the above Python_ script is:

.. code-block:: bash

  dict: ◎　《汉译パーリ语辞典》 黃秉榮譯 词数 7735.
  word: abbhāna
  explain: %3a%6e%2e%20%5b%61%62%68%69%2d%c4%81%79%c4%81%6e%61%5d%20%e5%87%ba%e7%bd%aa%2c%20%e5%ae%b9%e8%a8%b1%2c%20%e5%be%a9%e6%ad%b8%28%e6%81%a2%e5%be%a9%e5%8e%9f%e7%8b%80%29%2e
  dict: ◎　《パーリ语辞典》 日本水野弘元教授 词数 13772.
  word: abbhāna
  explain: %3a%6e%2e%20%5b%61%62%68%69%2d%c4%81%79%c4%81%6e%61%5d%20%e5%87%ba%e7%bd%aa%2c%20%e8%a8%b1%e5%ae%b9%2c%20%e5%be%a9%e5%b8%b0%2e
  dict: ◎　《巴汉词典》 明法尊者增订
  word: Abbhāna
  explain: %2c%20%28%61%62%68%69%20%2b%20%c4%81%79%61%6e%61%20%6f%66%20%c4%81%20%2b%20%79%c4%81%20%28%69%29%29%2c%e3%80%90%e4%b8%ad%e3%80%91%e5%a4%8d%e5%bd%92%28%e6%af%94%e4%b8%98%e8%ba%ab%e4%bb%bd%29%28%63%6f%6d%69%6e%67%20%62%61%63%6b%2c%20%72%65%68%61%62%69%6c%69%74%61%74%69%6f%6e%20%6f%66%20%61%20%62%68%69%6b%6b%68%75%20%77%68%6f%20%68%61%73%20%75%6e%64%65%72%67%6f%6e%65%20%61%20%70%65%6e%61%6e%63%65%20%66%6f%72%20%61%6e%20%65%78%70%69%61%62%6c%65%20%6f%66%66%65%6e%63%65%29%e3%80%82
  dict: ◎　《PTS Pali-English dictionary》 The Pali Text Society's Pali-English dictionary
  word: Abbhāna
  explain: %2c%28%6e%74%2e%29%20%5b%61%62%68%69%20%2b%20%c4%81%79%61%6e%61%20%6f%66%20%c4%81%20%2b%20%3c%65%6d%3e%79%c4%81%3c%2f%65%6d%3e%3c%69%3e%20%28%3c%2f%69%3e%3c%65%6d%3e%69%3c%2f%65%6d%3e%3c%69%3e%29%3c%2f%69%3e%5d%20%63%6f%6d%69%6e%67%20%62%61%63%6b%2c%20%72%65%68%61%62%69%6c%69%74%61%74%69%6f%6e%20%6f%66%20%61%20%62%68%69%6b%6b%68%75%20%77%68%6f%20%68%61%73%20%75%6e%64%65%72%67%6f%6e%65%20%61%20%70%65%6e%61%6e%63%65%20%66%6f%72%20%61%6e%20%65%78%70%69%61%62%6c%65%20%6f%66%66%65%6e%63%65%20%56%69%6e%2e%49%2c%34%39%20%28%c2%b0%c3%a2%72%61%68%61%29%2c%20%35%33%20%28%69%64%2e%29%2c%20%31%34%33%2c%20%33%32%37%3b%20%49%49%2c%33%33%2c%20%34%30%2c%20%31%36%32%3b%20%41%2e%49%2c%39%39%2e%20%2d%2d%20%43%70%2e%20%3c%69%3e%61%62%62%68%65%74%69%3c%2f%69%3e%2e%20%28%50%61%67%65%20%36%30%29


----

*Python Library xml.dom.minidom Howto* series:

.. [1] `Python Library xml.dom.minidom Howto (1) <{filename}../24/python-xml-dom-minidom-howto-1%en.rst>`_

.. [2] `Python Library xml.dom.minidom Howto (2) <{filename}../24/python-xml-dom-minidom-howto-2%en.rst>`_

.. [3] `Python Library xml.dom.minidom Howto (3) <{filename}../24/python-xml-dom-minidom-howto-3%en.rst>`_

.. [4] `Python Library xml.dom.minidom Howto (4) <{filename}../24/python-xml-dom-minidom-howto-4%en.rst>`_

.. [5] `Python Library xml.dom.minidom Howto (5) <{filename}../24/python-xml-dom-minidom-howto-5%en.rst>`_

.. [6] `Python Library xml.dom.minidom Howto (6) <{filename}../24/python-xml-dom-minidom-howto-6%en.rst>`_

.. [7] `Python Library xml.dom.minidom Howto (7) <{filename}python-xml-dom-minidom-howto-7%en.rst>`_

----

Reference: `MiniDom - Python Wiki <https://wiki.python.org/moin/MiniDom>`_

.. _Python: https://www.python.org/
.. _XML: https://www.google.com/search?q=XML
.. _HTML: https://www.google.com/search?q=HTML
.. _DOM: https://www.google.com/search?q=DOM
.. _xml.dom.minidom: https://www.google.com/search?q=xml.dom.minidom
.. _Pāli: https://en.wikipedia.org/wiki/Pali
