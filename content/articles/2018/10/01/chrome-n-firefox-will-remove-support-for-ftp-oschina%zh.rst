Chrome 與 Firefox 將取消對 FTP 的支持
#####################################

:date: 2018-11-28T08:59+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: h4cd(開源中國)
:summary: 據 BleepingComputer 報導，Chrome 與 Firefox 都在採取措施取消對 FTP 協議的支持。
:og_image: https://oscimg.oschina.net/oscnet/6671a8c3a823a3f7cd0a0ae4692d2ba400a.jpg

據 BleepingComputer 報導，Chrome 與 Firefox 都在採取措施 `取消對 FTP 協議的支持`_ 。

首先是 Chrome，根據 BleepingComputer 的追蹤，其實開發者早在 2014 年就提出要取消對 FTP 的支持，但是因為影響較大一直未能實施。而在即將到來的版本中，有一個關於如何在 Chrome 上加載圖片的變化似乎預示著 Google 朝著完全拋棄 FTP 的方向邁出了一步。

.. image:: https://oscimg.oschina.net/oscnet/6671a8c3a823a3f7cd0a0ae4692d2ba400a.jpg
   :alt: Chrome 與 Firefox 將取消對 FTP 的支持
   :align: center

如上圖所示，當用戶使用 Chrome 打開 FTP 服務器上打開文件時，它將嘗試在瀏覽器中渲染並呈現該文件，這是目前瀏覽器支持 FTP 渲染的方式，然而谷歌覺得這需要改變，因為 FTP 是不安全的協議。新的方式將不再渲染並呈現像上面這張圖片一樣擁有完整路徑的文件，而是直接將其下載。但是當訪問目標是“ftp://目錄”時，Chrome 還會像以往一樣列出相關目錄項。Chrome 開發者表示這是不得以的舉措，因為目前還無法直接砍掉 FTP，它的使用率有些高，所以只能用這樣一種“合理的方法”，降低通過 FTP 進行網絡攻擊的風險。

而 Firefox 方面，在 Bugzilla 上一個關於 `支持 FTP over SSL 的 issue`_ 上，Firefox 開發者指出：“因為我們遲早要完全棄用 FTP，所以不應該在代碼庫中添加更多相關的代碼。”

.. image:: https://oscimg.oschina.net/oscnet/a07ef91589d3e869aedebbf05b35c7011d1.jpg
   :alt: Chrome 與 Firefox 將取消對 FTP 的支持
   :align: center

FTP 協議的安全性問題主要包括它以明文傳輸，數據容易被竊取，並且由於其工作機制，使客戶端可以直接監聽服務器等。

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：Chrome 與 Firefox 將取消對 FTP 的支持
  | 本文地址：https://www.oschina.net/news/102189/chrome-n-firefox-will-remove-support-for-ftp

.. _取消對 FTP 協議的支持: https://www.bleepingcomputer.com/news/google/chrome-and-firefox-developers-aim-to-remove-support-for-ftp/
.. _支持 FTP over SSL 的 issue: https://bugzilla.mozilla.org/show_bug.cgi?id=85464
