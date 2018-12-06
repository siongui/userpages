期待已久，Flutter 1.0 正式發佈！
################################

:date: 2018-12-05T09:16+08:00
:tags: 轉錄, 開源中國
:category: 轉錄
:author: 王練(開源中國)
:summary: 在昨晚的 Flutter Live 2018 上，Google 宣佈 Flutter 1.0 正式發佈。這是一個基於 Dart 的移動開發平台，旨在幫助開發者在 iOS 和 Android 兩個平台上開發高質量的原生應用界面。此外，Google 還宣佈了 Flutter 運行時基於 Web 的實驗性實現，旨在將 Flutter 應用引入標準 Web 瀏覽器。
:og_image: https://oscimg.oschina.net/oscnet/a259ea1ad55a0ac98d3c6293caa0489ae49.jpg

在昨晚的 Flutter Live 2018 上，Google 宣佈 Flutter 1.0 正式發佈。這是一個基於 Dart 的移動開發平台，旨在幫助開發者在 iOS 和 Android 兩個平台上開發高質量的原生應用界面。此外，Google 還宣佈了 Flutter 運行時基於 Web 的實驗性實現，旨在將 Flutter 應用引入標準 Web 瀏覽器。

.. image:: https://oscimg.oschina.net/oscnet/a259ea1ad55a0ac98d3c6293caa0489ae49.jpg
   :alt: 期待已久，Flutter 1.0 正式發佈！
   :align: center

從我們首次看到 Flutter 的 Beta 測試版，到如今的 1.0 正式版，總共經過了 9個多月。

- 2 月底在世界移動大會 (MWC) 上發佈首個 Beta 版;

- 5 月的 Google I/O 大會上發佈 Beta 3 ;

- 6 月底的 GMTC 發佈首個預覽版；

- 9 月的谷歌開發者大會上，發佈 `預覽版 2`_ ；

- 12月初的 Flutter Live 2018 上，發佈1.0 穩定版。

Flutter 1.0 主要聚焦於穩定性和 bug 修復，同時還包含兩項新功能的預覽 ——  Add to App 和 platform views：

Flutter 1.0 使用的是最新的 `Dart 2.1`_ 。Dart 2.1 提供更小的代碼體積，更快的類型檢查和更好的診斷。按照 Google 的說法，Dart 2.1 將輸出代碼的大小減少了 17％，並將編譯時間縮短了 15％ 。Dart 2.1 還包含新的語言特性，例如通過新mixin關鍵字改進 mixin 支持、支持 int-to-double 的轉換、編譯時類型檢查、新的 HTTP 狀態碼等等。

Add to App 是一種逐步將現有應用移植到 Flutter 的方法，主要用於將 Flutter 用於現有應用，或者將現有應用分階段轉換為 Flutter 。

platform views 則是一種相反的方式，用於將 Android 或 iPhone 平台的控件嵌入到 Flutter 應用。

.. image:: https://oscimg.oschina.net/oscnet/a6e8d868d4ea596d623307f42dffb94e0ea.jpg
   :alt: 期待已久，Flutter 1.0 正式發佈！
   :align: center

如開頭所述，Google 還宣佈了 Flutter 運行時基於 Web 的實驗性實現 —— Hummingbird_ ，旨在將 Flutter 應用引入瀏覽器。它利用 Dart 平台的特性不僅可以編譯原生 ARM 代碼，還可以編譯 JavaScript 。這使得 Flutter 代碼可以在基於標準的 Web 上運行而無需任何更改。

Hummingbird 可讓 Flutter 覆蓋更多平台，包括 Windows、macOS 和 Linux 。

.. image:: https://oscimg.oschina.net/oscnet/996a4cf1ed2b754e34bfc721069d4f91f61.jpg
   :alt: 期待已久，Flutter 1.0 正式發佈！
   :align: center

發行說明：

https://developers.googleblog.com/2018/12/flutter-10-googles-portable-ui-toolkit.html

.. highlights::

  | 本站文章除註明轉載外，均為本站原創或編譯。歡迎任何形式的轉載，但請務必註明出處，尊重他人勞動共創開源社區。
  | 轉載請註明：文章轉載自 開源中國社區 [https://www.oschina.net]
  | 本文標題：期待已久，Flutter 1.0 正式發佈！
  | 本文地址：https://www.oschina.net/news/102408/flutter-1-0-released

.. _預覽版 2: http://mp.weixin.qq.com/s?__biz=MzAwODY4OTk2Mg==&mid=2652047108&idx=2&sn=242ef51c15fef1f386ebb2ac322e490c&chksm=808ca741b7fb2e5723a16b633f12cacb124b60bd37b3c59929aae3610628e86326885275467f&scene=21#wechat_redirect
.. _Dart 2.1: https://www.oschina.net/news/102002/dart-2-1-released
.. _Hummingbird: https://medium.com/p/e687c2a023a8
