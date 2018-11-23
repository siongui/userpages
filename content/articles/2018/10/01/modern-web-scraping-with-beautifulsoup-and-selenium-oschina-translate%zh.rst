使用 BeautifulSoup 和 Selenium 進行網頁爬取
###########################################

:date: 2018-09-21T18:33+08:00
:tags: Python, 轉錄, 開源中國, Linux
:category: Python
:author: 開源中國翻譯
:summary: 在本教程中，您將瞭解在瀏覽器中看到的內容是如何實際呈現的，以及如何在必要時進行抓取。特別是，您將學習如何計算Disqus評論。我們的工具是Python和這門語言的很棒的包，比如request、BeautifulSoup和Selenium。
:og_image: https://oscimg.oschina.net/oscnet/879d4b68295e6c808b4a1417e1abdf062a1.jpg

.. highlights::

  | 英文原文： `Modern Web Scraping With BeautifulSoup and Selenium`_
  | 轉錄來源： `使用 BeautifulSoup 和 Selenium 进行网页爬取 - 开源中国`_
  | 參與翻譯 (3人) : `ZICK_ZEON`_, xiaoaiwhc1_, `边城`_

----

概述
++++

HTML幾乎是平鋪直敘的。CSS是一個偉大的進步，它清晰地區分了頁面的結構和外觀。JavaScript添加一些魅力。道理上講是這樣的。現實世界還是有點不一樣。

在本教程中，您將瞭解在瀏覽器中看到的內容是如何實際呈現的，以及如何在必要時進行抓取。特別是，您將學習如何計算Disqus評論。我們的工具是Python和這門語言的很棒的包，比如request、BeautifulSoup和Selenium。


什麼時候應該使用網頁爬取?
+++++++++++++++++++++++++

網頁爬取是一種自動獲取被設計於實現人工用戶交互式網頁的內容、解析它們並提取一些信息(可能是導航到其他頁面的鏈接)的實踐。如果沒有其他方法來提取必要的網頁信息時，網頁爬取是很必要有效的技術方法。理想情況下，應用程序依靠提供好的專用API來編程自動獲得網頁的數據。可在下面幾種場所景之下你最好就別用網頁抓取技術了:

- 被爬取的網頁是脆弱的(您正在爬取的網頁可能會被頻繁更改)。
- 爬取被禁止(一些web應用程序有禁止爬取的策略)。
- 爬取速度可能會很慢和爬取內容過於繁雜的(如果你需要在很多無用信息中尋找和涉獵你想要的東東)。

----

瞭解真實的網頁
++++++++++++++

讓我們通過查看一些常見web應用程序代碼的實現情況，來瞭解我們面臨的問題。例如在“ `Vagrant技術入門`_ ”這篇帖子的頁面底部有一些Disqus的評論:

.. image:: https://oscimg.oschina.net/oscnet/879d4b68295e6c808b4a1417e1abdf062a1.jpg
   :alt: 使用 BeautifulSoup 和 Selenium 進行網頁爬取
   :align: center

為了爬取這些評論，我們需要首先在頁面上找到它們。


查看頁面代碼
++++++++++++

自20世紀90年代以來，每個瀏覽器都支持查看當前頁面的HTML代碼。下面是在源碼視圖下觀看到的是“ `Vagrant技術入門`_ ”這篇帖子對應的源碼內容的一個片段，這篇源碼以大量與本文本身內容無關的被壓縮過的和醜陋的JavaScript代碼開始。下面是其中的一”小“部分:

.. image:: https://oscimg.oschina.net/oscnet/f90d94b626069bb9009fc6f3c62f5ce7000.jpg
   :alt: 使用 BeautifulSoup 和 Selenium 進行網頁爬取
   :align: center

這是頁面中的一些實際HTML代碼:

.. image:: https://oscimg.oschina.net/oscnet/444c92812cde124c1d3414d4381f855e2e9.jpg
   :alt: 使用 BeautifulSoup 和 Selenium 進行網頁爬取
   :align: center

代碼看起來亂糟糟，你竟然在頁面的源代碼中找不到Disqus評論，這讓你有些吃驚。

----

強大的內聯框架
++++++++++++++

原來頁面是一個”混搭“, Disqus評論被嵌入到iframe(內聯框架)元素中。你可以通過右鍵點擊評論區域找到它，你會看到那裡有框架信息和源碼:

.. image:: https://oscimg.oschina.net/oscnet/3bd0f23073a40cb98a2948c448f981479ce.jpg
   :alt: 使用 BeautifulSoup 和 Selenium 進行網頁爬取
   :align: center

這是有意義的。將第三方內容嵌入iframe是使用iframe的主要應用場景之一。讓我們在主頁源中找到iframe標記。完蛋了!主頁源中沒有iframe標記。


JavaScript-Generated標記
++++++++++++++++++++++++

這個遺漏的原因是view page source顯示了從服務器獲取的內容。但是，由瀏覽器呈現的最終DOM(文檔對象模型)可能非常不同。JavaScript開始工作，可以隨意操縱DOM。無法找到iframe，因為從服務器檢索頁面時，它就是不存在。

----

靜態抓取 vs. 動態抓取
=====================

靜態抓取會忽略 JavaScript, 它可以不依靠瀏覽器而直接從服務器端獲取網頁代碼. 這就是你通過"查看源碼"所看到的東西, 然後你就可以進行信息提取了. 如果你要查找的內容已經存在於源碼中, 那就不需要進一步的動作了. 可是, 如果你要查找的內容像上文的 Disqus 評論一樣被嵌入iframe 中, 你就必須使用動態爬取來獲取內容.

動態爬取使用一個真實的瀏覽器(或無界面瀏覽器), 它先讓頁面內的 JavaScript 運行起來, 完成動態內容處理加載. 之後, 它再通過查詢 DOM 來獲取所要尋找的內容. 有時候, 你還需要讓瀏覽器自動模擬人的操作來得到你所需要的內容.

使用 Requests 和 BeautifulSoup 進行靜態抓取
===========================================

讓我們來看看如何使用 Python 的兩個經典包來進行靜態抓取: requests 用來抓取網頁內容. BeautifulSoup用來解析 HTML.

安裝 Requests 和 BeautifulSoup
==============================

首先安裝 pipenv, 然後運行命令: pipenv install requests beautifulsoup4

它首先為你創建一個虛擬環境, 然後安裝這兩個包在虛擬環境裡. 如果你的代碼在gitlab上, 你可以使用命令 pipenv install 來安裝.

----

獲取網頁內容
============

用 requests 抓取網頁內容只需要一行代碼: r = requests.get(url).

代碼返回一個 response 對象, 它包含大量有用的屬性. 其中最重要的屬性是 ok 和 content. 如果請求失敗, r.ok 為 False 並且 r.content 包含該錯誤信息. content 代表一個字節流, 做文本處理時, 你最好將它解碼成 utf-8.

.. code-block:: bash

  >>> r = requests.get('http://www.c2.com/no-such-page')
  >>> r.ok
  False
  >>> print(r.content.decode('utf-8'))
  <!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
  <html><head>
  <title>404 Not Found</title>
  </head><body>
  <h1>Not Found</h1>
  <p>The requested URL /ggg was not found on this server.</p>
  <hr>
  <address>
  Apache/2.0.52 (CentOS) Server at www.c2.com Port 80
  </address>
  </body></html>

如果代碼正常返回沒有報錯, 那 r.content 會包含請求的網頁源碼(就是"查看源碼"所看到的內容).

用 BeautifulSoup 查找元素
=========================

下面的 get_page() 函數會獲取給定 URL 的網頁源碼, 然後解碼成 utf-8, 最後再將 content 傳遞給 BeautifulSoup 對象並返回, BeautifulSoup 使用 HTML 解析器進行解析.

.. code-block:: python

  def get_page(url):
      r = requests.get(url)
      content = r.content.decode('utf-8')
      return BeautifulSoup(content, 'html.parser')

我們獲取到 BeautifulSoup 對象後, 就可以開始解析所需要的信息了.

BeautifulSoup 提供了很多查找方法來定位網頁中的元素, 並可以深入挖掘出嵌套的元素.

Tuts+ 網站包含了很多培訓教程, `這裡`_ 是我的主頁. 在每一個頁面包含最多12篇教程, 如果你已經獲取了12篇的教程, 你就可以進入下一頁面了. 每一篇文章都被 <article> 標籤包圍著. 下面的函數就是發現頁面裡的所有 article 元素, 然後找到對應的鏈接, 最後提取出教程的 URL.

.. code-block:: python

  def get_page_articles(page):
      elements = page.findAll('article')
      articles = [e.a.attrs['href'] for e in elements]
      return articles

The following code gets all the articles from my page and prints them (without the common prefix):

.. code-block:: bash

  page = get_page('https://tutsplus.com/authors/gigi-sayfan')
  articles = get_page_articles(page)
  prefix = 'https://code.tutsplus.com/tutorials'
  for a in articles:
      print(a[len(prefix):])

  Output:

  building-games-with-python-3-and-pygame-part-5--cms-30085
  building-games-with-python-3-and-pygame-part-4--cms-30084
  building-games-with-python-3-and-pygame-part-3--cms-30083
  building-games-with-python-3-and-pygame-part-2--cms-30082
  building-games-with-python-3-and-pygame-part-1--cms-30081
  mastering-the-react-lifecycle-methods--cms-29849
  testing-data-intensive-code-with-go-part-5--cms-29852
  testing-data-intensive-code-with-go-part-4--cms-29851
  testing-data-intensive-code-with-go-part-3--cms-29850
  testing-data-intensive-code-with-go-part-2--cms-29848
  testing-data-intensive-code-with-go-part-1--cms-29847
  make-your-go-programs-lightning-fast-with-profiling--cms-29809

----

使用 Selenium 動態爬取
++++++++++++++++++++++

安裝 Selenium
=============

用這個命令安裝 Selenium：pipenv install selenium

選擇你的 Web 驅動
=================

Selenium 需要一個 Web 驅動（自動化用的瀏覽器）。對於網頁爬取來說，一般不需要在意選用哪個驅動。我建議使用 Chrome 驅動。 `Selenium 手冊`_ 中有相關的介紹。

對比 Chrome 和 PhantomJS
========================

某些情況下你可能想用沒有用戶界面的（headless）瀏覽器。理論上來說，PhantomJS 正好就是那款 Web 驅動。但是實際上有人報告一些只會在 PhantomJS 中出現的問題，這些問題在 Selenium 使用 Chrome 或 Firefox 時並不會出現。我喜歡從等式中刪除這一變量，使用實際的 Web 瀏覽器驅動。

----

統計 Disqus 評論數量
====================

我們來搞點動態抓取，使用 Selenium 統計 Tuts+ 手機的 Disqus 評論數量。下面需要導入的內容。

.. code-block:: python

  from selenium import webdriver
  from selenium.webdriver.common.by import By
  from selenium.webdriver.support.expected_conditions import (
      presence_of_element_located)
  from selenium.webdriver.support.wait import WebDriverWait

get_comment_count() 函數需要傳入 Selenium 驅動和 URL 作為參數。它使用驅動的 get() 方法從 URL 獲取內容。這和requests.get()相似，其不同之處在於使用驅動對象管理 DOM 的實時呈現。

然後，它獲取教程的標題，並使用 iframe 的父級 id，disqus_thread，和 iframe 標籤來定位 iframe：

.. code-block:: python

  def get_comment_count(driver, url):
      driver.get(url)
      class_name = 'content-banner__title'
      name = driver.find_element_by_class_name(class_name).text
      e = driver.find_element_by_id('disqus_thread')
      disqus_iframe = e.find_element_by_tag_name('iframe')
      iframe_url = disqus_iframe.get_attribute('src')

接下來獲取 iframe 的內容。注意我們要等到 comment-count 元素出現，因為評論是動態加載的，不一定可用。

.. code-block:: python

  driver.get(iframe_url)
  wait = WebDriverWait(driver, 5)
  commentCountPresent = presence_of_element_located(
      (By.CLASS_NAME, 'comment-count'))
  wait.until(commentCountPresent)

  comment_count_span = driver.find_element_by_class_name(
      'comment-count')
  comment_count = int(comment_count_span.text.split()[0])

最後部分是返回最新的評論, 當然不包括我自己的評論. 方法是檢查我還沒有回覆的評論.

.. code-block:: python

  last_comment = {}
  if comment_count > 0:
      e = driver.find_elements_by_class_name('author')[-1]
      last_author = e.find_element_by_tag_name('a')
      last_author = e.get_attribute('data-username')
      if last_author != 'the_gigi':
          e = driver.find_elements_by_class_name('post-meta')
          meta = e[-1].find_element_by_tag_name('a')
          last_comment = dict(
              author=last_author,
              title=meta.get_attribute('title'),
              when=meta.text)
  return name, comment_count, last_comment


結論
++++

網頁爬取是一個非常實用的技術, 尤其當你需要處理的信息瀏覽器並不提供有用的API支持的時候. 它通常需要一些技巧來從現代web應用中提取信息, 不過一些成熟的、設計良好的工具, 比如: requests、BeautifulSoup、Selenium 都會減輕你的工作並提高效率.

最後, 你可以試一下我寫的一些工具, 它們在 `Envato Market`_ 有售, 歡迎提問和反饋.

----

.. highlights::

  | 本文中的所有譯文僅用於學習和交流目的，轉載請務必註明文章譯者、出處、和本文鏈接。
  | 我們的翻譯工作遵照 `CC 協議`_ ，如果我們的工作有侵犯到您的權益，請及時聯繫我們。


.. _Modern Web Scraping With BeautifulSoup and Selenium: https://code.tutsplus.com/tutorials/modern-web-scraping-with-beautifulsoup-and-selenium--cms-30486
.. _使用 BeautifulSoup 和 Selenium 进行网页爬取 - 开源中国: https://www.oschina.net/translate/modern-web-scraping-with-beautifulsoup-and-selenium--cms
.. _ZICK_ZEON: https://my.oschina.net/u/658291
.. _xiaoaiwhc1: https://my.oschina.net/xiaoaiwhc
.. _边城: https://my.oschina.net/jamesfancy
.. _Vagrant技術入門: https://code.tutsplus.com/tutorials/introduction-to-vagrant--cms-25917
.. _這裡: https://tutsplus.com/authors/gigi-sayfan
.. _Selenium 手冊: http://selenium-python.readthedocs.io/installation.html#drivers
.. _Envato Market: https://codecanyon.net/search/python
.. _CC 協議: https://zh.wikipedia.org/wiki/Wikipedia:CC
