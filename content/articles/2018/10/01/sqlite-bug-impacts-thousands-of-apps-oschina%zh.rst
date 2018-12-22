SQLite 被曝存在漏洞，所有 Chromium 瀏覽器受影響
###############################################

:date: 2018-12-16T09:15+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: h4cd(開源中國)
:summary: SQLite 被曝存在一個影響數千應用的漏洞，受害應用包括所有基於 Chromium 的瀏覽器。
:og_image: https://upload.wikimedia.org/wikipedia/commons/3/38/SQLite370.svg


SQLite 被曝存在一個影響數千應用的漏洞，受害應用包括所有基於 Chromium 的瀏覽器。

據 `ZDNet 報導`_ ，該漏洞由騰訊 Blade 安全團隊發現，允許攻擊者在受害者的計算機上運行惡意代碼，並在危險較小的情況下洩漏程序內存或導致程序崩潰。由於 SQLite 嵌入在數千個應用程序中，因此該漏洞會影響各種軟件，包括物聯網設備、桌面軟件、Web 瀏覽器、Android 與 iOS 應用。

如果底層瀏覽器支持 SQLite 和 Web SQL API，那麼將漏洞利用代碼轉換為常規 SQL 語法也可以通過訪問網頁等操作遠程利用此漏洞。Chromium 瀏覽器引擎支持此 API，這意味著像 Chrome、Vivaldi、Opera 和 Brave 等瀏覽器都會受到影響，而 Firefox 和 Edge 由於不支持此 API，因此不受影響。

騰訊 Blade  研究人員表示，他們在今年秋季早些時候向 SQLite 團隊報告了此問題，SQLite 3.26.0 中進行了修復。Chrome 71 已經解決了該問題，但是 Opera 的 Chromium 版本還沒有更新，這意味著它很可能還會受到影響。

另一方面，雖然 Firefox 不支持 Web SQL API，不會受到遠程利用攻擊，但是其自帶了一個本地可訪問的 SQLite 數據庫，這意味著本地攻擊者可能利用此漏洞來執行代碼。

ZDNet 表示，由於開發者很少更新代碼庫及其應用的組件部分，因此很可能這個漏洞在接下來幾年內還會持續產生影響，因此，騰訊 Blade 團隊表示暫時不會發佈任何概念驗證漏洞利用代碼。

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：SQLite 被曝存在漏洞，所有 Chromium 瀏覽器受影響
  | 本文地址：https://www.oschina.net/news/102738/sqlite-bug-impacts-thousands-of-apps

.. _ZDNet 報導: https://www.zdnet.com/article/sqlite-bug-impacts-thousands-of-apps-including-all-chromium-based-browsers/
