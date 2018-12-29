微軟正在為 Windows 10 添加對 AVIF 圖像的支持
############################################

:date: 2018-12-29T11:08+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: 局长(開源中國)
:summary: 據外媒報導，Windows 10 1903 版本是微軟 Windows 10 操作系統的下一個功能更新版本，將支持 AVIF 圖像格式。AVIF 或 AV1 圖像文件格式是基於 AV1 的新圖像格式，它使用 HEIF 作為容器和 AV1 幀。
:og_image: https://oscimg.oschina.net/oscnet/477303090e35b197780d215e02f15e7dbec.jpg


`據外媒報導`_ ，Windows 10 1903 版本是微軟 Windows 10 操作系統的下一個功能更新版本，將支持 AVIF 圖像格式。AVIF 或 AV1 圖像文件格式是基於 AV1 的新圖像格式，它使用 HEIF 作為容器和 AV1 幀。

微軟在 Windows 10 1809 版本中引入了對 AV1 格式視頻的支持，但使用者必須安裝 AV1 視頻擴展，以便為 Windows 10 設備添加對 AV1 視頻的支持。不過在本文發佈時，該擴展程序仍處於測試版階段。

.. image:: https://oscimg.oschina.net/oscnet/477303090e35b197780d215e02f15e7dbec.jpg
   :alt: 微軟正在為 Windows 10 添加對 AVIF 圖像的支持
   :align: center

事實上，從 Windows 10 Insider build 18305 開始，微軟已開始添加對 AVIF 圖像格式的支持。此支持首先添加到畫板和文件資源管理器這兩個程序中，但可以通過 API 擴展到其他應用程序。

.. image:: https://oscimg.oschina.net/oscnet/0be185d938ec07250c8552d133f7510daf3.jpg
   :alt: 微軟正在為 Windows 10 添加對 AVIF 圖像的支持
   :align: center

**Build 1809 不會生成 AVIF 格式圖像的縮略圖**

.. image:: https://oscimg.oschina.net/oscnet/f7e04d1d5a11adcfd2ac42efa5a6f0c5c09.jpg
   :alt: 微軟正在為 Windows 10 添加對 AVIF 圖像的支持
   :align: center

**但是在 Build 18305 中，可以在文件資源管理器中預覽和查看此圖像格式的縮略圖**

此外，微軟的畫板程序也支持 AVIF 文件，並且是與此文件格式關聯的默認應用程序。

.. image:: https://oscimg.oschina.net/oscnet/97cad8da5cac50b47797ad07967081ccbdd.jpg
   :alt: 微軟正在為 Windows 10 添加對 AVIF 圖像的支持
   :align: center

微軟並不是唯一一個致力於添加對 AVIF 圖像格式支持的組織。VLC 早已 `引入了支持`_ ，Mozilla 也計畫在未來 `增加支持`_ 。

.. image:: https://oscimg.oschina.net/oscnet/b39318cb6fd55f96a598a640725bdffbdf3.jpg
   :alt: 微軟正在為 Windows 10 添加對 AVIF 圖像的支持
   :align: center

AV1 已得到了所有主要瀏覽器開發商和其他主要技術公司的支持。雖然目前尚不清楚這是否足以使新格式成為視頻和圖像的事實標準，但有理由相信它比以前的嘗試更好。那麼你對 AV1 和 AVIF 又有什麼期望呢，歡迎留言討論。

..
  .. image:: 
   :alt: 
   :align: center

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：微軟正在為 Windows 10 添加對 AVIF 圖像的支持
  | 本文地址：https://www.oschina.net/news/103135/windows-10-1903-to-support-avif-format

.. _據外媒報導: https://www.bleepingcomputer.com/news/microsoft/microsoft-is-adding-avif-image-support-to-windows-10/
.. _引入了支持: https://git.videolan.org/?p=vlc.git;a=commit;h=2f3675ebd867b9a23152fa38674e8dbffccbeabb
.. _增加支持: https://bugzilla.mozilla.org/show_bug.cgi?id=avif
